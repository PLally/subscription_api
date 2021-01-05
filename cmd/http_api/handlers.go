package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Resource struct {
	Create http.HandlerFunc
	Index  http.HandlerFunc
	Delete http.HandlerFunc
	Get    http.HandlerFunc

	Router *mux.Router
	model  interface{}
}

func resources(r *mux.Router, model interface{}, DB *gorm.DB) *Resource {
	resource := &Resource{
		Create: createHandler(model, DB),
		Index:  indexHandler(model, DB),
		Delete: deleteHandler(model, DB),
		Get:    getHandler(model, DB),

		model: model,
	}
	name := reflect.TypeOf(resource.model).Name()
	tableName := DB.NamingStrategy.TableName(name)

	resource.addHandlers(tableName, r)
	return resource
}

func (resource *Resource) Use(middleware mux.MiddlewareFunc) {
	resource.Router.Use(middleware)
}

func (resource *Resource) addHandlers(tableName string, r *mux.Router) {
	prefixName := "/" + tableName
	r = r.PathPrefix(prefixName).Subrouter()

	resource.Router = r
	r.HandleFunc("/{id:[0-9]+}", resource.Get).
		Methods("GET")

	r.HandleFunc("", resource.Index).
		Methods("GET")

	r.HandleFunc("", resource.Create).
		Methods("POST")

	r.Handle("/{id:[0-9]+}", resource.Delete).
		Methods("DELETE")
}

func constructWhere(values url.Values, isAllowed func(string) bool) (string, []interface{}) {
	var conditionValues []interface{}
	var condition []string

	for k, v := range values {
		if isAllowed(k) {
			conditionValues = append(conditionValues, v[0])
			condition = append(condition, k+"=?")
		}
	}

	return strings.Join(condition, " AND "), conditionValues
}

func indexHandler(model interface{}, DB *gorm.DB) http.HandlerFunc {
	t := reflect.TypeOf(model)
	modelType := reflect.SliceOf(
		t,
	)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var dbModel = reflect.New(modelType).Interface()

		condition, values := constructWhere(r.URL.Query(), func(s string) bool {

			s = strings.ReplaceAll(s, "_", "")
			_, ok := t.FieldByNameFunc(func(name string) bool {
				name = strings.ToLower(name)
				return name == s
			})
			return ok
		})
		db := DB.Where(condition, values...)
		if joinable, ok := model.(database.Joinable); ok {
			db = joinable.DoJoins(db)
		}
		db.Find(dbModel)
		writeJson(w, dbModel, http.StatusOK)
	}
}

func getHandler(model interface{}, DB *gorm.DB) http.HandlerFunc {
	modelType := reflect.TypeOf(model)
	name := DB.NamingStrategy.TableName(modelType.Name())

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var dbModel = reflect.New(modelType).Interface()

		vars := mux.Vars(r)
		idString := vars["id"]
		id, _ := strconv.Atoi(idString)

		db := DB

		if joinable, ok := dbModel.(database.Joinable); ok {
			db = joinable.DoJoins(db)
		}
		db = db.First(dbModel, id)
		if db.RowsAffected < 1 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("%v does not exist", name)))
			return
		}

		writeJson(w, dbModel, http.StatusOK)

	}
}

func createHandler(model interface{}, DB *gorm.DB) http.HandlerFunc {
	modelType := reflect.TypeOf(model)
	name := DB.NamingStrategy.TableName(modelType.Name())

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var dbModel = reflect.New(modelType).Interface()

		unmarshalledBody := r.Context().Value("unmarshalled_body")

		if unmarshalledBody == nil {
			data, _ := ioutil.ReadAll(r.Body)
			err := json.Unmarshal(data, dbModel)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid json"))
				return
			}
		} else {
			dbModel = unmarshalledBody
		}

		db := DB.Create(dbModel)

		err := db.Error
		if err != nil {
			onDatabaseError(err, w, name)
			return
		}

		if db.RowsAffected < 1 {
			return
		}

		writeJson(w, dbModel, http.StatusCreated)
	}
}

func deleteHandler(model interface{}, DB *gorm.DB) http.HandlerFunc {
	modelType := reflect.TypeOf(model)
	name := DB.NamingStrategy.TableName(modelType.Name())
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var dbModel = reflect.New(modelType).Interface()

		vars := mux.Vars(r)
		idString := vars["id"]
		id, _ := strconv.Atoi(idString)

		if err := DB.Raw(fmt.Sprintf("DELETE FROM %v WHERE id=? RETURNING *", name), id).Scan(dbModel).Error; err != nil {
			onDatabaseError(err, w, name)
			return
		}

		writeJson(w, dbModel, http.StatusOK)
	}

}

func subscribeHandler(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		var subCreateStruct struct {
			DestinationType       string `json:"destination_type"`
			DestinationIdentifier string `json:"destination_identifier"`
			SubscriptionType      string `json:"subscription_type"`
			SubscriptionTags      string `json:"subscription_tags"`
		}
		_ = json.Unmarshal(data, &subCreateStruct)
		handler := subscription.GetSubTypeHandler(subCreateStruct.SubscriptionType)
		if handler == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid subscription type"))
			return
		}
		tags, err := handler.Validate(subCreateStruct.SubscriptionTags)
		subCreateStruct.SubscriptionTags = tags
		if err != nil {
			log.Info(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid subscription type tags"))
			return
		}

		subtype := database.SubscriptionType{
			Type: subCreateStruct.SubscriptionType,
			Tags: subCreateStruct.SubscriptionTags,
		}
		dest := database.Destination{
			DestinationType:    subCreateStruct.DestinationType,
			ExternalIdentifier: subCreateStruct.DestinationIdentifier,
		}

		DB.FirstOrCreate(&subtype, subtype)
		DB.FirstOrCreate(&dest, dest)

		sub := database.Subscription{
			SubscriptionTypeID: subtype.ID,
			DestinationID:      dest.ID,
		}
		status := http.StatusOK
		if DB.FirstOrCreate(&sub, sub).RowsAffected == 0 {
			status = http.StatusConflict
		}
		sub.Destination = dest
		sub.SubscriptionType = subtype
		writeJson(w, sub, status)
	}
}

func writeJson(w http.ResponseWriter, obj interface{}, status int) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.WriteHeader(status)
	w.Write(data)
}

func onDatabaseError(err error, w http.ResponseWriter, name string) {
	switch err := err.(type) {
	case *pq.Error:
		switch err.Code.Name() {
		case "foreign_key_violation":
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(
				"Foreign key violation",
			))
		case "unique_violation":
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(
				fmt.Sprintf("That %v already exists", name),
			))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Code.Name()))
		}
	default:
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 not found"))
		}
	}
}

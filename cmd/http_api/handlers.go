package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)
type Resource struct {
	Create http.HandlerFunc
	Index http.HandlerFunc
	Delete http.HandlerFunc
	Get http.HandlerFunc

	Router *mux.Router
	model interface{}
}
func resources(r *mux.Router, model interface{}, DB *gorm.DB) *Resource{
	resource := &Resource{
		Create: createHandler(model, DB),
		Index: indexHandler(model, DB),
		Delete: deleteHandler(model, DB),
		Get: getHandler(model, DB),

		model: model,
	}
	resource.addHandlers(r)
	return resource
}

func (resource *Resource) Use(middleware mux.MiddlewareFunc) {
	resource.Router.Use(middleware)
}

func (resource *Resource) addHandlers(r *mux.Router){
	name := reflect.TypeOf(resource.model).Name()
	prefixName := "/" + gorm.ToTableName(name) + "s"
	r = r.PathPrefix(prefixName).Subrouter()
	fmt.Println(prefixName)
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
func indexHandler(model interface{}, DB *gorm.DB) http.HandlerFunc {
	modelType := reflect.SliceOf(
		reflect.TypeOf(model),
	)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var dbModel = reflect.New(modelType).Interface()

		DB.Find(dbModel)
		writeJson(w, dbModel, http.StatusOK)
	}
}

func getHandler(model interface{}, DB *gorm.DB) http.HandlerFunc {
	modelType := reflect.TypeOf(model)
	name := gorm.ToTableName(modelType.Name())

	return func(w http.ResponseWriter, r *http.Request) {
		var dbModel = reflect.New(modelType).Interface()

		vars := mux.Vars(r)
		idString := vars["id"]
		id, _ := strconv.Atoi(idString)

		db := DB.First(dbModel, id)
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
	name := gorm.ToTableName(modelType.Name())

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var dbModel = reflect.New(modelType).Interface()


		unmarshalledBody := r.Context().Value("unmarshalled_body")

		if unmarshalledBody == nil {
			fmt.Println("Creating database entry")
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
	name := gorm.ToTableName(modelType.Name()) + "s"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var dbModel = reflect.New(modelType).Interface()

		vars := mux.Vars(r)
		idString := vars["id"]
		id, _ := strconv.Atoi(idString)

		if err := DB.Raw(fmt.Sprintf("DELETE FROM %v WHERE id=? RETURNING *", name), id).Scan(dbModel).Error
		err != nil {
			onDatabaseError(err, w, name)
			return
		}

		writeJson(w, dbModel, http.StatusOK)
	}

}
func writeJson(w http.ResponseWriter, obj interface{}, status int) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}
	w.Write(data)
}

func onDatabaseError(err error, w http.ResponseWriter, name string) {
	switch err := err.(type) {
	case *pq.Error:
		switch err.Code.Name() {
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
			w.WriteHeader(http.StatusNotFound,)
			w.Write([]byte("404 not found"))
		}
	}
}

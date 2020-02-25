package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type subscriptionServer struct {
	*gorm.DB
}

func (db subscriptionServer) createHandler(getModel func() interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		model := getModel()
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(model)
		if err != nil {
			log.Error(err)
		}

		err = db.Create(model).Error
		if err != nil {
			log.Error(err)
		}

		if err != nil {
			log.Error("db: ", err)
		}

		data, err := json.Marshal(model)
		if err != nil {
			log.Error(err)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(data)

	}
}

func (db subscriptionServer) indexHandler(getModel interface{}) {

}

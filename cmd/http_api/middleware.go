package main

import (
	"context"
	"encoding/json"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func CheckSubscriptionType(h http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			h.ServeHTTP(w, r)
			return
		}
		var dbModel database.SubscriptionType
		data, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(data, &dbModel)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid json"))
			return
		}
		handler := subscription.GetSubTypeHandler(dbModel.Type)
		if handler == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid subscription type"))
			return
		}
		dbModel.Tags, err = handler.Validate(dbModel.Tags)
		if err != nil {
			log.Info(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid subscription type tags"))
			return
		}

		ctx := context.WithValue(r.Context(), "unmarshalled_body", &dbModel)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
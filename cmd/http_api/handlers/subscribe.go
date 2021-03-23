package handlers

import (
	"encoding/json"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

func SubscribeHandler(DB *gorm.DB) http.HandlerFunc {
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

		DB.FirstOrCreate(&dest, dest)
		DB.FirstOrCreate(&subtype, subtype)

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

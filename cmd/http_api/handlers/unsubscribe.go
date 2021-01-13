package handlers

import (
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/subscription"
	"net/http"
	"gorm.io/gorm"
	"io/ioutil"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func UnsubscribeHandler(DB *gorm.DB) http.HandlerFunc {
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

		DB.First(&subtype, subtype)
		DB.First(&dest, dest)

		sub := database.Subscription{
			SubscriptionTypeID: subtype.ID,
			DestinationID:      dest.ID,
		}


		status := http.StatusNoContent

		if DB.Delete(&sub, sub).RowsAffected == 0 {
			status = http.StatusNotFound
		}

		sub.Destination = dest
		sub.SubscriptionType = subtype
		w.WriteHeader(status)
	}
}

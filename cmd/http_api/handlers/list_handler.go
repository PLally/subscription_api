package handlers

import (
	"github.com/plally/subscription_api/database"
	"gorm.io/gorm"
	"net/http"
)

func ListHandler(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var subscriptions []database.Subscription
		dest := database.Destination{
			ExternalIdentifier: r.URL.Query().Get("destination_identifier"),
			DestinationType:    r.URL.Query().Get("destination_type"),
		}
		DB.First(&dest, dest)
		db := DB.Model(database.Subscription{}).Where("destination_id = ?", dest.ID)

		db = database.Subscription{}.DoJoins(db)
		db.Find(&subscriptions)
		writeJson(w, subscriptions, http.StatusOK)
	}
}

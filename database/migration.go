package database

import (
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&SubscriptionType{}, &Destination{}, &Subscription{})

	db.Model(&Subscription{}).
		AddForeignKey("destination_id", "destinations(id)", "RESTRICT", "RESTRICT")
	db.Model(&Subscription{}).
		AddForeignKey("subscription_type_id", "subscription_types(id)", "RESTRICT", "RESTRICT")
}

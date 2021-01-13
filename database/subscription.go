package database

import (
	"gorm.io/gorm"
	"time"
)

type Subscription struct {
	ID                 uint             `gorm:"primary_key" json:"id"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
	Destination        Destination      `json:"destination"`
	SubscriptionType   SubscriptionType `json:"subscription_type"`
	DestinationID      uint             `gorm:"UniqueIndex:idx_destination_subtype" json:"destination_id"`
	SubscriptionTypeID uint             `gorm:"UniqueIndex:idx_destination_subtype" json:"subscription_type_id"`
	LastItem           int64            `json:"last_item"`
}

func (s Subscription) HasDispatched() bool {
	return s.LastItem != 0
}
func (s Subscription) DoJoins(db *gorm.DB) *gorm.DB {
	return db.Joins("Destination").Joins("SubscriptionType")
}


type Joinable interface {
	DoJoins(db *gorm.DB) *gorm.DB
}

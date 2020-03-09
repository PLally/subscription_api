package database

import "github.com/jinzhu/gorm"

type SubscriptionType struct {
	gorm.Model
	Type          string `gorm:"unique_index:idx_type_tags" json:"type"`
	Tags          string `gorm:"unique_index:idx_type_tags" json:"tags"`
	Subscriptions []Subscription
}

type Destination struct {
	gorm.Model
	ExternalIdentifier string `gorm:"unique_index:idx_destination_identifier" json:"external_identifier"`
	DestinationType    string `gorm:"unique_index:idx_destination_identifier" json:"destination_type"`
}

type Subscription struct {
	gorm.Model
	Destination        Destination
	SubscriptionType   SubscriptionType
	DestinationID      uint `gorm:"unique_index:idx_destination_subtype"`
	SubscriptionTypeID uint `gorm:"unique_index:idx_destination_subtype"`
	LastItem           int64
}

package database

import (
	"time"
)

type SubscriptionType struct {
	ID            uint `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Type          string         `gorm:"unique_index:idx_type_tags" json:"type"`
	Tags          string         `gorm:"unique_index:idx_type_tags" json:"tags"`
	Subscriptions []Subscription `gorm:"PRELOAD:false"`
}

type Destination struct {
	ID                 uint `gorm:"primary_key"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	ExternalIdentifier string `gorm:"unique_index:idx_destination_identifier" json:"external_identifier"`
	DestinationType    string `gorm:"unique_index:idx_destination_identifier" json:"destination_type"`
}

type Subscription struct {
	ID                 uint `gorm:"primary_key"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Destination        Destination
	SubscriptionType   SubscriptionType
	DestinationID      uint `gorm:"unique_index:idx_destination_subtype" json:"destination_id"`
	SubscriptionTypeID uint `gorm:"unique_index:idx_destination_subtype" json:"subscription_type_id"`
	LastItem           int64
}

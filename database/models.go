package database

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type SubscriptionType struct {
	ID            uint           `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Type          string         `gorm:"UniqueIndex:idx_type_tags" json:"type"`
	Tags          string         `gorm:"UniqueIndex:idx_type_tags" json:"tags"`
	Subscriptions []Subscription `json:"-"`
}

func (sub SubscriptionType) String() string {
	return fmt.Sprintf("[%v] %v - %v", sub.ID, sub.Type, sub.Tags)
}

type Destination struct {
	ID                 uint      `gorm:"primary_key" json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	ExternalIdentifier string    `gorm:"UniqueIndex:idx_destination_identifier" json:"external_identifier"`
	DestinationType    string    `gorm:"UniqueIndex:idx_destination_identifier" json:"destination_type"`
}

func (dest Destination) String() string {
	return fmt.Sprintf("[%v] %v - %v", dest.ID, dest.DestinationType, dest.ExternalIdentifier)
}

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

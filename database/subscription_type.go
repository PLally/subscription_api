package database

import (
	"fmt"
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
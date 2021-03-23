package database

import (
	"fmt"
	"time"
)

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

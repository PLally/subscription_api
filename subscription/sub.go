package subscription

import (
	"fmt"
	"github.com/plally/subscription_api/storage"
)

type SubscriptionItem struct {
	Title       string
	Url         string
	Image       string
	Author      string
	Description string
	TimeID      int64 // must be some sort of identifier that increases with time
}

type SubscriptionTypeHandler interface {
	GetType() string
	GetNewItems(tags string) []SubscriptionItem
	Validate(tags string) (string, error)
}

var typeHandlers = make(map[string]SubscriptionTypeHandler)

func SetSubTypeHandler(subType string, handler SubscriptionTypeHandler) {
	typeHandlers[subType] = handler
}

func GetSubTypeHandler(subType string) SubscriptionTypeHandler {
	return typeHandlers[subType]
}

func Subscribe(db storage.SubscriptionDatabase, subType string, tags string, destinationType string, destinationId string) *storage.Subscription{
	subTypeObj := db.SubscriptionType_Create(storage.SubscriptionType{
		Type: subType,
		Tags: tags,
	})

	destinationObj := db.Destination_Create(storage.Destination{
		ExternalIdentifier: destinationId,
		DestinationType:    destinationType,
	})
	fmt.Println(destinationObj.ID)
	fmt.Println(subTypeObj.ID,)
	subObj := db.Subscription_Create(storage.Subscription{
		DestinationID: destinationObj.ID,
		SubscriptionTypeID: subTypeObj.ID,
		SubscriptionType: subTypeObj,
		Destination: destinationObj,
	})
	fmt.Println(subObj.ID)
	return subObj
}
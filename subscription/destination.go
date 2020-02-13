package subscription

import (
	"github.com/plally/subscription_api/storage"
	log "github.com/sirupsen/logrus"
)

// a destination that a subscription item can be sent to such as a discord channel
type DestinationHandler interface {
	Dispatch(id string, item SubscriptionItem) error
	GetType() string
}

// map between destination types and a handler
var destinationHandlers = make(map[string]DestinationHandler)

func SetDestinationHandler(destType string, handler DestinationHandler) {
	destinationHandlers[destType] = handler
}

func GetDestinationHandler(destType string) DestinationHandler {
	return destinationHandlers[destType]
}

// dispatches all items in the slice to the subscriptions destination
func dispatch(sub storage.Subscription, items []SubscriptionItem) (mostRecent int64) {
	if sub.Destination == nil {
		log.Warn("Subscription does not include a valid destination %v", sub.ID)
		return
	}
	mostRecent = sub.LastItem
	for _, item := range items {
		if sub.LastItem >= item.TimeID { // item has already been dispatched
			continue
		}

		handler := GetDestinationHandler(sub.Destination.DestinationType)
		if handler == nil {
			log.Infof("Unrecognized destination handler %v", sub.Destination.DestinationType)
			return
		}
		handler.Dispatch(sub.Destination.ExternalIdentifier, item)

		if item.TimeID > mostRecent {
			mostRecent = item.TimeID
		}

	}

	return
}

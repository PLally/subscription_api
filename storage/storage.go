package storage
// this package defines an interface for the storage of subscription data
// and 3 structs to represent that data

type SubscriptionDatabase interface {
	SubscriptionType_Get(amount int) chan SubscriptionType
	Subscription_GetWithDestination_BySubType(subType int) chan Subscription
	Subscription_Update(subscription Subscription)

	Subscription_Create(subscription Subscription) *Subscription
	SubscriptionType_Create(SubscriptionType) *SubscriptionType
	Destination_Create(Destination) *Destination
	Connect()
}

type SubscriptionType struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Tags string `json:"tags"`
}

type Destination struct {
	ID                 int
	ExternalIdentifier string
	DestinationType    string
}

type Subscription struct {
	ID                 int
	SubscriptionTypeID int
	Destination        *Destination
	DestinationID      int
	SubscriptionType   *SubscriptionType
	LastItem           int64
}

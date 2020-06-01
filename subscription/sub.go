package subscription

type SubscriptionItem struct {
	Title       string
	Url         string
	Image       string
	Author      string
	Description string
	Tags        string
	Type        string
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

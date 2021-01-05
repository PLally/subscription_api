package destinations

import (
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
)

func init() {
	subscription.SetDestinationHandler("dummy", &DummyDestinationHandler{})
}

type DummyDestinationHandler struct{}

func (d *DummyDestinationHandler) GetType() string {
	return "dummy"
}

func (d *DummyDestinationHandler) Dispatch(id string, item subscription.SubscriptionItem) error {
	log.Infof("[Dummy Destination]: %v, %v (%v)", id, item.Title, item.Url)
	return nil
}

package subscription

// this package contains much of the logic
// checks subscription types for new items
// dispatches those items to their destinations
import (
	"github.com/jinzhu/gorm"
	"github.com/plally/subscription_api/database"
	log "github.com/sirupsen/logrus"
	"sort"
	"sync"
)

// the glue that holds my sphagetti together

// gets new items from a subscription type
func CheckOutDatedSubscriptionTypes(db *gorm.DB, max_workers int) {
	var subTypes []database.SubscriptionType
	db.Limit(1000).Order("updated_at", true).Find(&subTypes)

	log.Debugf("CheckOutDatedSubscriptionTypes: found %v subtypes", len(subTypes))

	subTypeChan := make(chan database.SubscriptionType)
	var wg sync.WaitGroup
	for i := 1; i <= max_workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			checkSubTypesWorker(db, subTypeChan)
		}()
	}

	for _, subType := range subTypes {
		subTypeChan <- subType
	}
	close(subTypeChan)
	wg.Wait()
}

// does the actual work for CheckOutDatedSubscriptionTypes
func checkSubTypesWorker(db *gorm.DB, subTypeChan chan database.SubscriptionType) {
	for subType := range subTypeChan {
		items := getItemsForSubType(subType)
		log.Debugf("found %v items", len(items))

		db.Where("subscription_type_id=?", subType.ID).Preload("Destination").
			Joins("JOIN destinations ON subscriptions.destination_id = destinations.id").
			Find(&subType.Subscriptions)
		var wg sync.WaitGroup
		for _, sub := range subType.Subscriptions {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sub.LastItem = dispatch(sub, items)

				db.Save(&sub)
			}()

		}
		wg.Wait()
	}
	return
}

func getItemsForSubType(subType database.SubscriptionType) ([]SubscriptionItem) {
	handler := GetSubTypeHandler(subType.Type)
	if handler == nil {
		log.Warnf("Unrecognized sub type %v", subType.Type)
		return []SubscriptionItem{}
	}
	log.Infof("Fetching New items for %v:%v", handler.GetType(), subType.Tags)
	items := handler.GetNewItems(subType.Tags)

	sort.Slice(items, func(i, j int) bool { return items[i].TimeID < items[j].TimeID })

	return items
}
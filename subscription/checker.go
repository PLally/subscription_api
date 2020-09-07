package subscription

// this package contains much of the logic
// checks subscription types for new items
// dispatches those items to their destinations
import (
	"github.com/plally/subscription_api/database"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sort"
	"sync"
)

func makeWorkers(amount int, workerMaker func()) *sync.WaitGroup {
	var wg sync.WaitGroup
	for i := 1; i <= amount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workerMaker()
		}()
	}
	return &wg
}

// gets new items from a subscription type
func CheckOutDatedSubscriptionTypes(db *gorm.DB, workerAmount int) error {
	subTypesToCheck := make(chan database.SubscriptionType)

	waitGroup := makeWorkers(workerAmount, func() {
		checkSubTypesWorker(db, subTypesToCheck)
	})

	var subTypes []database.SubscriptionType
	err := db.FindInBatches(&subTypes, 1000, func(tx *gorm.DB, batch int) error {
		for _, subType := range subTypes {
			subTypesToCheck <- subType
		}
		return nil
	}).Error
	close(subTypesToCheck)
	if err != nil {
		return err
	}
	waitGroup.Wait()
	return nil
}

// does the actual work for CheckOutDatedSubscriptionTypes
func checkSubTypesWorker(db *gorm.DB, subTypeChan chan database.SubscriptionType) {
	for subType := range subTypeChan {
		items := getSubscriptionTypeItems(subType)
		log.Debugf("found %v items", len(items))

		db.Where("subscription_type_id=?", subType.ID).Joins("Destination").
			Find(&subType.Subscriptions)

		var wg sync.WaitGroup
		for _, sub := range subType.Subscriptions {
			wg.Add(1)
			sub := sub
			go func() {
				defer wg.Done()
				if !sub.HasDispatched() && len(items) > 0 {
					items = []SubscriptionItem{
						items[len(items)-1],
					}
				}
				lastItem := dispatch(sub, items)
				sub.LastItem = lastItem
				db.Save(&sub)
			}()
		}
		wg.Wait()
	}
	return
}

func getSubscriptionTypeItems(subType database.SubscriptionType) []SubscriptionItem {
	defer func() {
		if r := recover(); r != nil {
			log.Error("Recovered from a fatal error ", r)
		}
	}()

	typeHandler := GetSubTypeHandler(subType.Type)
	if typeHandler == nil {
		log.Warnf("Unrecognized sub type %v", subType.Type)
		return []SubscriptionItem{}
	}

	log.Infof("Fetching New items for %v", subType.String())
	items := typeHandler.GetNewItems(subType.Tags)

	sort.Slice(items, func(i, j int) bool { return items[i].TimeID < items[j].TimeID })

	for k, item := range items {
		item.Type = subType.Type
		item.Tags = subType.Tags
		items[k] = item
	}
	return items
}

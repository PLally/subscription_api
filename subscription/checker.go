package subscription
// this package contains much of the logic
// checks subscription types for new items
// dispatches those items to their destinations
import (
	"github.com/plally/subscription_api/storage"
	log "github.com/sirupsen/logrus"
	"sort"
	"sync"
)

// the glue that holds my sphagetti together

// gets new items from a subscription type
func CheckOutDatedSubscriptionTypes(db storage.SubscriptionDatabase, max_workers int) {
	subTypeChan := db.SubscriptionType_Get(1000)
	var wg sync.WaitGroup
	for i := 1; i <= max_workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			checkSubTypesWorker(db, subTypeChan)
		}()
	}
	wg.Wait()

}

// does the actual work for CheckOutDatedSubscriptionTypes
func checkSubTypesWorker(db storage.SubscriptionDatabase, subTypeChan chan storage.SubscriptionType) {
	for subType := range subTypeChan {
		handler := GetSubTypeHandler(subType.Type)
		if handler == nil {
			log.Warnf("Unrecognized sub type %v", subType.Type)
			continue
		}
		log.Infof("Fetching New items for %v:%v", handler.GetType(), subType.Tags)
		items := handler.GetNewItems(subType.Tags)

		sort.Slice(items, func(i, j int) bool { return items[i].TimeID < items[j].TimeID })

		var wg sync.WaitGroup
		for dest := range db.Subscription_GetWithDestination_BySubType(subType.ID) {

			wg.Add(1)
			go func() {
				defer wg.Done()
				dest.LastItem = dispatch(dest, items)

				db.Subscription_Update(dest)
			}()

		}
		wg.Wait()
	}
	return
}


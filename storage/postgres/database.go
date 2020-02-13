package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/plally/subscription_api/storage"
	log "github.com/sirupsen/logrus"
)

// creates all necessary subscriptions types and destinations
type SubscriptionDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	*sql.DB
}

func (db *SubscriptionDatabase) SubscriptionType_Get(amount int) chan storage.SubscriptionType {

	subTypeChan := make(chan storage.SubscriptionType)
	go func() {
		rows, err := db.Query(`SELECT * FROM subscription_types`)
		if err != nil {
			log.Error("SubscriptionType_Get():", err)
		}
		for rows.Next() {

			s := storage.SubscriptionType{}
			rows.Scan(&s.ID, &s.Type, &s.Tags)
			subTypeChan <- s
		}
		close(subTypeChan)
	}()
	return subTypeChan
}

func (db *SubscriptionDatabase) SubscriptionType_Create(subscriptionType storage.SubscriptionType) (*storage.SubscriptionType) {
	row := db.QueryRow(
		`INSERT INTO subscription_types (type, tags) VALUES($1,$2)
 				ON CONFLICT (type, tags) DO UPDATE SET id=subscription_types.id
				RETURNING id;`,

		subscriptionType.Type, subscriptionType.Tags,
	)
	row.Scan(&subscriptionType.ID)
	return &subscriptionType
}

func (db *SubscriptionDatabase) Destination_Create(d storage.Destination) *storage.Destination {
	row := db.QueryRow(
		`INSERT INTO destinations (external_identifier, destination_type) VALUES($1, $2)		 
				ON CONFLICT (external_identifier, destination_type) DO UPDATE SET id=destinations.id
				RETURNING id;`,

		d.DestinationType,
		d.ExternalIdentifier,
	)
	row.Scan(&d.ID)
	return  &d
}

func (db *SubscriptionDatabase) Subscription_Create(s storage.Subscription) *storage.Subscription {
	row := db.QueryRow(
		`INSERT INTO subscriptions (subscription_type, destination) VALUES($1, $2) 
			ON CONFLICT (subscription_type, destination) DO UPDATE SET id=subscriptions.id
			RETURNING id`,
		s.SubscriptionTypeID, s.DestinationID,
	)
	row.Scan(&s.ID)
	return &s
}
func (s *SubscriptionDatabase) Subscription_GetWithDestination_BySubType(subType int) chan storage.Subscription {
	channel := make(chan storage.Subscription)
	go func() {
		rows, err := s.Query(`SELECT destinations.id, destinations.external_identifier, destinations.destination_type,
       						subscriptions.id, subscriptions.subscription_type, subscriptions.last_item
							FROM subscriptions JOIN destinations ON subscriptions.destination = destinations.id
							WHERE subscriptions.subscription_type = $1;`, subType)
		if err != nil {
			log.Error("Subscription_GetWithDestination_BySubType():", err)
		}
		for rows.Next() {
			sub := storage.Subscription{Destination: &storage.Destination{}}

			rows.Scan(&sub.Destination.ID, &sub.Destination.ExternalIdentifier, &sub.Destination.DestinationType,
				&sub.ID, &sub.SubscriptionTypeID, &sub.LastItem)
			sub.DestinationID = sub.Destination.ID

			channel <- sub
		}
		close(channel)
	}()
	return channel
}
func (s *SubscriptionDatabase) Subscription_Update(subscription storage.Subscription) {

	_, err := s.Exec(`UPDATE subscriptions SET last_item=$1, destination=$2, subscription_type=$3 WHERE id =$4`,
		subscription.LastItem, subscription.DestinationID, subscription.SubscriptionTypeID, subscription.ID,
	)
	if err != nil {
		log.Error("Subscription_Update():", err)
	}
}
func (s *SubscriptionDatabase) Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Host,
		s.Port,
		s.User,
		s.Password,
		s.Dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	s.DB = db
}

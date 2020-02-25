package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/plally/subscription_api/database"
	_ "github.com/plally/subscription_api/sub_types"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)


var (
	shouldMigrate  = flag.Bool("migrate", false, "Perform database migrations")
)

func main() {
	config, err := readConfig("cmd/conf.json")
	log.SetLevel(log.DebugLevel)
	if err != nil { log.Fatal("readConfig:", err) }
	db := makedb(config.Database)
	if *shouldMigrate {
		fmt.Println("Performing migrations")
		database.Migrate(db)
		os.Exit(0)
	}

	go startSubscriptionPoller(db)

	s := subscriptionServer{db}
	r := mux.NewRouter()

	r.HandleFunc("/subscriptions", s.createHandler(func() interface{}{
		return &database.Subscription{}
	})).
		Methods("POST")

	r.HandleFunc("/destinations", s.createHandler(func() interface{}{
		return &database.Destination{}
	})).
		Methods("POST")

	r.HandleFunc("/subscription_types", s.createHandler(func() interface{}{
		return &database.SubscriptionType{}
	})).
		Methods("POST")


	err = http.ListenAndServe( config.HttpPort, r)
	log.Fatal("ListenAndServer:", err)
}


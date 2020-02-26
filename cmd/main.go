package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/plally/subscription_api/database"
	subTypes "github.com/plally/subscription_api/types"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	shouldMigrate   = flag.Bool("migrate", false, "Perform database migrations")
	startHttp       = flag.Bool("doHttp", true, "Start the http api")
	startBackground = flag.Bool("background", true, "start background tasks")
)

func main() {
	subTypes.RegisterRSS()

	config, err := readConfig("cmd/conf.json")
	if err != nil {
		log.Fatal("readConfig:", err)
	}
	db := makedb(config.Database)

	log.SetLevel(log.DebugLevel)

	if *shouldMigrate {
		fmt.Println("Performing migrations")
		database.Migrate(db)
		fmt.Println("Migrations done")
		os.Exit(0)
	}

	if *startBackground {
		go startSubscriptionPoller(db)
	}

	if *startHttp {
		go runHttp(db, config)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func runHttp(db *gorm.DB, config *config) {
	s := subscriptionServer{db}
	r := mux.NewRouter()

	r.HandleFunc("/subscriptions", s.createHandler(func() interface{} {
		return &database.Subscription{}
	})).
		Methods("POST")

	r.HandleFunc("/destinations", s.createHandler(func() interface{} {
		return &database.Destination{}
	})).
		Methods("POST")

	r.HandleFunc("/subscription_types", s.createHandler(func() interface{} {
		return &database.SubscriptionType{}
	})).
		Methods("POST")

	err := http.ListenAndServe(config.HttpPort, r)
	log.Fatal("ListenAndServer:", err)
}

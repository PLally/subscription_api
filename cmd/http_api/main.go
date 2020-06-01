package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/internal/auth"
	"github.com/plally/subscription_api/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var noauth = flag.Bool("noauth", false, "dont use any authentication")

func makedb() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
	)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	database.Migrate(db)
	return db
}

func main() {
	flag.Parse()

	log.SetLevel(log.DebugLevel)
	viper.SetEnvPrefix("SUB_API")
	viper.SetConfigName("subapi_config")

	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.AddConfigPath("/etc/subscription_api")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	types.RegisterE621()
	types.RegisterRSS()

	r := mux.NewRouter()

	DB := makedb()
	DB.LogMode(true)
	resources(r, database.Destination{}, DB)
	resources(r, database.Subscription{}, DB)
	resources(r, database.SubscriptionType{}, DB).Use(CheckSubscriptionType)
	r.HandleFunc("/subscribe", subscribeHandler(DB))

	if *noauth {
		log.Warn("Starting with no authentication middleware!!!")
	} else {
		r.Use(
			auth.AuthMiddleware(viper.GetString("secret")),
		)
	}

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

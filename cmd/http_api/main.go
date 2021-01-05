package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/internal/auth"
	"github.com/plally/subscription_api/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"strings"
	"time"
)

var noauth = flag.Bool("noauth", true, "dont use any authentication")

func makedb() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
	)
	db, err := gorm.Open(
		postgres.Open(psqlInfo),
		&gorm.Config{Logger: logger.New(log.StandardLogger(), logger.Config{
			SlowThreshold: 0,
			Colorful:      false,
			LogLevel:      logger.Info,
		})},
	)
	if err != nil {
		panic(err)
	}

	database.Migrate(db)
	return db
}

func main() {
	flag.Parse()

	log.SetLevel(log.DebugLevel)

	viper.SetConfigName("subapi_config")

	viper.SetConfigType("yaml")

	viper.AddConfigPath("/etc/subscription_api")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()

	types.RegisterE621()
	types.RegisterRSS()

	r := mux.NewRouter()

	DB := makedb()

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
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

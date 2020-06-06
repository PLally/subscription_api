package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/plally/subscription_api/database"
	_ "github.com/plally/subscription_api/destinations"
	"github.com/plally/subscription_api/subscription"
	"github.com/plally/subscription_api/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func main() {
	log.SetLevel(log.InfoLevel)
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

	db := makedb()

	for {
		time.Sleep(10 * time.Second)
		subscription.CheckOutDatedSubscriptionTypes(db, 100)
	}
}

// TODO get rid of code duplication in http_api and this package
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

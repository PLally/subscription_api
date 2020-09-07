package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/plally/subscription_api/database"
	_ "github.com/plally/subscription_api/destinations"
	"github.com/plally/subscription_api/subscription"
	"github.com/plally/subscription_api/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func main() {
	log.SetLevel(log.InfoLevel)
	viper.SetEnvPrefix("SUB_API")
	viper.SetConfigName("subapi_config")
	viper.SetDefault("dispatcher.delay", "1m")
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.AddConfigPath("/etc/subscription_api")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	types.RegisterE621().StartPostCacheUpdater()
	types.RegisterRSS()

	db := makedb()

	for {
		time.Sleep(viper.GetDuration("dispatcher.delay"))

		err = subscription.CheckOutDatedSubscriptionTypes(db, 100)
		if err != nil {
			log.Error(err)
		}
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

	db, err := gorm.Open(
		postgres.Open(psqlInfo),
		&gorm.Config{Logger: logger.New(log.StandardLogger(), logger.Config{
			SlowThreshold: 0,
			Colorful:      false,
			LogLevel:      logger.Info,
		})})
	if err != nil {
		panic(err)
	}

	database.Migrate(db)
	return db
}

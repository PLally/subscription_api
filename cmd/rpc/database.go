package main

import (
	"fmt"
	"github.com/plally/subscription_api/database"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func connectToDatabase() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_DBNAME"),
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
	log.Info("database connected")
	database.Migrate(db)
	return db
}

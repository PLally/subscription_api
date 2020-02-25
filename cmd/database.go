package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"log"
)

func makedb(config databaseConfig) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Address,
		config.Port,
		config.User,
		config.Password,
		config.DatabaseName,
	)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil { log.Fatal(err) }
	logrus.SetLevel(logrus.DebugLevel)
	db = db.LogMode(true)
	return db
}


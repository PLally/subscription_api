package main

import (
	"github.com/jinzhu/gorm"
	"github.com/plally/subscription_api/subscription"
	"time"
)

func startSubscriptionPoller(db *gorm.DB){
	subscription.CheckOutDatedSubscriptionTypes(db, 10)
	time.Sleep(time.Second*100)
}
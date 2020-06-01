package subscription

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/plally/subscription_api/database"
	"github.com/sirupsen/logrus"
	"log"
	"testing"
)

func makedb() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1",
		"5432",
		"dev",
		"fox",
		"test",
	)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	db = db.LogMode(true)
	database.Migrate(db)
	return db
}

func teardowndb(db *gorm.DB) {
	db.Exec("drop table destinations CASCADE;")
	db.Exec("DROP TABLE subscription_types CASCADE")
	db.Exec("DROP TABLE subscriptions CASCADE")
}
func TestGetItemsForSubTyp(t *testing.T) {
}

package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/plally/subscription_api/database"
	"log"
)

func main() {
	makedb()
}

var shouldMigrate  = flag.Bool("migrate", false, "Perform database migrations")
func makedb() {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1",
		"5432",
		"dev",
		"fox",
		"subscription_dev",
	)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil { log.Fatal(err) }
	if *shouldMigrate {
		database.Migrate(db)
		return
	}

	db.Firs
}

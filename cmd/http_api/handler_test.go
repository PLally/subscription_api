package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/plally/subscription_api/database"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func maketestdb() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1",
		"5432",
		"dev",
		"fox",
		"fox_bot_dev",
	)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	database.Migrate(db)
	return db
}

func teardown(db *gorm.DB) {
	db.Exec("DROP TABLE destinations CASCADE;")
	db.Exec("DROP TABLE subscriptions CASCADE;")
	db.Exec("DROP TABLE subscription_types CASCADE;")
}

func populateDatabase(db *gorm.DB) {
	db.Create(&database.Destination{ExternalIdentifier: "a", DestinationType: "a"})
	db.Create(&database.Destination{ExternalIdentifier: "b", DestinationType: "b"})
	db.Create(&database.Destination{ExternalIdentifier: "c", DestinationType: "d"})
}

func TestSubscriptionServer_CreateDestination(t *testing.T) {
	jsonData, _ := json.Marshal(struct {
		ExternalIdentifier string `json:"external_identifier"`
		DestinationType    string `json:"destination_type"`
	}{
		"test",
		"discord",
	})


	body := bytes.NewReader(jsonData)
	req, _ := http.NewRequest("POST", "http://localhost/destinations", body)

	recorder := httptest.NewRecorder()

	DB := maketestdb()
	createHandler(database.Destination{}, DB)(recorder, req)
	defer teardown(DB)

	result := recorder.Result()

	dest := database.Destination{}
	bodyData, _ := ioutil.ReadAll(result.Body)

	json.Unmarshal(bodyData, &dest)

	if dest.ID < 1 {
		t.Fail()
	}

	if dest.DestinationType != "discord" {
		t.Errorf("dest types do not match %v", dest.DestinationType)
	}

	if dest.ExternalIdentifier != "test" {
		t.Error("identifiers do not match")
	}
}

func TestSubscriptionServer_IndexDestinations(t *testing.T) {
	db := maketestdb()
	defer teardown(db)
	populateDatabase(db)

	req, _ := http.NewRequest("POST", "http://localhost/destinations", nil)

	var destinations []database.Destination
	recorder := httptest.NewRecorder()

	indexHandler(database.Destination{}, db)(recorder, req)

	data, _ := ioutil.ReadAll(recorder.Result().Body)
	json.Unmarshal(data, &destinations)

	if len(destinations) < 2 {
		t.Fail()
	}
}

func TestSubscriptionServer_GetDestination(t *testing.T) {
	db := maketestdb()
	defer teardown(db)
	populateDatabase(db)

	req, _ := http.NewRequest("GET", "http://localhost/destinations/1", nil)

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	recorder := httptest.NewRecorder()
	getHandler(database.Destination{}, db)(recorder, req)

	var dest database.Destination
	data, _ := ioutil.ReadAll(recorder.Result().Body)

	json.Unmarshal(data, &dest)

	if dest.ExternalIdentifier != "a" {
		t.Fail()
	}

}

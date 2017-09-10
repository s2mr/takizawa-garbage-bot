package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/shimokp/takizawa-garbage-bot/manager/config"
)

type databaseManager struct {
	DB *sql.DB
}

var sharedInstance *databaseManager = newDatabaseManager()

func newDatabaseManager() *databaseManager {
	db, err := sql.Open("postgres", config.GetInstance().DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	return &databaseManager{db}
}

func GetInstance() *databaseManager {
	return sharedInstance
}

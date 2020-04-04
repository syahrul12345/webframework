package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {
	isDocker := os.Getenv("IS_DOCKER")
	var dbHost string
	if isDocker != "true" {
		// Running outside docker. Load the env file as usual.
		e := godotenv.Load("../.env") //Load .env file
		if e != nil {
			log.Print(e)
		}
		dbHost = "localhost"
	} else {
		dbHost = os.Getenv("DB_HOST")
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	log.Printf("Attempting to connect to %s\n", dbURI)
	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}) //Database migration
}

//GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

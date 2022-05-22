package database

import (
	"fmt"
	"log"

	"github.com/sanketitnal/gobasicrest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func ConnectPostgresDB() {
	var host string = "localhost"
	var db_username string = "postgres"
	var db_password string = "password"
	var db_name string = "gosimplerestDB"
	var db_port int16 = 5432
	var sslmode string = "disable" // enable or disable

	// https://gorm.io/docs/connecting_to_the_database.html
	database_conn_url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, db_username, db_password, db_name, db_port, sslmode)
	DB, err := gorm.Open(postgres.Open(database_conn_url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error: Cannot connect to POSTGRES Database")
	}
	PostgresDB = DB
	PostgresDB.AutoMigrate(&models.User{})
}

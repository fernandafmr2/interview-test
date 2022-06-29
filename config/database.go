package config

import (
	"fmt"
	"interview-test/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Setup for creating a new connection to database
func SetUpDatabaseConnection(c Config) *gorm.DB {

	dbUser := c.Get("DB_USER")
	dbPass := c.Get("DB_PASS")
	dbHost := c.Get("DB_HOST")
	dbName := c.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	// give a modul in here
	db.AutoMigrate(&entity.Animal{})

	return db
}

// close database connection
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}

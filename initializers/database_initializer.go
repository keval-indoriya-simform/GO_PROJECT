package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	// TO CREATE SINGLETON OBJECT
	once sync.Once

	// TO RETURN INSTANCE OF DB
	instance *gorm.DB
)

func DB() *gorm.DB {
	once.Do(func() {
		instance = initializeDB()
	})
	return instance
}

func initializeDB() *gorm.DB {
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, OpenPostgresDBError := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		FullSaveAssociations: true,
	})
	if OpenPostgresDBError != nil {
		log.Fatal(OpenPostgresDBError)
	}
	return db
}

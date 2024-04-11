package initializers

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dbPort := strings.ReplaceAll(os.Getenv("DB_PORT"), `"`, ``)
	dbHost := strings.ReplaceAll(os.Getenv("DB_HOST"), `"`, ``)
	dbUser := strings.ReplaceAll(os.Getenv("DB_USER"), `"`, ``)
	dbPassword := strings.ReplaceAll(os.Getenv("DB_PASSWORD"), `"`, ``)
	dbName := strings.ReplaceAll(os.Getenv("DB_NAME"), `"`, ``)
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, OpenPostgresDBError := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		FullSaveAssociations: true,
	})
	if OpenPostgresDBError != nil {
		log.Fatal(OpenPostgresDBError)
	}
	return db
}

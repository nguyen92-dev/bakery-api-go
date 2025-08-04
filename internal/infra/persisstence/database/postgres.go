package database

import (
	"bakery-api/configs"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *configs.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.SSLMode,
		cfg.Database.TimeZone)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Minute)

	log.Println("Connected to database successfully")
	return nil
}

func GetDbClient() *gorm.DB {
	if dbClient == nil {
		log.Fatal("Database client is not initialized. Call InitDb first.")
	}
	return dbClient
}

func CloseDb() {
	if dbClient == nil {
		log.Println("Database client is already nil.")
		return
	}
	sqlDB, err := dbClient.DB()
	if err != nil {
		log.Printf("Error getting database instance: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	} else {
		log.Println("Database connection closed successfully")
	}
	dbClient = nil
}

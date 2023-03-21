package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDBConfigFromEnv() DatabaseConfig {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	return DatabaseConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}
}

func NewDatabaseConnection(config DatabaseConfig) (*gorm.DB, error) {
	dbUser := config.User
	dbPassword := config.Password
	dbHost := config.Host
	dbName := config.DBName
	dbPort := config.Port

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

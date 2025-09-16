package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	// Load environment variables from .env file if present (non-fatal)
	_ = godotenv.Load()

	// // Ensure Viper is initialized
	// LoadConfig()

	// Prefer environment variables; otherwise fall back to config.yaml values
	host := firstNonEmpty(os.Getenv("DB_HOST"), viper.GetString("database.host"), "localhost")
	user := firstNonEmpty(os.Getenv("DB_USER"), viper.GetString("database.user"), "admin")
	password := firstNonEmpty(os.Getenv("DB_PASSWORD"), viper.GetString("database.password"), "password")
	dbname := firstNonEmpty(os.Getenv("DB_NAME"), viper.GetString("database.name"), "employees")
	port := firstNonEmpty(os.Getenv("DB_PORT"), fmt.Sprintf("%v", viper.Get("database.port")), "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	sqlDB, err := db.DB()
	if err == nil {
		// Reasonable defaults
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
		_ = sqlDB.Ping()
	}
	return db, nil
}

func NewDatabaseConnection() *gorm.DB {
	db, err := ConnectDatabase()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	return db
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

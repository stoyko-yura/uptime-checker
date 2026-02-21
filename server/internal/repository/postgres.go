package repository

import (
	"fmt"
	"os"
	"uptime-checker/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresPool() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=user password=password dbname=uptime_db port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	fmt.Println("connected to database")

	err = db.AutoMigrate(&model.Site{}, &model.Check{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate sites: %v", err)
	}
	fmt.Println("migrated sites")

	return db, nil
}

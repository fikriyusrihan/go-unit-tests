package db

import (
	"fmt"
	"go-product/config"
	"go-product/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func NewPostgresDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
		config.C.Database.SSLMode,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Debug().AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		return nil, err
	}

	log.Println("database connection successfully created")
	return db, nil
}

func GetDBInstance() *gorm.DB {
	return db
}

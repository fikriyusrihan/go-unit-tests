package db

import (
	"fmt"
	"go-product/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "postgres"
	dbname   = "db-go-sql"
)

func NewPostgresDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Debug().AutoMigrate(&domain.User{}, &domain.Product{})
	if err != nil {
		return nil, err
	}

	log.Println("database connection successfully created")
	return db, nil
}

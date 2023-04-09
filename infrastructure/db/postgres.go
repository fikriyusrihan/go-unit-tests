package db

import (
	"fmt"
	"go-product/domain/config"
	"go-product/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func NewPostgresDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.C.Database.Host,
		config.C.Database.Username,
		config.C.Database.Password,
		config.C.Database.DBName,
		config.C.Database.Port,
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

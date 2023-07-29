package database

import (
	"clockwork-server/domain/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Connect() *gorm.DB
}

type database struct {
}

func NewDatabase() Database {
	return &database{}
}

func (d database) Connect() *gorm.DB {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	var db *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, databaseName, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(
		&model.AttributeItem{},
		&model.CartItemAttributeItem{},
		&model.User{},
		&model.Product{},
		&model.Order{},
		&model.Cart{},
		&model.CartItem{},
		&model.Category{},
		&model.Customer{},
		&model.Image{},
		&model.Inventory{},
		&model.Address{},
		&model.Organization{},
		&model.Location{},
		&model.Voucher{},
	)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
		return db
	}
}

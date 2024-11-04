package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Posts struct {
	gorm.Model
	Name      string
	Author    string
	Text      string
	Thumbnail string
}

type Users struct {
	gorm.Model
	Username string
	Password string
}

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
}

func RunMigrations() error {
	ConnectDB()
	err := DB.AutoMigrate(&Posts{}, &Users{})
	if err != nil {
		return err
	}
	InsertData()
	return nil
}

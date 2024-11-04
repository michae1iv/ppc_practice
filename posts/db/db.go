package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Template string
	AuthorID int
	Users    Users `gorm:"foreignKey:AuthorID"`
}

type Users struct {
	gorm.Model
	Username string
	Password string
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./posts.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&Posts{}, &Users{})
	if err != nil {
		return err
	}
	return nil
}

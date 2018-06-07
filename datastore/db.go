package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tomrom95/go-rush/datastore/models"
)

type DataStore interface {
	Connect() error
	Close()
}

type GormDataStore struct {
	DB *gorm.DB
}

func NewDataStore() *GormDataStore {
	return &GormDataStore{}
}

func (ds *GormDataStore) Connect() error {
	db, err := gorm.Open(
		"postgres",
		getPostgresConfig(),
	)
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Rushee{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Reply{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(&models.Attendance{})
	db.AutoMigrate(&models.FeedEvent{})
	ds.DB = db
	return nil
}

func (ds *GormDataStore) Close() {
	if ds.DB != nil {
		ds.DB.Close()
	}
}

func getPostgresConfig() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		"127.0.0.1",
		5432,
		"tomrom95",
		"tomrom95",
	)
}

package services_test

import (
  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connectToTestDB() *gorm.DB {
  db, err := gorm.Open(
    "postgres",
    "host=127.0.0.1 port=5432 user=test dbname=test sslmode=disable",
  )
  if err != nil {
    panic(err)
  }
  return db
}

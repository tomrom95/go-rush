package services_test

import (
  "testing"

  "github.com/jinzhu/gorm"
  "github.com/stretchr/testify/assert"
  "github.com/tomrom95/go-rush/datastore/models"
  "github.com/tomrom95/go-rush/datastore/services"
)

func setup() (*gorm.DB, *services.GormRusheeService) {
  db := connectToTestDB()
  db.AutoMigrate(&models.Rushee{})
  db.Delete(&models.Rushee{}) // delete all
  return db, services.NewRusheeService(db)
}

func teardown(db *gorm.DB) {
  db.Delete(&models.Rushee{})
  db.Close()
}

func TestCreateRushee(t *testing.T) {
  db, service := setup()
  defer teardown(db)

  newRushee := &models.Rushee{ID: 1, Name: "test name"}
  err := service.Create(newRushee)
  assert.NoError(t, err)
  rushee, err := service.Get(1)
  assert.NoError(t, err)
  assert.Equal(t, rushee.Name, newRushee.Name)
  assert.Equal(t, rushee.ID, newRushee.ID)

  service.Create(&models.Rushee{ID: 2, Name: "test2"})
  service.Create(&models.Rushee{ID: 3, Name: "test3"})
  rushees, err := service.GetAll()
  assert.NoError(t, err)
  assert.Equal(t, len(rushees), 3)
}

func TestDeleteRushee(t *testing.T) {
  db, service := setup()
  defer teardown(db)

  service.Create(&models.Rushee{ID: 1, Name: "test2"})
  service.Create(&models.Rushee{ID: 2, Name: "test2"})
  service.Create(&models.Rushee{ID: 3, Name: "test3"})
  rushees, err := service.GetAll()
  assert.NoError(t, err)
  assert.Equal(t, len(rushees), 3)
  service.Delete(1)
  rusheesAfter, err := service.GetAll()
  assert.NoError(t, err)
  assert.Equal(t, len(rusheesAfter), 2)
  service.Delete(2)
  service.Delete(3)
  rusheesLast, err := service.GetAll()
  assert.NoError(t, err)
  assert.Equal(t, len(rusheesLast), 0)
}

func TestUpdateRushee(t *testing.T) {
  db, service := setup()
  defer teardown(db)

  service.Create(&models.Rushee{ID: 1, Name: "test2"})
  rushee, err := service.Get(1)
  assert.NoError(t, err)
  assert.Equal(t, rushee.Name, "test2")

  err = service.Update(&models.Rushee{ID: 1, Name: "test10"})
  assert.NoError(t, err)
  rusheeUpdated, err := service.Get(1)
  assert.NoError(t, err)
  assert.Equal(t, rusheeUpdated.Name, "test10")
}

package services

import (
	"github.com/jinzhu/gorm"
	"github.com/tomrom95/go-rush/datastore/models"
)

type RusheeService interface {
	Create(rushee *models.Rushee) error
	GetAll(id uint) ([]*models.Rushee, error)
	Get(id uint) (*models.Rushee, error)
	Update(rushee *models.Rushee) error
	Delete(id uint) error
}

type GormRusheeService struct {
	db *gorm.DB
}

func NewRusheeService(db *gorm.DB) *GormRusheeService {
	return &GormRusheeService{db: db}
}

func (service *GormRusheeService) Get(id uint) (*models.Rushee, error) {
	var rushee models.Rushee
	err := service.db.First(&rushee, id).Error
	if err != nil {
		return nil, err
	}
	return &rushee, nil
}

func (service *GormRusheeService) GetAll() ([]*models.Rushee, error) {
	var rushees []*models.Rushee
	err := service.db.Find(&rushees).Error
	if err != nil {
		return []*models.Rushee{}, err
	}
	return rushees, nil
}

func (service *GormRusheeService) Create(rushee *models.Rushee) error {
	return service.db.Create(rushee).Error
}

func (service *GormRusheeService) Update(rushee *models.Rushee) error {
	return service.db.Save(rushee).Error
}

func (service *GormRusheeService) Delete(id uint) error {
	return service.db.Delete(&models.Rushee{ID: id}).Error
}

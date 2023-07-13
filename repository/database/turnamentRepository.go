package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type TurnamentRepository interface {
	GetTurnament() (turnament []models.Turnament, err error)
	GetTurnamentById(id uint) (turnament *models.Turnament, err error)
	CreateTurnament(turnament *models.Turnament) error
	DeleteTurnament(turnament *models.Turnament) error
	UpdateTurnament(turnament *models.Turnament) error
	CountTurnament() (res int64)
}

type turnamentRepository struct {
	db *gorm.DB
}

func NewTurnamentRepository(db *gorm.DB) *turnamentRepository {
	return &turnamentRepository{db}
}

func (t *turnamentRepository) GetTurnament() (turnament []models.Turnament, err error) {

	if err := config.DB.Preload("Participation.Player").Preload("Match").Find(&turnament).Error; err != nil {
		return nil, err
	}

	return turnament, nil
}

func (t *turnamentRepository) GetTurnamentById(id uint) (turnament *models.Turnament, err error) {

	if err := config.DB.Preload("Participation.Player").Preload("Match.Player_1").Preload("Match.Player_2").Where("id = ?", id).First(&turnament).Error; err != nil {
		return nil, err
	}

	return turnament, nil
}

func (t *turnamentRepository) CountTurnament() (res int64) {
	res = 0
	turnament := []models.Turnament{}

	if err := config.DB.Model(&turnament).Count(&res).Error; err != nil {
		return 0
	}

	return res
}

func (t *turnamentRepository) CreateTurnament(turnament *models.Turnament) error {
	if err := config.DB.Create(&turnament).Error; err != nil {
		return err
	}

	return nil
}

func (t *turnamentRepository) UpdateTurnament(turnament *models.Turnament) error {

	if err := config.DB.Save(&turnament).Error; err != nil {
		return err
	}

	return nil
}

func (t *turnamentRepository) DeleteTurnament(turnament *models.Turnament) error {

	if err := config.DB.Delete(&turnament).Error; err != nil {
		return err
	}

	return nil
}

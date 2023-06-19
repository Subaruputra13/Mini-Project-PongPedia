package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type MatchRepository interface {
	GetMatch() ([]models.Match, error)
	GetMatchById(id uint) (*models.Match, error)
	GetParticipationByTurnamentIdAndPlayerId(idT, idP uint) (*models.Participation, error)
	GetMatchByTurnamentIdAndPlayerId(idT, id1, id2 uint) (*models.Match, error)
	CreateMatch(match *models.Match) error
	UpdateMatch(match *models.Match) error
	DeleteMatch(match *models.Match) error
	CountMatch() (res int64)
}

type matchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) MatchRepository {
	return &matchRepository{db}
}

func (m *matchRepository) GetMatch() ([]models.Match, error) {
	match := []models.Match{}

	err := config.DB.Preload("Turnament").Find(&match).Error
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (m *matchRepository) GetMatchById(id uint) (*models.Match, error) {
	match := models.Match{}

	err := config.DB.Preload("Turnament").Where("id = ?", id).First(&match).Error
	if err != nil {
		return nil, err
	}

	return &match, nil
}

func (m *matchRepository) CountMatch() (res int64) {
	res = 0
	match := []models.Match{}

	if err := config.DB.Model(&match).Count(&res).Error; err != nil {
		return 0
	}

	return res
}

func (m *matchRepository) CreateMatch(match *models.Match) error {
	err := config.DB.Preload("Turnament").Create(&match).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *matchRepository) UpdateMatch(match *models.Match) error {
	err := config.DB.Preload("Turnament").Save(&match).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *matchRepository) DeleteMatch(match *models.Match) error {
	err := config.DB.Delete(&match).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *matchRepository) GetParticipationByTurnamentIdAndPlayerId(idT, idP uint) (*models.Participation, error) {
	participation := models.Participation{}

	err := config.DB.Where("turnament_id = ? AND player_id = ?", idT, idP).First(&participation).Error
	if err != nil {
		return nil, err
	}

	return &participation, nil
}

func (m *matchRepository) GetMatchByTurnamentIdAndPlayerId(idT, id1, id2 uint) (*models.Match, error) {
	match := models.Match{}

	err := config.DB.Where("turnament_id = ? AND player_1 = ? AND player_2 = ?", idT, id1, id2).First(&match).Error
	if err != nil {
		return nil, err
	}

	return &match, nil
}

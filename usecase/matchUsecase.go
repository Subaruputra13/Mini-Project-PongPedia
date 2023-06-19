package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"errors"
	"time"
)

type MatchUsecase interface {
	GetMatch() ([]models.Match, error)
	GetMatchById(id uint) (*models.Match, error)
	CreateMatch(req *payload.CreateMatchRequest) (match *models.Match, err error)
	UpdateMatch(match *models.Match) error
}

type matchUsecase struct {
	matchRepository         database.MatchRepository
	participationRepository database.ParticipationRepository
}

func NewMatchUsecase(
	matchRepository database.MatchRepository,
	participationRepository database.ParticipationRepository,
) *matchUsecase {
	return &matchUsecase{matchRepository, participationRepository}
}

func (m *matchUsecase) GetMatch() ([]models.Match, error) {
	match, err := m.matchRepository.GetMatch()
	if err != nil {
		return nil, errors.New("failed to get match")
	}

	return match, nil
}

func (m *matchUsecase) GetMatchById(id uint) (*models.Match, error) {
	match, err := m.matchRepository.GetMatchById(id)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (m *matchUsecase) CreateMatch(req *payload.CreateMatchRequest) (match *models.Match, err error) {
	matchDate, err := time.Parse("02/01/2006", req.MatchDate)
	if err != nil {
		return nil, errors.New("Failed to parse birthdate")
	}

	if matchDate.Before(time.Now().AddDate(0, 0, -1)) {
		return nil, errors.New("Start date must be today")
	}

	matchReq := &models.Match{
		MatchName:      req.MatchName,
		MatchDate:      &matchDate,
		Player_1:       req.Player_1,
		Player_2:       req.Player_2,
		Player_1_Score: req.Player_1_Score,
		Player_2_Score: req.Player_2_Score,
		TurnamentID:    req.TurnamentID,
	}

	// check participation in turnament
	_, err = m.matchRepository.GetParticipationByTurnamentIdAndPlayerId(req.TurnamentID, req.Player_1)
	if err != nil {
		return nil, errors.New("Player not participate in this turnament")
	}

	_, err = m.matchRepository.GetParticipationByTurnamentIdAndPlayerId(req.TurnamentID, req.Player_2)
	if err != nil {
		return nil, errors.New("Player not participate in this turnament")
	}

	// check if match already exist
	_, err = m.matchRepository.GetMatchByTurnamentIdAndPlayerId(req.TurnamentID, req.Player_1, req.Player_2)
	if err == nil {
		return nil, errors.New("Match already exist")
	}

	err = m.matchRepository.CreateMatch(matchReq)
	if err != nil {
		return nil, errors.New("Failed to create match")
	}

	return matchReq, nil
}

func (m *matchUsecase) UpdateMatch(match *models.Match) error {
	err := m.matchRepository.UpdateMatch(match)
	if err != nil {
		return err
	}

	return nil
}

package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"errors"
	"time"

	"github.com/labstack/echo"
)

type TurnamentUsecase interface {
	GetTurnament() ([]payload.TurnamentResponse, error)
	GetTurnamentById(id uint) (turnament *models.Turnament, err error)
	CreateTurnament(req *payload.TurnamentRequest) (res payload.TurnamentResponse, err error)
	RegisterTurnament(id uint, req *payload.RegisterTurnamentRequest) error
}

type turnamentUsecase struct {
	turnamentRepository database.TurnamentRepository
	playerRepository    database.PlayerRespository
	userReposistory     database.UserRepository
	particapationRepo   database.ParticipationRepository
}

func NewTurnamentUsecase(
	turnamentRepository database.TurnamentRepository,
	playerRepository database.PlayerRespository,
	userReposistory database.UserRepository,
	participationRepo database.ParticipationRepository,
) TurnamentUsecase {
	return &turnamentUsecase{
		turnamentRepository,
		playerRepository,
		userReposistory,
		participationRepo,
	}
}

func (t *turnamentUsecase) GetTurnament() ([]payload.TurnamentResponse, error) {
	turnament, err := t.turnamentRepository.GetTurnament()
	if err != nil {
		return nil, errors.New("Failed to get turnament")
	}

	res := []payload.TurnamentResponse{}
	for _, v := range turnament {
		res = append(res, payload.TurnamentResponse{
			ID:        v.ID,
			Name:      v.Name,
			StartDate: v.StartDate,
			EndDate:   v.EndDate,
			Location:  v.Location,
			Slot:      v.Slot,
		})
	}
	return res, nil
}

func (t *turnamentUsecase) GetTurnamentById(id uint) (turnament *models.Turnament, err error) {
	// Check Turnament ID
	turnament, err = t.turnamentRepository.GetTurnamentById(id)
	if err != nil {
		echo.NewHTTPError(400, err.Error())
		return
	}

	return turnament, nil
}

func (t *turnamentUsecase) CreateTurnament(req *payload.TurnamentRequest) (res payload.TurnamentResponse, err error) {
	startDate, err := time.Parse("02/01/2006", req.StartDate)
	if err != nil {
		return res, errors.New("Failed to parse start date")
	}

	endDate, err := time.Parse("02/01/2006", req.EndDate)
	if err != nil {
		return res, errors.New("Failed to parse end date")
	}

	if startDate.Before(time.Now().AddDate(0, 0, -1)) {
		return res, errors.New("Start date must be after today")
	}

	if startDate == endDate {
		return res, errors.New("Start date and end date must be different")
	}

	if startDate.After(endDate) {
		return res, errors.New("Start date must be before end date")
	}

	turnamentReq := &models.Turnament{
		Name:      req.Name,
		StartDate: &startDate,
		EndDate:   &endDate,
		Location:  req.Location,
		Slot:      req.Slot,
	}

	err = t.turnamentRepository.CreateTurnament(turnamentReq)
	if err != nil {
		errors.New("Failed to create turnament")
		return
	}

	res = payload.TurnamentResponse{
		ID:        turnamentReq.ID,
		Name:      turnamentReq.Name,
		StartDate: turnamentReq.StartDate,
		EndDate:   turnamentReq.EndDate,
		Location:  turnamentReq.Location,
		Slot:      turnamentReq.Slot,
	}

	return
}

func (t *turnamentUsecase) RegisterTurnament(id uint, req *payload.RegisterTurnamentRequest) error {

	player, err := t.playerRepository.GetPlayerUserId(id)
	if err != nil {
		return errors.New("fill your player profile first")
	}

	regisReq := &models.Participation{
		PlayerID:    uint(player.ID),
		TurnamentID: uint(req.TurnamentID),
	}

	// Check slot availability
	turnament, err := t.turnamentRepository.GetTurnamentById(regisReq.TurnamentID)
	if err != nil {
		return errors.New("Turnament not found")
	}

	if turnament.Slot == 0 {
		return errors.New("Turnament slot is full")
	}

	// Check if user already registered
	err = t.particapationRepo.CheckPartisipasion(regisReq)
	if err == nil {
		return errors.New("Player already registered")
	}

	err = t.particapationRepo.RegisterTurnament(regisReq)
	if err != nil {
		return err
	}

	// Update slot
	turnament.Slot = turnament.Slot - 1

	err = t.turnamentRepository.UpdateTurnament(turnament)
	if err != nil {
		return err
	}

	return nil
}

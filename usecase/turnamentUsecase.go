package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type TurnamentUsecase interface {
	GetTurnament() ([]models.Turnament, error)
	GetTurnamentById(id int) (*models.Turnament, error)
	CreateTurnament(req *payload.TurnamentRequest) (res payload.TurnamentRequest, err error)
	RegisterTurnament(id int, req *payload.RegisterTurnamentRequest) error
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

func (t *turnamentUsecase) GetTurnament() ([]models.Turnament, error) {
	turnament, err := t.turnamentRepository.GetTurnament()
	if err != nil {
		return nil, err
	}

	return turnament, nil
}

func (t *turnamentUsecase) GetTurnamentById(id int) (*models.Turnament, error) {
	turnament, err := t.turnamentRepository.GetTurnamentById(id)
	if err != nil {
		return nil, err
	}

	return turnament, nil
}

func (t *turnamentUsecase) CreateTurnament(req *payload.TurnamentRequest) (res payload.TurnamentRequest, err error) {
	turnamentReq := &models.Turnament{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Location:  req.Location,
	}

	err = t.turnamentRepository.CreateTurnament(turnamentReq)
	if err != nil {
		echo.NewHTTPError(400, err.Error())
		return
	}

	res = payload.TurnamentRequest{
		Name:      turnamentReq.Name,
		StartDate: turnamentReq.StartDate,
		EndDate:   turnamentReq.EndDate,
		Location:  turnamentReq.Location,
	}

	return
}

func (t *turnamentUsecase) RegisterTurnament(id int, req *payload.RegisterTurnamentRequest) error {

	player, err := t.playerRepository.GetPlayerId(id)

	if err != nil {
		return echo.NewHTTPError(400, "Fill your player profile first")
	}

	regisReq := &models.Participation{
		PlayerID:    int(player.ID),
		TurnamentID: req.TurnamentID,
	}

	// Check if user already registered
	err = t.particapationRepo.CheckPartisipasion(regisReq)

	if err == nil {
		return echo.NewHTTPError(400, "You already registered")
	}

	err = t.particapationRepo.RegisterTurnament(regisReq)

	if err != nil {
		return err
	}

	return nil
}

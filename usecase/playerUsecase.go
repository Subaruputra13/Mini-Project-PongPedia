package usecase

import (
	"PongPedia/models"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type PlayerUsecase interface {
	CreatePlayer(id int, player *models.Player) (*models.User, error)
	UpdatePlayer(id int, player *models.Player) (*models.User, error)
	GetPlayer(id int) (*models.Player, error)
<<<<<<< Updated upstream
=======
	UpdatePlayer(id int, req *payload.CreateUpdatePlayerRequest) error
>>>>>>> Stashed changes
}

type playerUsecase struct {
	playerRespository database.PlayerRespository
}

func NewPlayerUsecase(playerRespository database.PlayerRespository) *playerUsecase {
	return &playerUsecase{playerRespository}
}

func (p *playerUsecase) GetPlayer(id int) (*models.Player, error) {

	user, err := p.playerRespository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to read token")
	}

	player, err := p.playerRespository.GetPlayerWithCookie(int(user.ID))

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to get player")
	}

	return player, nil
}

<<<<<<< Updated upstream
// Logic for update player
func (p *playerUsecase) CreatePlayer(id int, player *models.Player) (*models.User, error) {

	user, err := p.playerRespository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to read token")
	}

	if err := p.playerRespository.CreatePlayerWithCookie(id, player); err != nil {
		return nil, echo.NewHTTPError(400, "Failed to create player")
	}

	return user, nil
}

// Logic for update player
func (p *playerUsecase) UpdatePlayer(id int, player *models.Player) (*models.User, error) {

	user, err := p.playerRespository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to read token")
	}

	if err := p.playerRespository.UpdatePlayerWithCookie(id, player); err != nil {
		return nil, echo.NewHTTPError(400, "Failed to create player")
	}

	return user, nil
=======
// // Logic for Create and update player
// func (p *playerUsecase) CreatePlayer(id int, req *payload.CreateUpdatePlayerRequest) (res payload.PlayerResponse, err error) {
// 	userReq := &models.Player{
// 		Name:      req.Name,
// 		Age:       req.Age,
// 		BirthDate: req.BirthDate,
// 		Gender:    req.Gender,
// 		UserID:    id,
// 	}

// 	err = p.playerRespository.CreatePlayer(userReq)

// 	if err != nil {
// 		echo.NewHTTPError(400, "Failed to create player")
// 		return
// 	}

// 	res = payload.PlayerResponse{
// 		ID:        int(userReq.ID),
// 		Name:      userReq.Name,
// 		Age:       userReq.Age,
// 		BirthDate: userReq.BirthDate,
// 		Gender:    userReq.Gender,
// 		UserID:    userReq.UserID,
// 	}

// 	return res, nil
// }

func (p *playerUsecase) UpdatePlayer(id int, req *payload.CreateUpdatePlayerRequest) error {

	player, err := p.playerRespository.GetPlayerId(id)

	if err == nil {
		player.Name = req.Name
		player.Age = req.Age
		player.BirthDate = req.BirthDate
		player.Gender = req.Gender

		err = p.playerRespository.UpdatePlayer(player)
		if err != nil {
			return echo.NewHTTPError(400, "Failed to update player")
		}
	} else {
		userReq := &models.Player{
			Name:      req.Name,
			Age:       req.Age,
			BirthDate: req.BirthDate,
			Gender:    req.Gender,
			UserID:    id,
		}

		err = p.playerRespository.UpdatePlayer(userReq)
		if err != nil {
			return echo.NewHTTPError(400, "Failed to update player")
		}
	}

	return nil
>>>>>>> Stashed changes
}

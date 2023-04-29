package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name      string  `json:"name" form:"name" gorm:"unique;not null"`
	Age       int     `json:"age" form:"age"`
	BirthDate string  `json:"birth_date" form:"birth_date"`
	Gender    string  `json:"gender" form:"gender"`
	UserID    int     `json:"user_id" form:"user_id" gorm:"unique;not null"`
	Matchs    []Match `gorm:"foreignKey:PlayerID"`
}

// type PlayerResponse struct {
// 	gorm.Model
// 	FullName  string    `json:"fullname" form:"fullname"`
// 	Age       uint      `json:"age" form:"age"`
// 	BirthDate time.Time `json:"birth_date" form:"birth_date"`
// 	Gender    string    `json:"gender" form:"gender"`
// 	UserID    uint      `json:"-" form:"user_id"`
// }

// func (PlayerResponse) TableName() string {
// 	return "players"
// }

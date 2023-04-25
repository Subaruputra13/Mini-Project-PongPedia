package models

import "gorm.io/gorm"

type Turnament struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Date  string `json:"date" form:"date"`
	Type  string `json:"type" form:"type"`
	Match []Match
}

package models

import (
	"github.com/jinzhu/gorm"
)

//Pet model
type Pet struct {
	gorm.Model
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Photo string `json:"photo"`
}

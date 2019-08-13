package models

import (
	"fmt"

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

//GetPets will get all pets
func GetPets() []Pet {

	pets := make([]Pet, 0)
	err := GetDB().Table("pets").Find(&pets).Error
	if err != nil {
		fmt.Println(err)

		return nil
	}

	return pets
}

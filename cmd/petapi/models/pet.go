package models

//Pet model
type Pet struct {
	Id    int    `gorm:"primary_key" "AUTO_INCREMENT"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Photo string `json:"photo"`
}

// //GetPets will get all pets
// func GetPets() []Pet {

// 	pets := make([]Pet, 0)
// 	err := GetDB().Table("pets").Find(&pets).Error
// 	if err != nil {
// 		fmt.Println(err)

// 		return nil
// 	}

// 	return pets
// }

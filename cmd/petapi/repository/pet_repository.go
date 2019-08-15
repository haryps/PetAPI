package repository

import (
	m "github.com/PetAPI/cmd/petapi/models"
)

//IRepository interface
type IRepository interface {
	GetPets() ([]m.Pet, error)
	Register(pet m.Pet) error
	GetPetById(id int) (m.Pet, error)
	Update(id int, pet m.Pet) error
	UploadPetImage(id int, image string) error
	Delete(id int) error
}

type PetRepository struct {
}

func (repo *PetRepository) GetPets() ([]m.Pet, error) {
	allPets := []m.Pet{}
	err := m.Db.Find(&allPets)

	return allPets, err.Error
}

func (repo *PetRepository) Register(pet m.Pet) error {
	err := m.Db.Create(pet)

	return err.Error
}

func (repo *PetRepository) GetPetById(id int) (m.Pet, error) {
	pet := m.Pet{}
	err := m.Db.First(&pet, id)

	return pet, err.Error
}

func (repo *PetRepository) Update(id int, pet m.Pet) error {

	existingPet := m.Pet{}
	m.Db.First(&existingPet, id)

	existingPet.Name = pet.Name
	existingPet.Age = pet.Age
	err := m.Db.Save(existingPet)

	return err.Error
}

func (repo *PetRepository) UploadPetImage(id int, image string) error {

	existingPet := m.Pet{}
	m.Db.First(&existingPet, id)

	existingPet.Photo = image
	err := m.Db.Save(&existingPet)

	return err.Error
}

func (repo *PetRepository) Delete(id int) error {
	err := m.Db.Unscoped().Where("id=?", id).Delete(&m.Pet{})

	return err.Error
}

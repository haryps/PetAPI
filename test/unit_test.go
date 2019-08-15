package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	model "github.com/PetAPI/cmd/petapi/models"
)

func TestRegisterPet(t *testing.T) {
	repoMock := PetRepoMock{}
	input := model.Pet{Name: "Felix", Age: 2}

	repoMock.On("Register", mock.MatchedBy(func(pet model.Pet) bool {
		assert.Equal(t, input.Name, pet.Name)
		assert.Equal(t, input.Age, pet.Age)
		return true
	})).Return(nil)

	err := repoMock.Register(input)
	assert.NoError(t, err)
}

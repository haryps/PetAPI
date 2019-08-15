package test

import (
	model "github.com/PetAPI/cmd/petapi/models"
	"github.com/stretchr/testify/mock"
)

type PetRepoMock struct {
	mock.Mock
}

func (p *PetRepoMock) Register(pet model.Pet) error {
	args := p.Called(pet)
	return args.Error(0)
}

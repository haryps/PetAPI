package controllers

import (
	"net/http"

	helper "github.com/PetAPI/cmd/petapi/helpers"
	"github.com/PetAPI/cmd/petapi/models"
)

//GetPets will Get all the pets
var GetPets = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetPets()
	resp := helper.Message(true, "success")
	resp["data"] = data
	helper.Respond(w, resp)
}

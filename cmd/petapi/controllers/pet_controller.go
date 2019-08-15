package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	u "github.com/PetAPI/cmd/petapi/helpers"
	"github.com/PetAPI/cmd/petapi/models"
	repo "github.com/PetAPI/cmd/petapi/repository"
)

type PetController struct {
	petRepo repo.PetRepository
}

//RegisterPet will create a new pet
func (petCtrl *PetController) RegisterPet(w http.ResponseWriter, r *http.Request) {

	var err error

	pet := models.Pet{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &pet)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	err = petCtrl.petRepo.Register(pet)
	resp := u.Message(true, "success")
	u.Respond(w, resp)
}

//GetPetById will get a pet by its id
func (petCtrl *PetController) GetPetById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	data, err := petCtrl.petRepo.GetPetById(id)
	if err != nil {
		panic(err)
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func (petCtrl *PetController) UpdatePetById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	pet := models.Pet{}
	err = json.Unmarshal(reqBody, &pet)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	err = petCtrl.petRepo.Update(id, pet)
	if err != nil {
		panic(err)
	}

	resp := u.Message(true, "success")
	u.Respond(w, resp)
}

func (petCtrl *PetController) UploadPetImageById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = petCtrl.petRepo.UploadPetImage(id, handler.Filename)
	if err != nil {
		panic(err)
	}

	resp := u.Message(true, "success")
	u.Respond(w, resp)
}

func (petCtrl *PetController) DeletePet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	if err := petCtrl.petRepo.Delete(id); err != nil {
		panic(err)
	}
	resp := u.Message(true, "success")
	u.Respond(w, resp)
}

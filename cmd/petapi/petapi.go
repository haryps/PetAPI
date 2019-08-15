package main

import (
	"fmt"
	"net/http"
	"os"

	pCtrl "github.com/PetAPI/cmd/petapi/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	petController := new(pCtrl.PetController)
	router.HandleFunc("/pet", petController.RegisterPet).Methods("POST")
	router.HandleFunc("/pet/{id}", petController.GetPetById).Methods("GET")
	router.HandleFunc("/pet/{id}", petController.UpdatePetById).Methods("PUT")
	router.HandleFunc("/pet/{id}/uploadimage", petController.UploadPetImageById).Methods("POST")
	router.HandleFunc("/pet/{id}", petController.DeletePet).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}

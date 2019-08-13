package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PetAPI/cmd/petapi/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pet", controllers.GetPets).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}

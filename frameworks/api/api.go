package api

import (
	"net/http"

	"github.com/lucasasoaresmar/clean-go/adapters/controllers/api"
	"github.com/gorilla/mux"
)

// API framework
type API struct {
}

// Start server
func (a API) Start() {
	userHandler := api.NewUserHandler()
	router := mux.NewRouter()
	router.HandleFunc("/login", userHandler.Login).Methods("POST")
	http.ListenAndServe(":8000", router)
}

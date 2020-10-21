package router

import (
	"github.com/qimpl/housing/handlers"

	"github.com/gorilla/mux"
)

func createHousingStatusRouter(router *mux.Router) {
	housingStatusRouter := router.PathPrefix("/housing/status").Subrouter()

	housingStatusRouter.
		HandleFunc("", handlers.CreateHousingStatus).
		Methods("POST")

	housingStatusRouter.
		HandleFunc("", handlers.GetAllHousingStatuses).
		Methods("GET")
}

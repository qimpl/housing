package router

import (
	"github.com/qimpl/housing/handlers"

	"github.com/gorilla/mux"
)

func createHousingRouter(router *mux.Router) {
	healthyRouter := router.PathPrefix("/housing").Subrouter()

	healthyRouter.
		HandleFunc("", handlers.CreateHousing).
		Methods("POST")

	healthyRouter.
		HandleFunc("", handlers.GetAllHousing).
		Methods("GET")

	healthyRouter.
		HandleFunc("/{housing_type_id}", handlers.GetAllHousingByType).
		Methods("GET")
}

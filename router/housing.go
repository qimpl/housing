package router

import (
	"github.com/qimpl/housing/handlers"

	"github.com/gorilla/mux"
)

func createHousingRouter(router *mux.Router) {
	housingRouter := router.PathPrefix("/housing").Subrouter()

	housingRouter.
		HandleFunc("", handlers.CreateHousing).
		Methods("POST")

	housingRouter.
		HandleFunc("", handlers.GetAllHousing).
		Methods("GET")

	housingRouter.
		HandleFunc("/{housing_type_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.GetAllHousingByType).
		Methods("GET")

	housingRouter.
		HandleFunc("/type", handlers.CreateHousingType).
		Methods("POST")

	housingRouter.
		HandleFunc("/type", handlers.GetAllHousingTypes).
		Methods("GET")
}

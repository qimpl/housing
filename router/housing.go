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
		HandleFunc("/{housing_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.GetHousingByID).
		Methods("GET")

	housingRouter.
		HandleFunc("/{housing_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.DeleteHousingByID).
		Methods("DELETE")

	housingRouter.
		HandleFunc("/{housing_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.UpdateHousingByID).
		Methods("PUT")

	housingRouter.
		HandleFunc("/{housing_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/status", handlers.UpdateHousingStatus).
		Methods("PUT")

	housingRouter.
		HandleFunc("/{housing_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/publication", handlers.UpdateHousingPublicationStatus).
		Methods("PATCH")

	housingRouter.
		HandleFunc("/owner/{owner_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.GetHousingByOwnerID).
		Methods("GET")
}

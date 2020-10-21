package router

import (
	"github.com/qimpl/housing/handlers"

	"github.com/gorilla/mux"
)

func createHousingTypeRouter(router *mux.Router) {
	housingTypeRouter := router.PathPrefix("/housing/type").Subrouter()

	housingTypeRouter.
		HandleFunc("", handlers.CreateHousingType).
		Methods("POST")

	housingTypeRouter.
		HandleFunc("", handlers.GetAllHousingTypes).
		Methods("GET")

	housingTypeRouter.
		HandleFunc("/{housing_type_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.GetAllHousingByType).
		Methods("GET")

}

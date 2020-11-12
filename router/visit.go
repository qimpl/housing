package router

import (
	"github.com/qimpl/housing/handlers"

	"github.com/gorilla/mux"
)

func createVisitRouter(router *mux.Router) {
	housingRouter := router.PathPrefix("/visit").Subrouter()

	housingRouter.
		HandleFunc("", handlers.CreateVisit).
		Methods("POST")

	housingRouter.
		HandleFunc("/housing/{housing_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", handlers.GetVisitByHousingID).
		Methods("GET")

	housingRouter.
		HandleFunc("/{visit_id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/accept", handlers.AcceptVisit).
		Methods("PATCH")
}

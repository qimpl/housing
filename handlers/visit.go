package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qimpl/housing/db"
	"github.com/qimpl/housing/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateVisit add a new visit inside the database
// @Summary Insert a visit
// @Description Create a new housing visit with given data
// @Tags Visits
// @Accept json
// @Produce json
// @Param visit body models.Visit true "Visit data"
// @Success 201 {string} models.Visit
// @Failure 400 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Router /visit [post]
func CreateVisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var visit *models.Visit

	if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Body malformed data"))

		return
	}

	visit, err := db.CreateVisit(visit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("Visit creation failed"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(visit)
}

// GetVisitByHousingID get all visits of a given housing
// @Summary Get all visits of a housing
// @Description Get all visits of a housing with its ID
// @Tags Visits
// @Produce json
// @Param housing_id path string true "Housing ID"
// @Success 200 {string} []models.Visit
// @Failure 400 {string} models.ErrorResponse
// @Failure 404 {string} models.ErrorResponse
// @Router /visit/housing/{housing_id} [get]
func GetVisitByHousingID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	housingID := uuid.MustParse(mux.Vars(r)["housing_id"])

	if _, err := db.GetHousingByID(housingID); err != nil {
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing ID doesn't exist"))

		return
	}

	housingVisits, err := db.GetVisitByHousingID(housingID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing visits retrieval"))

		return
	}

	json.NewEncoder(w).Encode(housingVisits)
}

// AcceptVisit update a given visit booking IsAccepted field to true
// @Summary Accept a non accepted visit booking
// @Description Update a given visit booking IsAccepted field to true
// @Tags Visits
// @Produce json
// @Param visit_id path string true "Visit booking ID"
// @Success 204 ""
// @Failure 400 {string} models.ErrorResponse
// @Failure 404 {string} models.ErrorResponse
// @Router /visit/{visit_id}/accept [patch]
func AcceptVisit(w http.ResponseWriter, r *http.Request) {
	visitID := uuid.MustParse(mux.Vars(r)["visit_id"])

	if _, err := db.GetVisitBookingByID(visitID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given visit booking ID doesn't exist"))

		return
	}

	if err := db.AcceptVisit(visitID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during visit booking acceptation retrieval"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

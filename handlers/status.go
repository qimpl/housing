package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/qimpl/housing/db"
	"github.com/qimpl/housing/models"
)

// CreateHousingStatus handles housing status POST request and return a housing status if no error is returned
// @Summary Insert a housing status
// @Description Create a new housing status with given data
// @Tags Housing Status
// @Accept json
// @Produce json
// @Param status body models.Status true "Housing status data"
// @Success 200 {object} models.Status
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Router /housing/status [post]
func CreateHousingStatus(w http.ResponseWriter, r *http.Request) {
	var housingStatus *models.Status

	if err := json.NewDecoder(r.Body).Decode(&housingStatus); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))
		log.Printf("Status - Create - Body Error - %s : %s\n", time.Now(), err)

		return
	}

	housingStatus, err := db.CreateHousingStatus(housingStatus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("Housing status creation failed"))
		log.Printf("Status - Create - DB Creation Error - %s : %s\n", time.Now(), err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(housingStatus)
}

// GetAllHousingStatuses returns all the housing statuses found inside the database if no error is returned
// @Summary Get all housing statuses
// @Description Search all housing statuses
// @Tags Housing Status
// @Produce json
// @Success 200 {object} []models.Status
// @Failure 400 {object} models.ErrorResponse
// @Router /housing/status [get]
func GetAllHousingStatuses(w http.ResponseWriter, r *http.Request) {
	housingStatuses, err := db.GetAllHousingStatuses()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing statuses retrieval"))
		log.Printf("Status - Get All - DB Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housingStatuses)
}

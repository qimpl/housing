package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qimpl/housing/db"
	"github.com/qimpl/housing/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateHousing insert a new housing inside the database
// @Summary Insert a housing
// @Description Create a new housing with given data
// @Accept json
// @Produce json
// @Param housing body models.Housing true "Housing data"
// @Success 200 {string} models.Housing
// @Failure 400 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Router /housing [post]
func CreateHousing(w http.ResponseWriter, r *http.Request) {
	var housing *models.Housing

	if err := json.NewDecoder(r.Body).Decode(&housing); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Body malformed data"))

		return
	}

	housing, err := db.CreateHousing(housing)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("Housing creation failed"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(housing)
}

// GetAllHousing returns all the housing found on the request response
// @Summary Get all housing
// @Description Search all housing
// @Produce json
// @Success 200 {string} []models.Housing
// @Failure 400 {string} models.ErrorResponse
// @Router /housing [get]
func GetAllHousing(w http.ResponseWriter, r *http.Request) {
	housing, err := db.GetAllHousing()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing retrieval"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housing)
}

// GetAllHousingByType returns all the housing with a given type found on the request response
// @Summary Get all housing by type
// @Description Search all housing of a given type
// @Produce json
// @Param housing_type_id path string true "Housing type ID"
// @Success 200 {string} []models.Housing
// @Failure 400 {string} models.ErrorResponse
// @Failure 404 {string} models.ErrorResponse
// @Router /housing/{housing_type_id} [get]
func GetAllHousingByType(w http.ResponseWriter, r *http.Request) {
	validUUID := uuid.MustParse(mux.Vars(r)["housing_type_id"])

	if housingType, _ := db.GetHousingTypeByID(validUUID); housingType == nil {
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing type ID doesn't exist"))

		return
	}

	housing, err := db.GetHousingByType(validUUID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing retrieval"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housing)
}

// DeleteHousingByID delete a given housing with its ID
// @Summary Delete a housing by ID
// @Description Delete a given housing by ID
// @Param housing_id path string true "Housing ID"
// @Success 204 ""
// @Failure 400 {string} models.ErrorResponse
// @Router /housing/{housing_id} [delete]
func DeleteHousingByID(w http.ResponseWriter, r *http.Request) {
	if err := db.DeleteHousingByID(uuid.MustParse(mux.Vars(r)["housing_id"])); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing update"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdateHousingByID update a given housing with its ID
// @Summary Update a housing by ID
// @Description Update a given housing by ID
// @Accept json
// @Produce json
// @Param housing_id path string true "Housing ID"
// @Param housing body models.HousingBody true "Housing data"
// @Success 200 {string} models.Housing
// @Failure 400 {string} models.ErrorResponse
// @Failure 404 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Router /housing/{housing_id} [put]
func UpdateHousingByID(w http.ResponseWriter, r *http.Request) {
	var updatedHousing *models.HousingBody
	var housing *models.Housing

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&updatedHousing); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Body malformed data"))

		return
	}

	housing, err := db.GetHousingByID(uuid.MustParse(mux.Vars(r)["housing_id"]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing ID doesn't exist"))

		return
	}

	if _, err := db.UpdateHousingByID(housing, updatedHousing); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing update"))

		return
	}

	json.NewEncoder(w).Encode(housing)
}

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
// @Tags Housing
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
// @Tags Housing
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

// DeleteHousingByID delete a given housing with its ID
// @Summary Delete a housing by ID
// @Description Delete a given housing by ID
// @Tags Housing
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
// @Tags Housing
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

// UpdateHousingStatus update the status a given housing
// @Summary Update status of housing
// @Description Update the status of a given housing
// @Tags Housing
// @Accept json
// @Param housing_id path string true "Housing ID"
// @Param status body models.UpdateStatusBody true "Status ID"
// @Success 204 ""
// @Failure 400 {string} models.ErrorResponse
// @Failure 404 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Router /housing/{housing_id}/status [put]
func UpdateHousingStatus(w http.ResponseWriter, r *http.Request) {
	var body *models.UpdateStatusBody
	housingID := uuid.MustParse(mux.Vars(r)["housing_id"])

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Body malformed data"))

		return
	}

	if _, err := db.GetHousingByID(housingID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing ID doesn't exist"))

		return
	}

	if _, err := db.GetStatusByID(body.StatusID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given status ID doesn't exist"))

		return
	}

	if err := db.UpdateHousingStatus(housingID, body.StatusID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing status update"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetHousingByID returns a given housing with a given ID found on the request response
// @Summary Get a housing by ID
// @Description Search for a given housing with its ID
// @Tags Housing
// @Produce json
// @Param housing_id path string true "Housing ID"
// @Success 200 {string} models.Housing
// @Failure 400 {string} models.ErrorResponse
// @Router /housing/{housing_id} [get]
func GetHousingByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	housing, err := db.GetHousingByID(uuid.MustParse(mux.Vars(r)["housing_id"]))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing retrieval"))

		return
	}

	json.NewEncoder(w).Encode(housing)
}

// GetHousingByOwnerID returns all housings of a given owner
// @Summary Get all housings of a owner
// @Description Search for all housings with a given owner ID
// @Tags Housing
// @Produce json
// @Param owner_id path string true "Owner ID"
// @Success 200 {string} []models.Housing
// @Failure 400 {string} models.ErrorResponse
// @Router /housing/owner/{owner_id} [get]
func GetHousingByOwnerID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	housings, err := db.GetHousingByOwnerID(uuid.MustParse(mux.Vars(r)["owner_id"]))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housings list retrieval"))

		return
	}

	json.NewEncoder(w).Encode(housings)
}

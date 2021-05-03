package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/qimpl/housing/db"
	"github.com/qimpl/housing/models"
)

// CreateHousingType handles housing type POST request and return a housing type if no error is returned
// @Summary Insert a housing type
// @Description Create a new housing type with given data
// @Tags Housing Types
// @Accept json
// @Produce json
// @Param type body models.HousingType true "Housing type data"
// @Success 200 {object} models.HousingType
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Router /housing/type [post]
func CreateHousingType(w http.ResponseWriter, r *http.Request) {
	var housingType *models.HousingType

	if err := json.NewDecoder(r.Body).Decode(&housingType); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))
		log.Printf("Housing Type - Create - Body Error - %s : %s\n", time.Now(), err)

		return
	}

	housingType, err := db.CreateHousingType(housingType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("Housing type creation failed"))
		log.Printf("Housing Type - Create - Creation Error - %s : %s\n", time.Now(), err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(housingType)
}

// GetAllHousingTypes returns all the housing types found inside the database if no error is returned
// @Summary Get all housing types
// @Description Search all housing types
// @Tags Housing Types
// @Produce json
// @Success 200 {object} []models.HousingType
// @Failure 400 {object} models.ErrorResponse
// @Router /housing/type [get]
func GetAllHousingTypes(w http.ResponseWriter, r *http.Request) {
	housingTypes, err := db.GetAllHousingTypes()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing types retrieval"))
		log.Printf("Housing Type - Get All - DB Error - %s : %s\n", time.Now(), err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housingTypes)
}

// GetAllHousingByType returns all the housing with a given type found on the request response
// @Summary Get all housing by type
// @Description Search all housing of a given type
// @Tags Housing Types
// @Produce json
// @Param housing_type_id path string true "Housing type ID"
// @Success 200 {object} []models.Housing
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /housing/type/{housing_type_id} [get]
func GetAllHousingByType(w http.ResponseWriter, r *http.Request) {
	validUUID := uuid.MustParse(mux.Vars(r)["housing_type_id"])

	if housingType, err := db.GetHousingTypeByID(validUUID); housingType == nil {
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing type ID doesn't exist"))
		log.Printf("Housing Type - Get All By Type - ID not found - %s : %s\n", time.Now(), err)

		return
	}

	housing, err := db.GetHousingByType(validUUID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing retrieval"))
		log.Printf("Housing Type - Get All By Type - DB Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housing)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qimpl/housing/db"
	"github.com/qimpl/housing/models"
)

// CreateHousingType handles housing type POST request and return a housing type if no error is returned
// @Summary Insert a housing type
// @Description Create a new housing type with given data
// @Accept json
// @Produce json
// @Param type body models.HousingType true "Housing type data"
// @Success 200 {string} models.HousingType
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /housing/type [post]
func CreateHousingType(w http.ResponseWriter, r *http.Request) {
	var housingType *models.HousingType

	if err := json.NewDecoder(r.Body).Decode(&housingType); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")
		return
	}

	housingType, err := db.CreateHousingType(housingType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Housing type creation failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(housingType)
}

// GetAllHousingTypes returns all the housing types found inside the database if no error is returned
// @Summary Get all housing types
// @Description Search all housing types
// @Produce json
// @Success 200 {string} []models.HousingType
// @Failure 400 {string} string
// @Router /housing/type [get]
func GetAllHousingTypes(w http.ResponseWriter, r *http.Request) {
	housingTypes, err := db.GetAllHousingTypes()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("An error occurred during housing types retrieval")

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housingTypes)
}

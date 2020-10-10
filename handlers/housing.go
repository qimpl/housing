package handlers

import (
	"encoding/json"
	"log"
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
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /housing [post]
func CreateHousing(w http.ResponseWriter, r *http.Request) {
	var housing *models.Housing

	if err := json.NewDecoder(r.Body).Decode(&housing); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Println(err)
		json.NewEncoder(w).Encode("Body malformed data")
		return
	}

	log.Println(housing)

	housing, err := db.CreateHousing(housing)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		json.NewEncoder(w).Encode("Housing creation failed")
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
// @Failure 404 {string} string
// @Router /housing [get]
func GetAllHousing(w http.ResponseWriter, r *http.Request) {
	housing, err := db.GetAllHousing()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("An error occurred during housing retrieval")

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
// @Failure 404 {string} string
// @Router /housing/{housing_type_id} [get]
func GetAllHousingByType(w http.ResponseWriter, r *http.Request) {
	housingTypeID := mux.Vars(r)["housing_type_id"]

	housing, err := db.GetHousingByType(uuid.MustParse(housingTypeID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("An error occurred during housing retrieval")

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housing)
}

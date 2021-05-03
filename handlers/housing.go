package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/qimpl/housing/db"
	"github.com/qimpl/housing/models"
	"github.com/qimpl/housing/services"
	"github.com/qimpl/housing/storage"

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
// @Success 200 {object} models.Housing
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Router /housing [post]
func CreateHousing(w http.ResponseWriter, r *http.Request) {
	var housing *models.Housing

	if err := json.NewDecoder(r.Body).Decode(&housing); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Body malformed data"))
		log.Printf("Housing - Create - Body Error - %s : %s\n", time.Now(), err)

		return
	}

	coordinates, _ := services.GetCoordinatesByAddress(
		fmt.Sprintf("%s, %s %s %s", housing.Street, housing.City, housing.Zip, housing.Country),
	)

	if len(coordinates) > 0 {
		housing.Latitude = coordinates[0].Geometry.Location.Lat
		housing.Longitude = coordinates[0].Geometry.Location.Lng
	}

	housing, err := db.CreateHousing(housing)
	var badRequest *models.BadRequest

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(badRequest.GetError("Housing creation failed"))
		log.Printf("Housing - Create - Creation Error - %s : %s\n", time.Now(), err)

		return
	}

	if housing.Pictures != nil {
		for index, picture := range housing.Pictures {
			decodedImage, err := base64.StdEncoding.DecodeString(picture)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during profile picture decoding"))
				log.Printf("Housing - Create - Picture Decoding Error - %s : %s\n", time.Now(), err)

				return
			}

			imageReader := bytes.NewReader(decodedImage)

			storage.AddToBucket(
				fmt.Sprintf("%s/housing_picture_%s_%d.png", housing.ID, housing.ID, index+1),
				imageReader,
				imageReader.Size(),
				http.DetectContentType(decodedImage),
			)
		}
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
// @Success 200 {object} []models.Housing
// @Failure 400 {object} models.ErrorResponse
// @Router /housing [get]
func GetAllHousing(w http.ResponseWriter, r *http.Request) {
	housings, err := db.GetAllHousing()
	var badRequest *models.BadRequest

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing retrieval"))
		log.Printf("Housing - Get All - DB Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	for index, housing := range housings {
		for i := 1; i <= 5; i++ {
			if picture, err := storage.GetFromBucket(
				fmt.Sprintf("%s/housing_picture_%s_%d.png", housing.ID, housing.ID, i),
			); err == nil {
				housing.Pictures = append(housing.Pictures, picture)
			}
		}

		housings[index] = housing
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housings)
}

// DeleteHousingByID delete a given housing with its ID
// @Summary Delete a housing by ID
// @Description Delete a given housing by ID
// @Tags Housing
// @Param housing_id path string true "Housing ID"
// @Success 204 ""
// @Failure 400 {object} models.ErrorResponse
// @Router /housing/{housing_id} [delete]
func DeleteHousingByID(w http.ResponseWriter, r *http.Request) {
	if err := db.DeleteHousingByID(uuid.MustParse(mux.Vars(r)["housing_id"])); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing deletion"))
		log.Printf("Housing - Delete - DB Error - %s : %s\n", time.Now(), err)

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
// @Success 200 {object} models.Housing
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
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
		log.Printf("Housing - Update - Body Error - %s : %s\n", time.Now(), err)

		return
	}

	housing, err := db.GetHousingByID(uuid.MustParse(mux.Vars(r)["housing_id"]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing ID doesn't exist"))
		log.Printf("Housing - Update - DB retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	if housing.Street != updatedHousing.Street ||
		housing.City != updatedHousing.City ||
		housing.Zip != updatedHousing.Zip ||
		housing.Country != updatedHousing.Country {

		coordinates, _ := services.GetCoordinatesByAddress(fmt.Sprintf("%s, %s %s %s", updatedHousing.Street, updatedHousing.City, updatedHousing.Zip, updatedHousing.Country))
		if len(coordinates) > 0 {
			updatedHousing.Latitude = coordinates[0].Geometry.Location.Lat
			updatedHousing.Longitude = coordinates[0].Geometry.Location.Lng
		}
	}

	if _, err := db.UpdateHousingByID(housing, updatedHousing); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing update"))
		log.Printf("Housing - Update - DB Update Error - %s : %s\n", time.Now(), err)

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
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
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
		log.Printf("Housing - Update Status - Body Error - %s : %s\n", time.Now(), err)

		return
	}

	if _, err := db.GetHousingByID(housingID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing ID doesn't exist"))
		log.Printf("Housing - Update Status - Housing ID Error - %s : %s\n", time.Now(), err)

		return
	}

	if _, err := db.GetStatusByID(body.StatusID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given status ID doesn't exist"))
		log.Printf("Housing - Update Status - Status ID Error - %s : %s\n", time.Now(), err)

		return
	}

	if err := db.UpdateHousingStatus(housingID, body.StatusID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing status update"))
		log.Printf("Housing - Update Status - DB Update Error - %s : %s\n", time.Now(), err)

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
// @Success 200 {object} models.Housing
// @Failure 400 {object} models.ErrorResponse
// @Router /housing/{housing_id} [get]
func GetHousingByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	housing, err := db.GetHousingByID(uuid.MustParse(mux.Vars(r)["housing_id"]))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing retrieval"))
		log.Printf("Housing - Get By ID - DB Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	for i := 1; i <= 5; i++ {
		if picture, err := storage.GetFromBucket(
			fmt.Sprintf("%s/housing_picture_%s_%d.png", housing.ID, housing.ID, i),
		); err == nil {
			housing.Pictures = append(housing.Pictures, picture)
		}
	}

	json.NewEncoder(w).Encode(housing)
}

// GetHousingByOwnerID returns all housings of a given owner
// @Summary Get all housings of a owner
// @Description Search for all housings with a given owner ID
// @Tags Housing
// @Produce json
// @Param owner_id path string true "Owner ID"
// @Success 200 {object} []models.Housing
// @Failure 400 {object} models.ErrorResponse
// @Router /housing/owner/{owner_id} [get]
func GetHousingByOwnerID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	housings, err := db.GetHousingByOwnerID(uuid.MustParse(mux.Vars(r)["owner_id"]))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housings list retrieval"))
		log.Printf("Housing - Get By Owner ID - DB Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	json.NewEncoder(w).Encode(housings)
}

// UpdateHousingPublicationStatus updates the publication status of a housing
// @Summary Update publication status
// @Description Update the publication status of a housing
// @Tags Housing
// @Accept json
// @Param housing_id path string true "Housing ID"
// @Param status body models.UpdatePublicationStatus true "Status ID"
// @Success 204 ""
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Router /housing/{housing_id}/publication [patch]
func UpdateHousingPublicationStatus(w http.ResponseWriter, r *http.Request) {
	var body *models.UpdatePublicationStatus
	housingID := uuid.MustParse(mux.Vars(r)["housing_id"])

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Body malformed data"))
		log.Printf("Housing - Update Publication Status - Body Error - %s : %s\n", time.Now(), err)

		return
	}

	if _, err := db.GetHousingByID(housingID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing ID doesn't exist"))
		log.Printf("Housing - Update Publication Status - DB ID Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	if err := db.UpdateHousingPublicationStatus(housingID, body.IsPublished); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing status update"))
		log.Printf("Housing - Update Publication Status - Status Update Error - %s : %s\n", time.Now(), err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetFilteredHousings returns all housings of given filters
// @Summary Get all housings depending on filters
// @Description Search for all housings with given filters
// @Tags Housing
// @Produce json
// @Param type_id path string true "Type ID"
// @Param city path string true "City"
// @Param price path string true "Rent Price"
// @Param size path string true "Surface Area"
// @Param status_id path string true "Status ID"
// @Success 200 {object} []models.Housing
// @Failure 400 {object} models.ErrorResponse
// @Router /housing/filter/{type_id} [get]
func GetFilteredHousings(w http.ResponseWriter, r *http.Request) {
	housingType := uuid.MustParse(mux.Vars(r)["type_id"])
	city := mux.Vars(r)["city"]
	housingStatus := uuid.MustParse(mux.Vars(r)["status"])
	v := r.URL.Query()
	housingRentPrice := v.Get("max_price")
	housingSurfaceArea := v.Get("min_size")

	housings, err := db.GetFilteredHousing(housingType, city, housingRentPrice, housingSurfaceArea, housingStatus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing retrieval"))
		log.Printf("Housing - Get Filtered - DB Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	for index, housing := range housings {
		if picture, err := storage.GetFromBucket(
			fmt.Sprintf("%s/housing_picture_%s_1.png", housing.ID, housing.ID),
		); err == nil {
			housing.Pictures = append(housing.Pictures, picture)
		}

		housings[index] = housing
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(housings)
}

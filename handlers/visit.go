package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/qimpl/housing/db"
	"github.com/qimpl/housing/models"
	"github.com/qimpl/housing/services"

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
// @Success 201 {object} models.Visit
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Router /visit [post]
func CreateVisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body *models.CreateVisitBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Body malformed data"))
		log.Printf("Visit - Create - Body Error - %s : %s\n", time.Now(), err)

		return
	}

	var err error
	visit := &models.Visit{
		Date:      body.Date,
		Hour:      body.Hour,
		HousingID: body.HousingID,
	}

	visit, err = db.CreateVisit(visit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("Visit creation failed"))
		log.Printf("Visit - Create - DB Creation Error - %s : %s\n", time.Now(), err)

		return
	}

	services.SendEmail(
		fmt.Sprintf("Qimpl | %s veut visiter votre bien ! 🎉", body.Owner.FirstName),
		"",
		body.Owner.Email,
		"visit_reservation",
		&services.TemplateValues{
			IsEng: false,
			Variables: map[string]string{
				"first_name": body.Owner.FirstName,
				"visit_link": fmt.Sprintf("https://qimpl.io/visit/%s", visit.ID),
			},
		},
	)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(visit)
}

// GetVisitByHousingID get all visits of a given housing
// @Summary Get all visits of a housing
// @Description Get all visits of a housing with its ID
// @Tags Visits
// @Produce json
// @Param housing_id path string true "Housing ID"
// @Success 200 {object} []models.Visit
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /visit/housing/{housing_id} [get]
func GetVisitByHousingID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	housingID := uuid.MustParse(mux.Vars(r)["housing_id"])

	if _, err := db.GetHousingByID(housingID); err != nil {
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given housing ID doesn't exist"))
		log.Printf("Visit - Get By ID - DB ID Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	housingVisits, err := db.GetVisitByHousingID(housingID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during housing visits retrieval"))
		log.Printf("Visit - Get By ID - DB Visit Retrieval Error - %s : %s\n", time.Now(), err)

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
// @Param visit body models.AcceptVisit true "Accept Visit data"
// @Success 204 ""
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /visit/{visit_id}/accept [patch]
func AcceptVisit(w http.ResponseWriter, r *http.Request) {
	var body *models.AcceptVisit

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Body malformed data"))
		log.Printf("Visit - Accept - Body Error - %s : %s\n", time.Now(), err)

		return
	}

	visitID := uuid.MustParse(mux.Vars(r)["visit_id"])
	visit, err := db.GetVisitBookingByID(visitID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		var notFound *models.NotFound
		json.NewEncoder(w).Encode(notFound.GetError("The given visit booking ID doesn't exist"))
		log.Printf("Visit - Accept - DB ID Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	if err := db.AcceptVisit(visitID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		var badRequest *models.BadRequest
		json.NewEncoder(w).Encode(badRequest.GetError("An error occurred during visit booking acceptation retrieval"))
		log.Printf("Visit - Accept - DB Booking Acceptation Retrieval Error - %s : %s\n", time.Now(), err)

		return
	}

	visitDate := fmt.Sprintf("%s %s", visit.Date.Format("02-Jan-2006"), visit.Hour)

	services.SendEmail(
		fmt.Sprintf("Qimpl | Votre demande de visite du %s a été acceptée ! 🔥", visitDate),
		"",
		body.UserEmail,
		"visit_accepted",
		&services.TemplateValues{
			IsEng: false,
			Variables: map[string]string{
				"visit_date": visitDate,
				"visit_link": fmt.Sprintf("https://qimpl.io/visit/%s", visitID),
			},
		},
	)

	w.WriteHeader(http.StatusNoContent)
}

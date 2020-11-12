package db

import (
	"github.com/qimpl/housing/models"

	"github.com/google/uuid"
)

// CreateVisit add a new housing visit inside the database
func CreateVisit(visit *models.Visit) (*models.Visit, error) {
	if _, err := Db.Model(visit).Insert(); err != nil {
		return nil, err
	}

	return visit, nil
}

// GetVisitByHousingID get all visits of a given housing
func GetVisitByHousingID(housingID uuid.UUID) ([]models.Visit, error) {
	var visit []models.Visit

	if err := Db.Model(&visit).Where("housing_id = ?", housingID).Select(); err != nil {
		return nil, err
	}

	return visit, nil
}

// GetVisitBookingByID get a given visit booking
func GetVisitBookingByID(visitID uuid.UUID) (*models.Visit, error) {
	visit := new(models.Visit)

	if err := Db.Model(visit).Where("id = ?", visitID).Select(); err != nil {
		return nil, err
	}

	return visit, nil
}

// AcceptVisit change IsAccepted field of a given visit booking to true
func AcceptVisit(visitID uuid.UUID) error {
	var visit *models.Visit

	if _, err := Db.Model(visit).Set("is_accepted = ?", true).Where("id = ?", visitID).Update(); err != nil {
		return err
	}

	return nil
}

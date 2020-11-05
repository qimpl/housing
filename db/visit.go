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

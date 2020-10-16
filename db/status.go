package db

import (
	"github.com/google/uuid"
	"github.com/qimpl/housing/models"
)

// CreateHousingStatus add a new housing status inside the database
func CreateHousingStatus(housingStatus *models.Status) (*models.Status, error) {
	if _, err := Db.Model(housingStatus).Insert(); err != nil {
		return nil, err
	}

	return housingStatus, nil
}

// GetAllHousingStatuses get all housing statuses from the database
func GetAllHousingStatuses() ([]models.Status, error) {
	var housingStatuses []models.Status

	if err := Db.Model(&housingStatuses).Order("created_at").Select(); err != nil {
		return nil, err
	}

	return housingStatuses, nil
}

// GetStatusByID find a status inside the database with its ID
func GetStatusByID(statusID uuid.UUID) (*models.Status, error) {
	status := new(models.Status)

	if err := Db.Model(status).Where("id = ?", statusID).Select(); err != nil {
		return nil, err
	}

	return status, nil
}

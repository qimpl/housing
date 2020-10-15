package db

import (
	"github.com/qimpl/housing/models"

	"github.com/google/uuid"
)

// CreateHousing add a new housing inside the database
func CreateHousing(housing *models.Housing) (*models.Housing, error) {
	if _, err := Db.Model(housing).Insert(); err != nil {
		return nil, err
	}

	return housing, nil
}

// GetAllHousing get all housings from the database
func GetAllHousing() ([]models.Housing, error) {
	var housing []models.Housing

	if err := Db.Model(&housing).Order("created_at").Select(); err != nil {
		return nil, err
	}

	return housing, nil
}

// GetHousingByType get all housings from the database the match a given type
func GetHousingByType(housingTypeID uuid.UUID) ([]models.Housing, error) {
	var housing []models.Housing

	if err := Db.Model(&housing).Order("created_at").Where("type_id = ?", housingTypeID.String()).Select(); err != nil {
		return nil, err
	}

	return housing, nil
}

// DeleteHousingByID delete a given housing
func DeleteHousingByID(housingID uuid.UUID) error {
	var housing *models.Housing

	if _, err := Db.Model(housing).Where("id = ?", housingID).Delete(); err != nil {
		return err
	}

	return nil
}

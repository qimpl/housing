package db

import (
	"github.com/google/uuid"
	"github.com/qimpl/housing/models"
)

// CreateHousingType add a new housing type inside the database
func CreateHousingType(housingType *models.HousingType) (*models.HousingType, error) {
	if _, err := Db.Model(housingType).Insert(); err != nil {
		return nil, err
	}

	return housingType, nil
}

// GetAllHousingTypes get all housings types from the database
func GetAllHousingTypes() ([]models.HousingType, error) {
	var housingTypes []models.HousingType

	if err := Db.Model(&housingTypes).Order("created_at").Select(); err != nil {
		return nil, err
	}

	return housingTypes, nil
}

// GetHousingTypeByID get a housings type from the database with a given ID
func GetHousingTypeByID(housingTypeID uuid.UUID) (*models.HousingType, error) {
	housingType := &models.HousingType{ID: housingTypeID}

	if err := Db.Select(housingType); err != nil {
		return nil, err
	}

	return housingType, nil
}

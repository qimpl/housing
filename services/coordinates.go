package services

import (
	"context"
	"os"

	"googlemaps.github.io/maps"
)

// GetCoordinatesByAddress return geocoding information from a given address
func GetCoordinatesByAddress(address string) ([]maps.GeocodingResult, error) {
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		return nil, err
	}

	request := &maps.GeocodingRequest{
		Address: address,
	}

	return client.Geocode(context.Background(), request)
}

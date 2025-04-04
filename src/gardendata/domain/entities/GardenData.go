package entities

import "time"

// GardenData represents a single record of sensor data from a garden kit.
type GardenData struct {
	DataID              int64     `json:"data_id"`
	KitID               int64     `json:"kit_id"` // Changed to int64 for consistency if IDs can grow large
	Temperature         float64   `json:"temperature"`
	GroundHumidity      float64   `json:"ground_humidity"`
	EnvironmentHumidity float64   `json:"environment_humidity"` // Corrected spelling
	PhLevel             float64   `json:"ph_level"`
	Time                int64     `json:"time"`      // Unix timestamp from device
	Timestamp           time.Time `json:"timestamp"` // DB insertion timestamp
}

// GardenDataResponse defines the structure returned by the API, potentially omitting fields if needed.
// In this case, it's the same as GardenData.
type GardenDataResponse struct {
	DataID              int64     `json:"data_id"`
	KitID               int64     `json:"kit_id"`
	Temperature         float64   `json:"temperature"`
	GroundHumidity      float64   `json:"ground_humidity"`
	EnvironmentHumidity float64   `json:"environment_humidity"`
	PhLevel             float64   `json:"ph_level"`
	Time                int64     `json:"time"`
	Timestamp           time.Time `json:"timestamp"`
}

// ToResponse converts GardenData to GardenDataResponse.
// Useful if you ever want to hide certain fields in the future.
func (gd *GardenData) ToResponse() GardenDataResponse {
	return GardenDataResponse{
		DataID:              gd.DataID,
		KitID:               gd.KitID,
		Temperature:         gd.Temperature,
		GroundHumidity:      gd.GroundHumidity,
		EnvironmentHumidity: gd.EnvironmentHumidity,
		PhLevel:             gd.PhLevel,
		Time:                gd.Time,
		Timestamp:           gd.Timestamp,
	}
}

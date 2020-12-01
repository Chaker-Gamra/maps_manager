package models

type DistanceMatrixEntity struct {
	DurationInTraffic float64 `json:"durationInTraffic" example:"3062"`
	Distance          float64 `json:"distance" example:"57824"`
	Duration          float64 `json:"duration" example:"2562"`
}

type DistanceMatrix struct {
	DirectionsArray [][]DistanceMatrixEntity `json:"directionsArray"`
}

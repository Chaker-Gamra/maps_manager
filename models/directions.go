package models

type Location struct {
	lat float64 `json:"lat" example:"34.1330949"`
	lng float64 `json:"lng" example:"-117.9143879"`
}

type Direction struct {
	DurationInTraffic float64    `json:"durationInTraffic" example:"3062"`
	Distance          float64    `json:"orderedAt" example:"57824"`
	Path              []Location `json:"path"`
}

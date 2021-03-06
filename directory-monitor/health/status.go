// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

// Package health is used by inway and healthceck to communicate
// health/status in a standardized json format.
package health

// Status is a structure used to indicate health of a service
// provided by an inway.
type Status struct {
	Healthy bool   `json:"healthy"`
	Version string `json:"version"`
}

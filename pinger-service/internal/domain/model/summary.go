package model

import "time"

type ResultSummary struct {
	URL         string    `json:"url"`
	Uptime      float64   `json:"uptime"`
	TotalChecks int       `json:"total_checks"`
	LastStatus  string    `json:"last_status"`
	LastChecked time.Time `json:"last_checked"`
}

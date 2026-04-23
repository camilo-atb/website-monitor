package model

import "time"

type PingResult struct {
	URL          string        `json:"url"`
	Status       string        `json:"status"`
	StatusCode   int           `json:"status_code"`
	ResponseTime time.Duration `json:"response_time"`
	CheckedAt    time.Time     `json:"checked_at"`
	Error        string        `json:"error"`
}

package model

import "time"

type MonitoredURL struct {
	ID           int       `json:"id"`
	URL          string    `json:"url"`
	ReviewTime   int       `json:"reviewTime"`
	CreationDate time.Time `json:"creationDate"`
	ModifyDate   time.Time `json:"modifyDate"`
}

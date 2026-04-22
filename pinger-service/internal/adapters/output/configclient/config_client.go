package configclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pinger/internal/domain/model"
	"pinger/internal/domain/ports"
	"time"
)

type configClient struct{}

func NewConfigClient() ports.ConfigServicePort {
	return &configClient{}
}

func (c *configClient) GetSites(ctx context.Context) ([]model.MonitoredURL, error) {

	log.Println("consultando config-service...")

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/sites", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	var sites []model.MonitoredURL

	err = json.NewDecoder(resp.Body).Decode(&sites)
	if err != nil {
		return nil, err
	}

	return sites, nil
}

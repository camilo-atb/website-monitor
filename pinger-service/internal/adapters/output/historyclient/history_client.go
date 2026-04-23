package historyclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"pinger/internal/domain/model"
	"pinger/internal/domain/ports"
	"time"
)

type historyClient struct {
	baseURL string
}

func NewHistoryClient(baseURL string) ports.HistoryPort {
	return &historyClient{baseURL: baseURL}
}

func (c *historyClient) Save(ctx context.Context, result model.PingResult) error {

	body, err := json.Marshal(result)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/results",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error status: %d", resp.StatusCode)
	}

	return nil
}

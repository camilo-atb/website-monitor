package httpclient

import (
	"context"
	"log"
	"net/http"
	"pinger/internal/domain/ports"
	"time"
)

type httpClient struct{}

func NewHTTPClient() ports.HTTPClientPort {
	return &httpClient{}
}

func (c *httpClient) Get(ctx context.Context, url string) (int, time.Duration, error) {

	log.Println("haciendo GET a:", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, 0, err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()

	resp, err := client.Do(req)
	duration := time.Since(start)

	if err != nil {
		return 0, duration, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, duration, nil
}

/*
Sin retries:
GET → falla → DOWN

Con retries:
GET → falla → esperar → retry → esperar → retry

Con circuit breaker:
muchos fallos →
	dejo de intentar →
	espero →
	pruebo otra vez
*/

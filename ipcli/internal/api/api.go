package api

import (
	"encoding/json"
	"fmt"
	"ipcli/internal/models"
	"net/http"
	"time"
)

const baseUrl string = "http://ip-api.com/json"

func Request(query string) (*models.ApiResponse, error) {
	client := &http.Client{Timeout: 4 * time.Second}
	resp, err := client.Get(baseUrl + "/" + query)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d", resp.StatusCode)
	}

	data := &models.ApiResponse{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode the json data: %w", err)
	}

	return data, nil
}

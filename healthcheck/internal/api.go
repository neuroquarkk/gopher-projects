package internal

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

func Request(endpoint string) (string, time.Duration) {
	client := http.Client{Timeout: 5 * time.Second}

	start := time.Now()
	resp, err := client.Get(endpoint)
	elapsed := time.Since(start)

	var status string

	if err != nil {
		var urlError *url.Error
		if errors.As(err, &urlError) && urlError.Timeout() {
			status = "Timeout"
		} else {
			status = "Error"
		}
	} else {
		status = resp.Status
	}

	return status, elapsed
}

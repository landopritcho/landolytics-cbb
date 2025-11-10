package espn

import (
	"fmt"
	"io"
	"net/http"
)

const (
	espnAllSchedulesBaseUrl = "https://site.api.espn.com/apis/site/v2/sports/basketball/mens-college-basketball/scoreboard?seasontype=2&groups=50"
)

func fetchJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func buildAllSchedulesUrl(date string) string {
	return fmt.Sprintf("%s&dates=%s",
		espnAllSchedulesBaseUrl,
		date,
	)
}

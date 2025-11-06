package cbb

import (
	"fmt"
	"strings"

	"github.com/landopritcho/landolytics-cbb/espn"
)

func FetchSchedule(date string) ([]Game, error) {

	cleanDate := strings.ReplaceAll(date, "-", "")

	if len(cleanDate) != 8 {
		return nil, fmt.Errorf("invalid date format: expected YYYYMMDD or YYYY-MM-DD, got %s", date)
	}

	espnGames, err := espn.FetchSchedule(cleanDate)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schedule: %w", err)
	}

	games := make([]Game, len(espnGames))
	for i, eg := range espnGames {
		games[i] = convertESPNGame(eg)
	}

	return games, nil
}

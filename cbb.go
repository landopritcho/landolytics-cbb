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

func FetchTeams() ([]InternalTeam, error) {
	espnTeams, err := espn.FetchTeams()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch teams: %w", err)
	}

	teams := make([]InternalTeam, len(espnTeams))
	for i, et := range espnTeams {
		teams[i] = convertESPNTeam(et)
	}

	return teams, nil
}

func FetchTeamById(id string) (InternalTeam, error) {
	teams, err := FetchTeams()
	if err != nil {
		return InternalTeam{}, fmt.Errorf("failed to fetch team: %w", err)
	}

	for _, team := range teams {
		if team.Id == id {
			return team, nil
		}
	}

	return InternalTeam{}, nil // There's probably a better way of returning no results but I'm leaving this for now
}

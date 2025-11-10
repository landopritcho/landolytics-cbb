package espn

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func FetchSchedule(date string) ([]ParsedGame, error) {
	url := buildAllSchedulesUrl(date)

	jsonData, err := fetchJSON(url)
	if err != nil {
		return nil, err
	}

	games, err := parseSchedule(jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse schedule: %w", err)
	}

	return games, nil
}

func parseSchedule(data []byte) ([]ParsedGame, error) {
	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	games := make([]ParsedGame, 0, len(response.Events))

	for _, event := range response.Events {
		game := parseEvent(event, response.DateInfo.Date)
		games = append(games, game)
	}

	return games, nil
}

func parseEvent(event Event, date string) ParsedGame {
	game := ParsedGame{
		Date:       date,
		LongDate:   event.Status.Details.DayDateTime,
		Time:       event.Status.Details.DateTime,
		LongTitle:  event.Name,
		ShortTitle: event.ShortName,
	}

	for _, comp := range event.Competitions {
		game.VenueName = comp.VenueInfo.Name
		game.VenueCity = comp.VenueInfo.Location.City
		game.VenueState = comp.VenueInfo.Location.State
		game.VenueLocation = fmt.Sprintf("%s, %s", comp.VenueInfo.Location.City, comp.VenueInfo.Location.State)
		game.IsNeutralSite = comp.NeutralSite
		game.IsConferenceGame = comp.ConferenceGame
		if len(comp.Notes) > 0 && strings.ToLower(comp.Notes[0].Type) == "event" {
			game.IsSpecialEvent = true
			game.EventTitle = comp.Notes[0].Headline
		} else {
			game.IsSpecialEvent = false
		}

		for _, team := range comp.Teams {
			parsedTeam := parseTeam(team)

			if strings.ToLower(team.HomeAway) == "home" {
				game.HomeTeam = parsedTeam
			} else {
				game.AwayTeam = parsedTeam
			}
		}
	}

	return game
}

func parseTeam(team Team) ParsedTeam {
	return ParsedTeam{
		School:         team.TeamInfo.SchoolName,
		Mascot:         team.TeamInfo.Mascot,
		Abbreviation:   team.TeamInfo.Abbreviation,
		FullName:       team.TeamInfo.DisplayName,
		ShortName:      team.TeamInfo.ShortDisplayName,
		PrimaryColor:   team.TeamInfo.PrimaryColor,
		AlternateColor: team.TeamInfo.AltColor,
		Rank:           parseRanking(team.Rank.TeamRanking),
	}
}

func parseRanking(rank json.Number) string {
	rankInt, err := strconv.Atoi(rank.String())
	if err != nil {
		return "Not Ranked"
	}

	if rankInt > 25 {
		return "Not Ranked"
	}

	return rank.String()
}

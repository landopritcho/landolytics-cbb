package espn

import "encoding/json"

type Response struct {
	Events   []Event `json:"events"`
	DateInfo struct {
		Date string `json:"date"`
	} `json:"day"`
}

type TeamsResponse struct {
	Sports []struct {
		Leagues []struct {
			Teams []struct {
				Team InternalTeam `json:"team"`
			} `json:"teams"`
		} `json:"leagues"`
	} `json:"sports"`
}

type InternalTeam struct {
	Id               string `json:"id"`
	Slug             string `json:"slug"`
	Abbreviation     string `json:"abbreviation"`
	DisplayName      string `json:"displayName"`
	ShortDisplayName string `json:"shortDisplayName"`
	Mascot           string `json:"name"`
}

type Event struct {
	Name         string        `json:"name"`
	ShortName    string        `json:"shortName"`
	Season       Season        `json:"season"`
	Competitions []Competition `json:"competitions"`
	Status       Status        `json:"status"`
}

type Season struct {
	Year int    `json:"year"`
	Type string `json:"slug"`
}

type Venue struct {
	Name     string `json:"fullName"`
	Location struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
}

type Competition struct {
	NeutralSite    bool   `json:"neutralSite"`
	ConferenceGame bool   `json:"conferenceCompetition"`
	VenueInfo      Venue  `json:"venue"`
	Teams          []Team `json:"competitors"`
	Notes          []struct {
		Type     string `json:"type"`
		Headline string `json:"headline"`
	} `json:"notes"`
}

type Status struct {
	Details struct {
		DayDateTime string `json:"detail"`
		DateTime    string `json:"shortDetail"`
	} `json:"type"`
}

type Team struct {
	HomeAway string `json:"homeAway"`
	TeamInfo struct {
		SchoolName       string `json:"location"`
		Mascot           string `json:"name"`
		Abbreviation     string `json:"abbreviation"`
		DisplayName      string `json:"displayName"`
		ShortDisplayName string `json:"shortDisplayName"`
		PrimaryColor     string `json:"color"`
		AltColor         string `json:"alternateColor"`
	} `json:"team"`
	Rank struct {
		TeamRanking json.Number `json:"current"`
	} `json:"curatedRank"`
}

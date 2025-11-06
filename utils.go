package cbb

import "github.com/landopritcho/landolytics-cbb/espn"

func convertESPNGame(eg espn.ParsedGame) Game {
	return Game{
		Date:       eg.Date,
		Time:       eg.Time,
		LongDate:   eg.LongDate,
		LongTitle:  eg.LongTitle,
		ShortTitle: eg.ShortTitle,
		Venue: Venue{
			Name:     eg.VenueName,
			City:     eg.VenueCity,
			State:    eg.VenueState,
			Location: eg.VenueLocation,
		},
		HomeTeam: Team{
			School:         eg.HomeTeam.School,
			Mascot:         eg.HomeTeam.Mascot,
			Abbreviation:   eg.HomeTeam.Abbreviation,
			FullName:       eg.HomeTeam.FullName,
			ShortName:      eg.HomeTeam.ShortName,
			PrimaryColor:   eg.HomeTeam.PrimaryColor,
			AlternateColor: eg.HomeTeam.AlternateColor,
			Rank:           eg.HomeTeam.Rank,
		},
		AwayTeam: Team{
			School:         eg.AwayTeam.School,
			Mascot:         eg.AwayTeam.Mascot,
			Abbreviation:   eg.AwayTeam.Abbreviation,
			FullName:       eg.AwayTeam.FullName,
			ShortName:      eg.AwayTeam.ShortName,
			PrimaryColor:   eg.AwayTeam.PrimaryColor,
			AlternateColor: eg.AwayTeam.AlternateColor,
			Rank:           eg.AwayTeam.Rank,
		},
		IsNeutralSite:    eg.IsNeutralSite,
		IsConferenceGame: eg.IsConferenceGame,
	}
}

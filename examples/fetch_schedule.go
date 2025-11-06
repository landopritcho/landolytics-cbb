package main

import (
	"fmt"
	"log"

	cbb "github.com/landopritcho/landolytics-cbb"
)

func main() {
	fmt.Println("Fetching games for November 18, 2025...")
	games, err := cbb.FetchSchedule("2025-11-18")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nFound %d games:\n\n", len(games))

	displayCount := 5
	if len(games) < 5 {
		displayCount = len(games)
	}

	for i := 0; i < displayCount; i++ {
		game := games[i]
		fmt.Printf("\nGame %d:\n", i+1)
		fmt.Printf("  Title: %s\n", game.ShortTitle)
		fmt.Printf("  Away: %s (Rank: %s)\n", game.AwayTeam.FullName, game.AwayTeam.Rank)
		fmt.Printf("  Home: %s (Rank: %s)\n", game.HomeTeam.FullName, game.HomeTeam.Rank)
		fmt.Printf("  Venue: %s\n", game.Venue.Name)
		fmt.Printf("  Location: %s\n", game.Venue.Location)
		fmt.Printf("  Time: %s\n", game.Time)
		fmt.Printf("  Neutral Site: %v\n", game.IsNeutralSite)
		fmt.Printf("  Conference Game: %v\n", game.IsConferenceGame)
		fmt.Printf("  Special Event: %v\n", game.IsSpecialEvent)
		if game.IsSpecialEvent {
			fmt.Printf("  Event Title: %s\n", game.EventTitle)
		}
	}

	if len(games) > displayCount {
		fmt.Printf("...and %d more games\n", len(games)-displayCount)
	}
}

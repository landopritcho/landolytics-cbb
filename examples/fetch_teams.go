package main

import (
	"fmt"
	"log"

	cbb "github.com/landopritcho/landolytics-cbb"
)

func main() {
	fmt.Println("Fetching teams...")
	teams, err := cbb.FetchTeams()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nFound %d teams:\n\n", len(teams))

	displayCount := 5
	if len(teams) < displayCount {
		displayCount = len(teams)
	}

	for i := 0; i < displayCount; i++ {
		team := teams[i]
		fmt.Printf("\nId: %s\n", team.Id)
		fmt.Printf("Slug: %s\n", team.Slug)
		fmt.Printf("Display Name: %s\n", team.DisplayName)
	}

	if len(teams) > displayCount {
		fmt.Printf("\n...and %d more teams\n", len(teams)-displayCount)
	}

}

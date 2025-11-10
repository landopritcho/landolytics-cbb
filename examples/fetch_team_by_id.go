package main

import (
	"fmt"
	"log"

	cbb "github.com/landopritcho/landolytics-cbb"
)

func main() {
	fmt.Println("Fetching team with ID 127")
	team, err := cbb.FetchTeamById("127")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nId: %s\n", team.Id)
	fmt.Printf("Slug: %s\n", team.Slug)
	fmt.Printf("Display Name: %s\n", team.DisplayName)

}

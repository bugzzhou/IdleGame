package main

import (
	"fight/everything"
	"fmt"
)

func main() {
	adventurer, err := everything.GetAdventurer()
	if err != nil {
		fmt.Printf("failed to get adventurer, and err is: %s\n", err.Error())
		return
	}

	for {
		monster := everything.Encounter(0)
		everything.Display(adventurer, monster)

		if !everything.HandleUserInput(&adventurer, monster) {
			return
		}
	}
}

package main

import (
	"fight/everything"
	"fmt"
)

func main() {
	fmt.Println(1)

	adventurer := everything.InitAdventurer()

	for {
		monster := everything.Encounter(0)

		everything.Display(adventurer, monster)

		var success bool
		var exp int
		var decide string
		fmt.Println("是否应对敌怪？")
		fmt.Scan(&decide)
		switch decide {
		case "y", "Y":
			success, exp, _ = everything.Combat(adventurer, monster)
		case "n", "N":
			continue
		default:
			fmt.Println("选择无效，重新选择，自动为你跳过该敌怪")
			continue
		}

		if success {
			everything.Develop(&adventurer, exp)
		} else {
			fmt.Println("输了")
			return
		}
	}

}

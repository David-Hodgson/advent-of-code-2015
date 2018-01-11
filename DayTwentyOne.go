package adventofcode2015

import (
	"fmt"
)

type character struct {
	hitPoints int
	damage int
	armour int
}

func attack(attacker *character, defender *character) {

		//Player
		playerStrike := attacker.damage - defender.armour
		if playerStrike < 1 {
			playerStrike = 1
		}

		defender.hitPoints = defender.hitPoints - playerStrike
}

func DayTwentyOneExample() {

	fmt.Println("Day 21 - Example")

	player := character{8,5,5}
	boss := character{12,7,2}

	for i := 0; i<4 ;i++ {
		fmt.Println("Round", i)
		attack(&player,&boss)
		if boss.hitPoints < 1 {
			fmt.Println("Boss loses")
			break
		}
		attack(&boss,&player)
		if player.hitPoints < 1 {
			fmt.Println("Player loses")
		}
	}

}

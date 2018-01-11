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

func isPlayerWinner(player, boss character) bool {
	playerWins := false;

	for ; ; {
		attack(&player,&boss)
		if boss.hitPoints < 1 {
			playerWins = true
			break
		}
		attack(&boss,&player)
		if player.hitPoints < 1 {
			playerWins = false
			break
		}
	}

	return playerWins
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

type ring struct {
	cost int
	damage int
	armour int
}

func DayTwentyOnePartOne() {

	fmt.Println("Day 21 - Part 1")

	weapons := []int {4,5,6,7,8}
	weaponCosts := []int {8,19,25,40,74}

	armourItems := []int {0,1,2,3,4,5}
	armourCosts := []int {0,13,31,53,75,102}

	rings := []ring {ring{0,0,0},ring{0,0,0},ring{25,1,0},ring{50,2,0},ring{100,3,0},ring{20,0,1},ring{40,0,2},ring{80,0,3}}

	minCost := 1000
	for i :=0; i <len(weapons); i++ {
		for k :=0; k < len(armourItems); k++ {
			for x := 0; x < len(rings); x ++ {
				for y :=0; y<len(rings);y++ {
					if x==y {
						continue
					}

					player := character{100,0,0}
					boss := character{104,8,1}

					player.damage += weapons[i]
					player.armour += armourItems[k]
					player.damage += rings[x].damage
					player.damage += rings[y].damage
					player.armour += rings[x].armour
					player.armour += rings[y].armour

					playerWins := isPlayerWinner(player,boss)

					if playerWins {
						cost := weaponCosts[i] + armourCosts[k] + rings[x].cost + rings[y].cost
						if cost < minCost {
							minCost = cost
						}
					}
				}
			}
		}
	}

	fmt.Println("Lowest Winner Cost is", minCost)
}

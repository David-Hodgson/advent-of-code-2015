package adventofcode2015

import (
	"fmt"
)

type spell struct {
	cost     int
	damage   int
	healing  int
	duration int
	recharge int
	armour   int
}

type magicCharacter struct {
	hitPoints int
	mana      int
	damage    int
}

func doTurn(player, boss magicCharacter, effects []spell, availableSpells []spell) {

	//appy effects
	for i := 0; i < len(effects); i++ {
		effectiveSpell := effects[i]

		boss.hitPoints -= effectiveSpell.damage
		player.hitPoints += effectiveSpell.healing
		player.mana += effectiveSpell.recharge

		effectiveSpell.duration -= 1

	}
}

func DayTwentyTwoExample() {

	fmt.Println("Day 22 - Example")

	magicMissile := spell{cost: 53, damage: 4, duration: 1}
	drain := spell{cost: 73, damage: 2, healing: 2, duration: 1}
	shield := spell{cost: 113, duration: 6, armour: 7}
	poison := spell{cost: 172, duration: 6, damage: 3}
	recharge := spell{cost: 229, duration: 5, recharge: 101}

	spells := []spell{magicMissile, drain, shield, poison, recharge}

	fmt.Println(spells)

	player := magicCharacter{10, 25, 0}
	boss := magicCharacter{13, 0, 8}

	fmt.Println(player)
	fmt.Println(boss)

	for i := 0; i < 2; i++ {

		//effects

		//player turn

		//effects

		//boss turn
	}
}

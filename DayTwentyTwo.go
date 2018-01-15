package adventofcode2015

import (
	"fmt"
)

type spell struct {
	name	 string
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
	canDoMagic bool
}

func doTurn(attacker,defender magicCharacter, attackerEffects,defenderEffects map[string]spell, availableSpells []spell,manaSpent int) {

	if attacker.hitPoints <= 0 || defender.hitPoints <= 0 {
		fmt.Println("Someone is dead. Mana Spent:", manaSpent)
		return
	}

	//DO attack stuff
	for name, attackerSpell := range attackerEffects {
		fmt.Println("Processing attacker spell:", name,"-",attackerSpell, "duration:", attackerSpell.duration)

		if attackerSpell.duration > 0 {
			attackerSpell.duration -= 1
			spell := attackerEffects[name]
			spell.duration -= 1
			attackerEffects[name] = spell

			attacker.hitPoints += spell.healing
			attacker.mana += spell.recharge
			defender.hitPoints -= spell.damage

			if defender.hitPoints <= 0 {
				fmt.Println("Some one killed by magic")
				return
			}
		} else {
			delete(attackerEffects, name)
		}
	}

	for name, defenderSpell := range defenderEffects {
		fmt.Println("Processing defender spell:", name,"-",defenderSpell, "duration:", defenderSpell.duration)

		if defenderSpell.duration > 0 {
			defenderSpell.duration -= 1

			spell := defenderEffects[name]
			spell.duration -= 1
			attackerEffects[name] = spell

			defender.hitPoints += spell.healing
			defender.mana += spell.recharge
			attacker.hitPoints -= spell.damage
		} else {
			delete(defenderEffects,name)
		}
	}
	if attacker.canDoMagic {
		fmt.Println("Attacker can do magic")

		for i:=0; i< len(availableSpells); i++ {
			spellToTry := availableSpells[i]
			if _,ok := attackerEffects[spellToTry.name]; !ok && attacker.mana > spellToTry.cost {
				fmt.Println("Casting spell", spellToTry.name)
				attackerEffects[spellToTry.name] = spellToTry

				doTurn(defender,attacker,defenderEffects,attackerEffects,availableSpells,manaSpent+spellToTry.cost)
			}
		}
	} else {
		fmt.Println("Attacker can't do magic")
		attackerDamage := attacker.damage

		defender.hitPoints -= attackerDamage
		fmt.Println("Defender hitPoints:", defender.hitPoints)
		if defender.hitPoints <= 0 {
			fmt.Println("Non magic attack killed")
			//return
		}

		doTurn(defender,attacker,defenderEffects,attackerEffects,availableSpells,manaSpent)
	}
}

func DayTwentyTwoExample() {

	fmt.Println("Day 22 - Example")

	magicMissile := spell{name: "Magic Missile", cost: 53, damage: 4, duration: 0}
	drain := spell{name: "Drain", cost: 73, damage: 2, healing: 2, duration: 0}
	shield := spell{name: "Shield", cost: 113, duration: 6, armour: 7}
	poison := spell{name: "Poison", cost: 172, duration: 6, damage: 3}
	recharge := spell{name: "Recharge", cost: 229, duration: 5, recharge: 101}

	spells := []spell{magicMissile, drain, shield, poison, recharge}

	fmt.Println(spells)

	player := magicCharacter{10, 250, 0,true}
	boss := magicCharacter{13, 0, 8,false}

	fmt.Println(player)
	fmt.Println(boss)

	playerEffects := make(map[string]spell)
	bossEffects := make(map[string]spell)

	doTurn(player,boss,playerEffects,bossEffects,spells,0)
}

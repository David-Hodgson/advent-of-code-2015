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

func copySpellMap(spellMap map[string]spell) map[string]spell {

	newSpellMap := make(map[string]spell)

	for name,spell := range spellMap {
		newSpell := spell
		newSpellMap[name] = newSpell
	}

	return newSpellMap
}

var minManaSpent int = 100000

func doTurn(attacker,defender magicCharacter, attackerEffects,defenderEffects map[string]spell, availableSpells []spell,manaSpent int, hardMode bool) {

	if manaSpent > minManaSpent {
		return
	}

	if hardMode && attacker.canDoMagic {
		attacker.hitPoints -= 1
		if attacker.hitPoints <=0 {
			return
		}
	}
	if attacker.hitPoints <= 0 || defender.hitPoints <= 0 {

		var deadCharacter magicCharacter
		if attacker.hitPoints <=0{
			deadCharacter = attacker
		} else {
			deadCharacter = defender
		}

		if !deadCharacter.canDoMagic {
			if manaSpent < minManaSpent {
				minManaSpent = manaSpent
			}
		}

		return
	}

	//DO attack stuff

	defenderArmourModifier := 0
	for name, attackerSpell := range attackerEffects {

		if attackerSpell.duration > 0 {
			attackerSpell.duration -= 1
			spell := attackerEffects[name]
			spell.duration -= 1
			attackerEffects[name] = spell

			attacker.hitPoints += spell.healing
			attacker.mana += spell.recharge
			defender.hitPoints -= spell.damage

			if defender.hitPoints <= 0 {
				if manaSpent < minManaSpent {
					minManaSpent = manaSpent
				}
				return
			}
			if spell.duration == 0 {
				delete(attackerEffects,name)
			}
		} else {
			delete(attackerEffects, name)
		}
	}

	for name, defenderSpell := range defenderEffects {

		if defenderSpell.duration > 0 {
			defenderSpell.duration -= 1

			spell := defenderEffects[name]
			spell.duration -= 1
			defenderEffects[name] = spell
			defender.hitPoints += spell.healing
			defender.mana += spell.recharge
			attacker.hitPoints -= spell.damage
			defenderArmourModifier += spell.armour

			if attacker.hitPoints <= 0 {

				if manaSpent < minManaSpent {
					minManaSpent = manaSpent
				}
				return
			}

			if spell.duration ==0 {
				delete(defenderEffects,name)
			}
		} else {
			delete(defenderEffects,name)
		}
	}
	if attacker.canDoMagic {

		for i:=0; i< len(availableSpells); i++ {
			spellToTry := availableSpells[i]
			spellCount := 0
			if _,ok := attackerEffects[spellToTry.name]; !ok && attacker.mana >= spellToTry.cost {
				spellCount++
				nextAttacker := defender
				nextDefender := attacker

				nextDefender.mana -= spellToTry.cost
				if spellToTry.duration == 0 {

					nextAttacker.hitPoints -= spellToTry.damage
					nextDefender.hitPoints += spellToTry.healing
				}
				newAttackerEffects := copySpellMap(attackerEffects)
				newDefenderEffects := copySpellMap(defenderEffects)

				newAttackerEffects[spellToTry.name] = spellToTry

				doTurn(nextAttacker,nextDefender,newDefenderEffects,newAttackerEffects,availableSpells,manaSpent+spellToTry.cost, hardMode)
			}

		}
	} else {
		attackerDamage := attacker.damage

		attackerDamage -= defenderArmourModifier

		if attackerDamage < 1 {
			attackerDamage = 1
		}

		defender.hitPoints -= attackerDamage
		if defender.hitPoints <= 0 {
			//return
		}

		nextAttacker := defender
		nextDefender := attacker

		newAttackerEffects := copySpellMap(attackerEffects)
		newDefenderEffects := copySpellMap(defenderEffects)

		doTurn(nextAttacker,nextDefender,newDefenderEffects,newAttackerEffects,availableSpells,manaSpent, hardMode)
	}
}

func DayTwentyTwoExample() {

	fmt.Println("Day 22 - Example")

	magicMissile := spell{name: "Magic Missile", cost: 53, damage: 4, duration: 0}
	drain := spell{name: "Drain", cost: 73, damage: 2, healing: 2, duration: 0}
	shield := spell{name: "Shield", cost: 113, duration: 6, armour: 7}
	poison := spell{name: "Poison", cost: 173, duration: 6, damage: 3}
	recharge := spell{name: "Recharge", cost: 229, duration: 5, recharge: 101}

	spells := []spell{magicMissile, drain, shield, poison, recharge}

	fmt.Println(spells)

	player := magicCharacter{10, 250, 0,true}
	boss := magicCharacter{13, 0, 8,false}

	fmt.Println(player)
	fmt.Println(boss)

	playerEffects := make(map[string]spell)
	bossEffects := make(map[string]spell)

	doTurn(player,boss,playerEffects,bossEffects,spells,0, false)

	fmt.Println("Min mana spent:", minManaSpent)
}

func DayTwentyTwoPartOne() {

	fmt.Println("Day 22 - Part One")

	minManaSpent = 10000000
	magicMissile := spell{name: "Magic Missile", cost: 53, damage: 4, duration: 0}
	drain := spell{name: "Drain", cost: 73, damage: 2, healing: 2, duration: 0}
	shield := spell{name: "Shield", cost: 113, duration: 6, armour: 7}
	poison := spell{name: "Poison", cost: 173, duration: 6, damage: 3}
	recharge := spell{name: "Recharge", cost: 229, duration: 5, recharge: 101}

	spells := []spell{magicMissile, drain, shield, poison, recharge}
	fmt.Println(spells)

	player := magicCharacter{50, 500, 0,true}
	boss := magicCharacter{51, 0, 9,false}

	fmt.Println(player)
	fmt.Println(boss)

	playerEffects := make(map[string]spell)
	bossEffects := make(map[string]spell)

	doTurn(player,boss,playerEffects,bossEffects,spells,0, false)

	fmt.Println("Min mana spent:", minManaSpent)
}

func DayTwentyTwoPartTwo() {

	fmt.Println("Day 22 - Part Two")

	minManaSpent = 100000000
	magicMissile := spell{name: "Magic Missile", cost: 53, damage: 4, duration: 0}
	drain := spell{name: "Drain", cost: 73, damage: 2, healing: 2, duration: 0}
	shield := spell{name: "Shield", cost: 113, duration: 6, armour: 7}
	poison := spell{name: "Poison", cost: 173, duration: 6, damage: 3}
	recharge := spell{name: "Recharge", cost: 229, duration: 5, recharge: 101}

	spells := []spell{magicMissile, drain, shield, poison, recharge}
	fmt.Println(spells)

	player := magicCharacter{50, 500, 0,true}
	boss := magicCharacter{51, 0, 9,false}

	fmt.Println(player)
	fmt.Println(boss)

	playerEffects := make(map[string]spell)
	bossEffects := make(map[string]spell)

	doTurn(player,boss,playerEffects,bossEffects,spells,0, true)

	fmt.Println("Min mana spent:", minManaSpent)
}

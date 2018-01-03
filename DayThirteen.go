package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

func getListPermutations(input []string) [][]string {
	permutations := make([][]string, 0)

	for i := 0; i < len(input); i++ {
		remainingItems := make([]string, 0)

		for j := 0; j < len(input); j++ {
			if input[i] != input[j] {
				remainingItems = append(remainingItems, input[j])
			}
		}

		if len(remainingItems) > 0 {
			followingPermutations := getListPermutations(remainingItems)
			for k := 0; k < len(followingPermutations); k++ {
				combination := make([]string, 0)
				combination = append(combination, input[i])
				combination = append(combination, followingPermutations[k]...)

				permutations = append(permutations, combination)
			}
		} else {

			combination := make([]string, 0)
			combination = append(combination, input[i])
			permutations = append(permutations, combination)
		}
	}

	return permutations
}

func parseHappinessMap(input string) map[string]map[string]int {

	happinessMap := make(map[string]map[string]int)

	inputLines := strings.Split(input, "\n")

	for i := 0; i < len(inputLines); i++ {
		line := inputLines[i]

		parts := strings.Split(line, " ")
		firstName := parts[0]
		secondName := parts[10][:len(parts[10])-1]
		happinessValues, exists := happinessMap[firstName]

		value, _ := strconv.Atoi(parts[3])

		if parts[2] == "lose" {
			value = value * -1
		}

		if !exists {
			happinessValues = make(map[string]int)
			happinessMap[firstName] = happinessValues
		}
		happinessValues[secondName] = value
	}

	return happinessMap
}

func getHappinessScoreForList(tableList []string, happinessMap map[string]map[string]int) int {

	score := 0

	for i := 0; i < len(tableList)-1; i++ {

		score += happinessMap[tableList[i]][tableList[i+1]]
		score += happinessMap[tableList[i+1]][tableList[i]]
	}

	score += happinessMap[tableList[len(tableList)-1]][tableList[0]]
	score += happinessMap[tableList[0]][tableList[len(tableList)-1]]

	return score
}

func DayThirteenExample() {

	fmt.Println("Day 13 - Example")

	people := []string{"Alice", "Bob", "Carol", "David"}
	combinations := getListPermutations(people)

	input := "Alice would gain 54 happiness units by sitting next to Bob.\nAlice would lose 79 happiness units by sitting next to Carol.\nAlice would lose 2 happiness units by sitting next to David.\nBob would gain 83 happiness units by sitting next to Alice.\nBob would lose 7 happiness units by sitting next to Carol.\nBob would lose 63 happiness units by sitting next to David.\nCarol would lose 62 happiness units by sitting next to Alice.\nCarol would gain 60 happiness units by sitting next to Bob.\nCarol would gain 55 happiness units by sitting next to David.\nDavid would gain 46 happiness units by sitting next to Alice.\nDavid would lose 7 happiness units by sitting next to Bob.\nDavid would gain 41 happiness units by sitting next to Carol."

	happinessMap := parseHappinessMap(input)
	score := 0

	for i := 0; i < len(combinations); i++ {
		newScore := getHappinessScoreForList(combinations[i], happinessMap)
		if newScore > score {
			score = newScore
		}

	}

	fmt.Println("Score:", score)
}

func DayThirteenPartOne() {

	fmt.Println("Day 13 - Part One")

	input := ReadFile("day13-input.txt")
	happinessMap := parseHappinessMap(input)

	people := make([]string, 0)

	for key, _ := range happinessMap {
		people = append(people, key)
	}

	combinations := getListPermutations(people)

	score := 0

	for i := 0; i < len(combinations); i++ {
		newScore := getHappinessScoreForList(combinations[i], happinessMap)
		if newScore > score {
			score = newScore
		}

	}

	fmt.Println("Score:", score)
}

func DayThirteenPartTwo() {

	fmt.Println("Day 13 - Part Two")

	input := ReadFile("day13-input.txt")
	happinessMap := parseHappinessMap(input)

	people := make([]string, 0)

	meMap := make(map[string]int)

	for key, _ := range happinessMap {
		people = append(people, key)
		meMap[key] = 0
	}

	people = append(people, "me")
	happinessMap["me"] = meMap

	combinations := getListPermutations(people)

	score := 0

	for i := 0; i < len(combinations); i++ {
		newScore := getHappinessScoreForList(combinations[i], happinessMap)
		if newScore > score {
			score = newScore
		}

	}

	fmt.Println("Score:", score)
}

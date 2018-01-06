package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

func getStoreCombinations(amountToStore int, containers []int) [][]int {

	combinations := make([][]int,0)
	for i:=0; i<len(containers);i++ {
		if containers[i] == amountToStore {
			combo := make([]int,1)
			combo[0] = containers[i]
			combinations = append(combinations,combo)
		} else if containers[i] < amountToStore {
			newAmount := amountToStore - containers[i]
			nextElement := i+1
			newCombos := getStoreCombinations(newAmount,containers[nextElement:])
			if len(newCombos) > 0 {
				for j:=0; j<len(newCombos); j++ {

					combo := make([]int,1)
					combo[0] = containers[i]
					combo = append(combo, newCombos[j]...)
					combinations = append(combinations, combo)
				}
			}

		}
	}

	return combinations
}

func DaySeventeenExample() {

	fmt.Println("Day 17 - Example")

	amountToStore := 25

	containers := []int {20,15,10,5,5}

	fmt.Println(amountToStore)
	fmt.Println(containers)

	fmt.Println("Combinations: ", len(getStoreCombinations(amountToStore, containers)))
}

func DaySeventeenPartOne() {

	fmt.Println("Day 17 - Part One")

	amountToStore := 150

	input := strings.Split(ReadFile("day17-input.txt"),"\n")
	containers := make([]int,len(input))

	for i := 0; i < len(input); i++ {
		size, _ := strconv.Atoi(input[i])
		containers[i] = size
	}
	fmt.Println(amountToStore)
	fmt.Println(containers)

	fmt.Println("Combinations: ", len(getStoreCombinations(amountToStore, containers)))
}

func DaySeventeenPartTwo() {

	fmt.Println("Day 17 - Part Two")

	amountToStore := 150

	input := strings.Split(ReadFile("day17-input.txt"),"\n")
	containers := make([]int,len(input))

	for i := 0; i < len(input); i++ {
		size, _ := strconv.Atoi(input[i])
		containers[i] = size
	}

	combinations := getStoreCombinations(amountToStore, containers)
	fmt.Println("Combinations: ", len(combinations))

	sizeMap := make(map[int]int)

	for j:=0; j<len(combinations); j++ {
		comboLength := len(combinations[j])

		sizeMap[comboLength] = sizeMap[comboLength]+1
	}

	fmt.Println(sizeMap)
}

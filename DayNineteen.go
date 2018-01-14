package adventofcode2015

import (
	"fmt"
	"strings"
)

type replacement struct {
	input  string
	output string
}

func performReplacement(starting string, replacements map[replacement]bool) map[string]bool {

	output := make(map[string]bool)

	for replacementRule, _ := range replacements {
		for i := 0; i < len(starting); i++ {
			endPos := i + len(replacementRule.input)
			endPosInRange := endPos <= len(starting)
			if endPosInRange && string(starting[i:endPos]) == replacementRule.input {
				newMolecule := starting[0:i]
				newMolecule += replacementRule.output
				newMolecule += starting[endPos:]
				output[newMolecule] = true
			}
		}
	}

	return output
}

func DayNineteenExample() {

	fmt.Println("Day 19 - Example")

	input := strings.Split("H => HO\nH => OH\nO => HH", "\n")

	replacements := make(map[replacement]bool)

	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], "=>")

		source := strings.Trim(parts[0], " ")
		dest := strings.Trim(parts[1], " ")
		replacements[replacement{source, dest}] = true

	}
	starting := "HOH"

	output := performReplacement(starting, replacements)
	fmt.Println(output)
}

func DayNineteenPartOne() {

	fmt.Println("Day 19 - Part One")

	input := strings.Split(ReadFile("day19-input.txt"), "\n")

	replacements := make(map[replacement]bool)

	for i := 0; i < len(input)-2; i++ {
		parts := strings.Split(input[i], "=>")

		source := strings.Trim(parts[0], " ")
		dest := strings.Trim(parts[1], " ")
		replacements[replacement{source, dest}] = true

	}
	starting := input[len(input)-1]

	output := performReplacement(starting, replacements)
	fmt.Println("Number of new molecules:", len(output))
}

func DayNineteenPartTwo() {

	fmt.Println("Day 19 - Part Two")

	input := strings.Split(ReadFile("day19-input.txt"), "\n")

	replacements := make(map[replacement]bool)

	for i := 0; i < len(input)-2; i++ {
		parts := strings.Split(input[i], "=>")

		source := strings.Trim(parts[0], " ")
		dest := strings.Trim(parts[1], " ")
		replacements[replacement{source, dest}] = true

	}
	aim := input[len(input)-1]

	starting := "e"
	fmt.Println(starting)
	for i := 1; i < 300; i++ {
		longestReplacement := 0
		currentReplace := replacement{"", ""}

		for replacementValue, _ := range replacements {
			if strings.Contains(aim, replacementValue.output) && longestReplacement < len(replacementValue.output) {
				currentReplace = replacementValue
				longestReplacement = len(replacementValue.output)
			}
		}

		if longestReplacement > 0 {
			fmt.Println("Going to use:", currentReplace)

			aim = strings.Replace(aim, currentReplace.output, currentReplace.input, 1)
			//	fmt.Println(aim)
			if aim == starting {
				fmt.Println("We are done at:", i)
				break
			}
		} else {
			fmt.Println("No more replacements")

			break
		}

	}

	fmt.Println(aim)
}

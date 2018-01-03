package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

func lookAndSay(input string) string {

	currentChar := ""
	currentCount := 0
	output := ""
	for j := 0; j < len(input); j++ {
		checkChar := input[j : j+1]
		if checkChar == currentChar {
			//Same char increase count
			currentCount++
		} else {
			//different char write previous and change
			if currentCount > 0 {
				output += strconv.Itoa(currentCount)
				output += currentChar
			}
			currentChar = checkChar
			currentCount = 1

		}
	}
	output += strconv.Itoa(currentCount)
	output += currentChar

	return output
}

func DayTenExample() {

	fmt.Println("Day 10 - Example")

	input := "1"

	for i := 0; i < 5; i++ {
		input = lookAndSay(input)
	}

	fmt.Println("Output:", input)
}

func DayTenPartOne() {

	fmt.Println("Day 10 - Part One")

	input := "1113122113"

	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}

	fmt.Println("Output:", len(input))
}

type element struct {
	number    int
	name      string
	pattern   string
	nextStage []string
}

func DayTenPartTwo() {

	fmt.Println("Day 10 - Part Two")

	input := "1113122113"

	fmt.Println(input)
	elementFile := strings.Split(ReadFile("day10-elements.txt"), "\n")

	elementList := make([]element, len(elementFile)+1)

	for i := 0; i < len(elementFile); i++ {

		elementParts := strings.Split(elementFile[i], "\t")
		elementNumber, _ := strconv.Atoi(elementParts[1])
		elementName := elementParts[2]
		elementPattern := elementParts[3]
		var nextStage []string

		elementList[elementNumber] = element{elementNumber, elementName, elementPattern, nextStage}
	}
	//for each element
	// look at next pattern
	// if simple next is next elment
	// else next is a list of elements including the next one

	elementNameMap := buildNextStateMap(elementList)
	var output []element
	output = append(output, elementNameMap["Fr"])
	for x := 0; x < 50; x++ {
		output = getNextState(output, elementNameMap)
	}
	finalPatternLength := 0

	for y := 0; y < len(output); y++ {
		finalPatternLength += len(output[y].pattern)
	}
	fmt.Println(finalPatternLength)
}

func getNextState(input []element, elementMap map[string]element) []element {

	output := make([]element, 0)

	for i := 0; i < len(input); i++ {
		inputElement := input[i]
		for j := 0; j < len(inputElement.nextStage); j++ {
			nextElement := elementMap[inputElement.nextStage[j]]
			output = append(output, nextElement)
		}
	}
	return output

}

func buildNextStateMap(elementList []element) map[string]element {

	elementNameMap := make(map[string]element)
	elementNames := make(map[string]element)

	for _, value := range elementList {
		elementNames[value.name] = value
	}

	for i := 1; i < len(elementList); i++ {

		currentElement := elementList[i]
		//fix pattern
		newPattern := ""
		var nextStageArray []string
		oldPatternGroups := strings.Split(currentElement.pattern, ".")
		for j := 0; j < len(oldPatternGroups); j++ {
			if _, exists := elementNames[oldPatternGroups[j]]; !exists {
				newPattern = oldPatternGroups[j]
				nextStageArray = append(nextStageArray, currentElement.name)
			} else {
				nextStageArray = append(nextStageArray, oldPatternGroups[j])
			}
		}

		if i+1 < len(elementList) {
			elementList[i+1].nextStage = nextStageArray
		}
		currentElement.pattern = newPattern
		elementNameMap[currentElement.name] = currentElement
	}

	firstElement := elementNameMap["H"]
	firstElement.nextStage = []string{"H"}
	elementNameMap["H"] = firstElement

	return elementNameMap
}

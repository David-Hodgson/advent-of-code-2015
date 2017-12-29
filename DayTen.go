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
	for j:= 0 ;j < len(input) ; j++ {
		checkChar := input[j:j+1]
		if checkChar == currentChar {
			//Same char increase count
			currentCount++
		} else {
			//different char write previous and change
			if (currentCount > 0) {
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

	for i := 0 ;i < 5; i++ {
		input = lookAndSay(input) 
	}

	fmt.Println("Output:", input)
}

func DayTenPartOne() {

	fmt.Println("Day 10 - Part One")

	input := "1113122113"

	for i := 0 ;i < 40; i++ {
		input = lookAndSay(input) 
	}

	fmt.Println("Output:", len(input))
}

type element struct {
	number int
	name string
	pattern string
}


func DayTenPartTwo() {

	fmt.Println("Day 10 - Part Two")

	input := "1113122113"
	//input = "13112221133211322112211213322112"

	fmt.Println(input)
	elementFile := strings.Split(ReadFile("day10-elements.txt"),"\n")

	elementList := make([]element, len(elementFile)+1)

	for i := 0 ; i < len(elementFile) ; i++ {

		elementParts := strings.Split(elementFile[i], "\t")
		elementNumber,_ := strconv.Atoi(elementParts[1])
		elementName := elementParts[2]
		elementPattern := elementParts[3]

		elementList[elementNumber] = element{elementNumber,elementName,elementPattern}
	}

	fmt.Println(elementList)

	//TODO parse elementList to build next elements

	//for each element
		// look at next pattern
		// if simple next is next elment
		// else next is a list of elements including the next one

	//TODO method to work through next states
}

func getNextState(input string, elementPatternMap, elementNameMap map[string]element, patternPosMap map[string]int, elementList []element) string {

	currentElement := elementPatternMap[input]
	currentPos := patternPosMap[input]

	fmt.Println(currentElement)
	fmt.Println(currentPos)

	if currentPos >= len(elementList) {
		fmt.Println("Last Element")
	} else {
		nextElement := elementList[currentPos+1]

		nextElementGroups := strings.Split(nextElement.pattern, ".")

		output := ""

		for i := 0 ; i < len(nextElementGroups); i++ {
			groupPattern := nextElementGroups[i]
			output += "."
			if otherElement, exists := elementNameMap[groupPattern]; exists {
				output += otherElement.pattern
			} else {
				output += groupPattern
			}

		}
		return output 
	}
	return input
}

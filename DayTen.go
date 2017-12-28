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
	name string
	pattern string
}


func DayTenPartTwo() {

	fmt.Println("Day 10 - Part Two")

	input := "1113122113"
	input = "13112221133211322112211213322112"
	elementFile := ReadFile("day10-elements.txt")
	elementFileList := strings.Split(elementFile, "\n")

	elementList := make([]element,len(elementFileList)) 

	for j := 0 ; j < len(elementFileList) ; j++ {
		elementParts := strings.Split(elementFileList[j], "\t")
		elementList[j] = element{elementParts[2],elementParts[3]}
	}

	elementPatternMap := make(map[string]element)
	elementNameMap := make(map[string]element)
	namePosMap := make(map[string]int)
	patternPosMap := make(map[string]int)

	for indx,element := range elementList {
		fmt.Println("indx:", indx)

		elementPatternMap[element.pattern] = element
		elementNameMap[element.name] = element

		namePosMap[element.name] = indx
		patternPosMap[element.pattern] = indx
	}

	fmt.Println("Pattern Map", elementPatternMap)
	fmt.Println("elementNameMap", elementNameMap)
	fmt.Println("namePosMap", namePosMap)
	fmt.Println("patternPosMap", patternPosMap)

	startPos := patternPosMap[input]

	fmt.Println("startPos:", startPos)

	output := ""
	stepCount := 2 
	for i := 0 ; i < stepCount; i++ {
		startPos++
		pattern := ""
		if startPos >= len(elementList) {
			pattern = "22"
		} else {
			nextElement := elementList[startPos]
			pattern = nextElement.pattern
		}

		patternParts := strings.Split(pattern, ".")

		output = ""

		for j := 0; j<len(patternParts) ; j++ {
			fmt.Println(patternParts[j])
			if _,exists := elementNameMap[patternParts[j]];exists {
				output += processGroup(patternParts[j], stepCount-i, elementNameMap, namePosMap, elementList)
			} else {
				output += patternParts[j]
			}
		}
	}
	fmt.Println("Output:", len(output))
}

func processGroup(name string, steps int, elementNameMap map[string]element, namePosMap map[string]int, elementList []element) string {

	startPos := namePosMap[name]

	fmt.Println("process Group start Pos:", startPos)
	output := ""
	for i := 0 ; i < steps; i++ {
		startPos++
		if startPos >= len(elementList) {
			startPos = len(elementList) -1
		}
		nextElement := elementList[startPos]

		patternParts := strings.Split(nextElement.pattern, ".")

		output = ""

		for j := 0; j<len(patternParts) ; j++ {
			fmt.Println(patternParts[j])
			if _,exists := elementNameMap[patternParts[j]];exists {
				fmt.Println("Name does not exist")
				output += processGroup(patternParts[j], steps-i, elementNameMap, namePosMap, elementList)
			} else {
				fmt.Println("Name does exist")
				output += patternParts[j]
			}
		}
		fmt.Println(output)

	}
	return output
}

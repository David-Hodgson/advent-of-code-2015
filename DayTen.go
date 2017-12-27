package adventofcode2015

import (
	"fmt"
	"strconv"
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

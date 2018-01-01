package adventofcode2015

import (
	"fmt"
	"strconv"
)

func DayTwelveExample() {

	fmt.Println("Day 12 - Example")

	input := "[1,2,3]"

	fmt.Println(input, ":", countJsonNumbers(input))

	input = "{\"a\":2,\"b\":4}"
	fmt.Println(input, ":", countJsonNumbers(input))

	input = "[[[3]]]"
	fmt.Println(input, ":", countJsonNumbers(input))

	input = "{\"a\":{\"b\":4},\"c\":-1}"
	fmt.Println(input, ":", countJsonNumbers(input))

	input = "{\"a\":[-1,1]}"
	fmt.Println(input, ":", countJsonNumbers(input))

	input = "[-1,{\"a\":1}]"
	fmt.Println(input, ":", countJsonNumbers(input))

	input = "[]"
	fmt.Println(input, ":", countJsonNumbers(input))

	input = "{}"
	fmt.Println(input, ":", countJsonNumbers(input))
}

func DayTwelvePartOne() {

	fmt.Println("Day 12 - Part One")
	input := ReadFile("day12-input.txt")

	fmt.Println("Total:", countJsonNumbers(input))
}

func DayTwelvePartTwo() {

	fmt.Println("Day 12 - Part Two")
	input := ReadFile("day12-input.txt")

	fmt.Println("Total:", countJsonNumbersExcludeRed(input))
}

const START = 0
const STRING = 1
const OBJECT = 2
const ARRAY = 3

func countJsonNumbers(json string) int {

	count := 0

	currentState := START

	stateStack := make([]int,0)
	stateStack = append(stateStack, currentState)

	value := ""

	for i :=0 ; i<len(json) ; i++ {

		currentChar := json[i]

		switch currentState{
			case ARRAY:
				switch currentChar{
					case ']':
						count += getIntFromString(value)
						value = ""
						currentState = stateStack[len(stateStack)-1]
						stateStack = stateStack[:len(stateStack)-1]
					case ',':
						count += getIntFromString(value)
						value = ""
					case '[':
						stateStack = append(stateStack, currentState)
						currentState = ARRAY
					case '{':
						stateStack = append(stateStack, currentState)
						currentState = OBJECT
					default:
						value += string(currentChar)
				}
			case OBJECT:
				switch currentChar {
					case '}':
						count += getIntFromString(value)
						value = ""
						currentState = stateStack[len(stateStack)-1]
						stateStack = stateStack[:len(stateStack)-1]
					case ',':
						count += getIntFromString(value)
						value = ""
					case ':' :
						value = ""
					case '{' :
						stateStack = append(stateStack, currentState)
						currentState = OBJECT
					case '[' :
						stateStack = append(stateStack, currentState)
						currentState = ARRAY
					default:
						value += string(currentChar)
				}
			case START:
				switch currentChar {
					case '{':
						stateStack = append(stateStack, currentState)
						currentState = OBJECT
						value = ""
					case '[':
						stateStack = append(stateStack, currentState)
						currentState = ARRAY
						value = ""
				}

		}
	}
	return count
}

type excludableCount struct {
	exclude bool
	count int
}

func countJsonNumbersExcludeRed(json string) int {

	currentState := START

	stateStack := make([]int,0)
	stateStack = append(stateStack, currentState)

	countStack := make([]excludableCount,0)
	countStack = append(countStack, excludableCount{false,0})

	value := ""

	for i :=0 ; i<len(json) ; i++ {

		currentChar := json[i]

		switch currentState{
			case ARRAY:
				switch currentChar{
					case ']':
						countStack[len(countStack)-1].count += getIntFromString(value)
						value = ""
						currentState = stateStack[len(stateStack)-1]
						stateStack = stateStack[:len(stateStack)-1]
						countStack[len(countStack)-2].count += countStack[len(countStack)-1].count
						countStack = countStack[:len(countStack)-1]
					case ',':
						countStack[len(countStack)-1].count += getIntFromString(value)
						value = ""
					case '[':
						stateStack = append(stateStack, currentState)

						countStack = append(countStack, excludableCount{false,0})
						currentState = ARRAY
					case '{':
						stateStack = append(stateStack, currentState)
						currentState = OBJECT
						countStack = append(countStack, excludableCount{false,0})
					default:
						value += string(currentChar)
				}
			case OBJECT:
				switch currentChar {
					case '}':
						if value == "\"red\"" {
							countStack[len(stateStack)-1].exclude = true
						}
						countStack[len(countStack)-1].count += getIntFromString(value)
						value = ""
						currentState = stateStack[len(stateStack)-1]
						stateStack = stateStack[:len(stateStack)-1]

						if !countStack[len(countStack)-1].exclude {
							countStack[len(countStack)-2].count += countStack[len(countStack)-1].count
						}
						countStack = countStack[:len(countStack)-1]
					case ',':

						if value == "\"red\"" {
							countStack[len(stateStack)-1].exclude = true
						}
						countStack[len(countStack)-1].count += getIntFromString(value)
						value = ""
					case ':' :
						value = ""
					case '{' :
						stateStack = append(stateStack, currentState)
						currentState = OBJECT
						countStack = append(countStack, excludableCount{false,0})
					case '[' :
						stateStack = append(stateStack, currentState)
						currentState = ARRAY

						countStack = append(countStack, excludableCount{false,0})
					default:
						value += string(currentChar)
				}
			case START:
				switch currentChar {
					case '{':
						stateStack = append(stateStack, currentState)
						currentState = OBJECT
						value = ""
						countStack = append(countStack, excludableCount{false,0})
					case '[':
						stateStack = append(stateStack, currentState)
						currentState = ARRAY
						countStack = append(countStack, excludableCount{false,0})
						value = ""
				}

		}
	}

	return countStack[0].count
}

func getIntFromString(input string) int {

	if value,err := strconv.Atoi(input) ; err == nil {
		return value
	}

	return 0
}

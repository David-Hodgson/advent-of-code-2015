package adventofcode2015

import (
	"fmt"
	"strings"
	"strconv"
)


func getRegisterValue(registers map[string]int, register string) int {

	register = strings.Trim(register, ",")
	if value,ok := registers[register]; ok {
		return value
	} else {
		parsedValue,_ := strconv.Atoi(register)
		return parsedValue
	}
}

func processInstructions(registers map[string]int, instructions []string) {

	fmt.Println("Going to process", len(instructions),"instructions")

	programCounter := 0

	for ; programCounter >=0 && programCounter < len(instructions); {
		incValue := 1

		instructionParts := strings.Split(instructions[programCounter]," ")

		opCode := instructionParts[0]

		switch (opCode) {
			case "inc":
				register := instructionParts[1]
				registers[register] = getRegisterValue(registers,register) + 1
			case "jio":
				register := instructionParts[1]
				if getRegisterValue(registers,register) == 1 {
					incValue = getRegisterValue(registers,instructionParts[2])
				}
			case "tpl":
				register := instructionParts[1]
				registers[register] = getRegisterValue(registers,register) * 3
			default:
				fmt.Println(opCode)
		}

		programCounter += incValue
	}
}

func DayTwentyThreeExample() {

	fmt.Println("Day 23 - Example")

	input := "inc a\njio a, +2\ntpl a\ninc a"
	instructions := strings.Split(input, "\n")

	registers := make(map[string]int)
	registers["a"] = 0
	registers["b"] = 0

	processInstructions(registers,instructions)

	fmt.Println(registers)
}


func DayTwentyThreePartOne() {

	fmt.Println("Day 23 - Part One")

	input := ReadFile("day23-input.txt") 
	instructions := strings.Split(input, "\n")

	registers := make(map[string]int)
	registers["a"] = 0
	registers["b"] = 0

	processInstructions(registers,instructions)

	fmt.Println(registers)
}

package adventofcode2015

import (
	"fmt"
	"strings"
	"strconv"
)

func getValue(circuit map[string]string, register string) string {

	bob := circuit[register]

	fmt.Println(bob)
	if strings.HasPrefix(bob, "NOT") {
		parts := strings.Split(bob, " ")
		value,_ := strconv.ParseUint(getValue(circuit, parts[1] ),10,16)
		return strconv.Itoa(int(^uint16(value)))

	} else {
		return bob
	}
}

func DaySevenExample() {


	input := "123 -> x\n456 -> y\nx AND y -> d\nx OR y -> e\nx LSHIFT 2 -> f\ny RSHIFT 2 -> g\nNOT x -> h\nNOT y -> i"

	fmt.Println(input)

	connections := strings.Split(input, "\n")

	positions := make(map[string]string)
	for i := 0 ; i < len(connections);i++ {

		//fmt.Println(connections[i])
		bits := strings.Split(connections[i], "-> ")

		fmt.Println("Input: ", bits[0])
		//fmt.Println("Output: ", bits[1])
		positions[bits[1]] = strings.Trim(bits[0], " ")
	}
	fmt.Println(positions)

	fmt.Println("d: ", getValue(positions, "d"))
	fmt.Println("e: ", getValue(positions, "e"))
	fmt.Println("f: ", getValue(positions, "f"))
	fmt.Println("g: ", getValue(positions, "g"))
	fmt.Println("h: ", getValue(positions, "h"))
	fmt.Println("i: ", getValue(positions, "i"))
	fmt.Println("x: ", getValue(positions, "x"))
	fmt.Println("y: ", getValue(positions, "y"))
}
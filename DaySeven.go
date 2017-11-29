package adventofcode2015

import (
	"fmt"
	"strings"
	"strconv"
	"log"
)

var (
	limit = 50
	count = 0
)
func getValue(circuit map[string]string, register string) string {

	count++;

	if count > limit {
		log.Panic("Done enough")
	}

	fmt.Println("Getting register: ", register)
	bob := circuit[register]

	fmt.Println(bob)
//	return bob

	if bob == "" {
		fmt.Println("Getting ", register, " and it's empty")
	}
	if strings.HasPrefix(bob, "NOT") {
		parts := strings.Split(bob, " ")
		value,_ := strconv.ParseUint(getValue(circuit, parts[1] ),10,16)
		return strconv.Itoa(int(^uint16(value)))
	} else if strings.Contains(bob, "AND") {
		parts := strings.Split(bob, " ")
		value1,_ := strconv.ParseUint(getValue(circuit, parts[0] ),10,16)
		value2,_ := strconv.ParseUint(getValue(circuit, parts[2] ),10,16)
		return strconv.Itoa(int(uint16(value1) & uint16(value2)))
	} else if strings.Contains(bob, "OR") {
		parts := strings.Split(bob, " ")
		value1,_ := strconv.ParseUint(getValue(circuit, parts[0] ),10,16)
		value2,_ := strconv.ParseUint(getValue(circuit, parts[2] ),10,16)
		return strconv.Itoa(int(uint16(value1) | uint16(value2)))
	} else if strings.Contains(bob, "LSHIFT") {
		parts := strings.Split(bob, " ")
		value1,_ := strconv.ParseUint(getValue(circuit, parts[0] ),10,16)
		shift,_ := strconv.ParseUint(parts[2],10,16)
		return strconv.Itoa(int(uint16(value1) << shift))
	} else if strings.Contains(bob, "RSHIFT") {
		parts := strings.Split(bob, " ")
		value1,_ := strconv.ParseUint(getValue(circuit, parts[0] ),10,16)
		shift,_ := strconv.ParseUint(parts[2],10,16)
		return strconv.Itoa(int(uint16(value1) >> shift))
	} else {
		fmt.Println("In else")
		if _,exists := circuit[bob]; exists {
			fmt.Println("register ", bob, "exists")
			return getValue(circuit, bob)
		}
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
	//fmt.Println(positions)
	//
	//fmt.Println("d: ", getValue(positions, "d"))
	//fmt.Println("e: ", getValue(positions, "e"))
	//fmt.Println("f: ", getValue(positions, "f"))
	//fmt.Println("g: ", getValue(positions, "g"))
	//fmt.Println("h: ", getValue(positions, "h"))
	//fmt.Println("i: ", getValue(positions, "i"))
	//fmt.Println("x: ", getValue(positions, "x"))
	//fmt.Println("y: ", getValue(positions, "y"))

	circuit := make(map[string]string)


	for i := 0 ; i< 50; i++ {
		for key, value := range positions {
			//fmt.Println("Key:", key, "Value:", value)
			if shift, err := strconv.ParseUint(value, 10, 16); err == nil {
				fmt.Println("Key:", key, "Value:", value)
				fmt.Println(shift)

				updateCircuit(positions, key, value)
				circuit[key] = value
				delete(positions, key)

			}
		}
	}

	fmt.Println(positions)
	fmt.Println(circuit)


}

func DaySevenPartOne() {


	input := ReadFile("day7-input.txt")

	sep := "\n"
	if strings.Contains(input, "\r\n") {
		sep = "\r\n"
	}

	connections := strings.Split(input, sep)

	positions := make(map[string]string)
	for i := 0 ; i < len(connections);i++ {

		bits := strings.Split(connections[i], "-> ")
		positions[bits[1]] = strings.Trim(bits[0], " ")
	}

	//fmt.Println("b: ", getValue(positions, "b"))

	//fmt.Println("a: ", getValue(positions, "a"))
	//fmt.Println("lx: ", getValue(positions, "lx"))
	//fmt.Println("lw: ", getValue(positions, "lw"))
	//fmt.Println("lv: ", getValue(positions, "lv"))
	//fmt.Println("1: ", getValue(positions, "1"))
	//fmt.Println("lu: ", getValue(positions, "lu"))

	circuit := make(map[string]string)


	for i := 0 ; i< 202; i++ {
		for key, value := range positions {
			//fmt.Println("Key:", key, "Value:", value)
			if shift, err := strconv.ParseUint(value, 10, 16); err == nil {
				fmt.Println("Key:", key, "Value:", value)
				fmt.Println(shift)

				updateCircuit(positions, key, value)
				circuit[key] = value
				delete(positions, key)

			}
		}
	}

	fmt.Println(positions)
	fmt.Println(circuit)
	fmt.Println("b:", circuit["b"])
	fmt.Println("a:", circuit["a"])
	fmt.Println("lx:", circuit["lx"])

}


func DaySevenPartTwo() {


	input := ReadFile("day7-input.txt")

	sep := "\n"
	if strings.Contains(input, "\r\n") {
		sep = "\r\n"
	}

	connections := strings.Split(input, sep)

	positions := make(map[string]string)
	for i := 0 ; i < len(connections);i++ {

		bits := strings.Split(connections[i], "-> ")
		positions[bits[1]] = strings.Trim(bits[0], " ")
	}

	positions["b"] = "16076"

	circuit := make(map[string]string)


	for i := 0 ; i< 202; i++ {
		for key, value := range positions {
			if _, err := strconv.ParseUint(value, 10, 16); err == nil {
				updateCircuit(positions, key, value)
				circuit[key] = value
				delete(positions, key)

			}
		}
	}

	fmt.Println(positions)
	fmt.Println(circuit)
	fmt.Println("b:", circuit["b"])
	fmt.Println("a:", circuit["a"])
	fmt.Println("lx:", circuit["lx"])

}


func updateCircuit(circuit map[string]string, register string, replacement string) {

	for key, value := range circuit {

		if strings.HasPrefix(value, "NOT") {
			parts := strings.Split(value, " ")

			if register == parts[1] {
				parsedReplacement,_ := strconv.ParseUint(replacement,10,16)
				newValue := strconv.Itoa(int(^uint16(parsedReplacement)))
				circuit[key] = newValue
			}

		} else if strings.Contains(value, "AND") {
			parts := strings.Split(value, " ")

			if parts[0] == register || parts[2] == register {
				circuit[key] = updateAND(value, register, replacement)
			}
		} else if strings.Contains(value, "OR") {
			parts := strings.Split(value, " ")

			if parts[0] == register || parts[2] == register {
				circuit[key] = updateOR(value, register, replacement)
			}
		} else if strings.Contains(value, "LSHIFT") {
			parts := strings.Split(value, " ")

			if register == parts[0] {
				shift,_ := strconv.ParseUint(parts[2],10,16)
				value1,_ := strconv.ParseUint(replacement,10,16)
				newValue := strconv.Itoa(int(uint16(value1) << shift))
				circuit[key] = newValue
			}

		} else if strings.Contains(value, "RSHIFT") {
			parts := strings.Split(value, " ")

			if register == parts[0] {
				shift,_ := strconv.ParseUint(parts[2],10,16)
				value1,_ := strconv.ParseUint(replacement,10,16)
				newValue := strconv.Itoa(int(uint16(value1) >> shift))
				circuit[key] = newValue
			}
		}

	}
}

func updateAND(oldValue string, register string, replacement string) string {

	parts := strings.Split(oldValue, " ")

	if parts[0] == register {
		parts[0] = replacement
	}

	if parts[2] == register {
		parts[2] = replacement
	}

	if value1, err := strconv.ParseUint(parts[0], 10, 16); err == nil {
		if value2, err2 := strconv.ParseUint(parts[2], 10, 16); err2 == nil {
			return strconv.Itoa(int(uint16(value1) & uint16(value2)))
		}
	}

	return parts[0] + " AND " + parts[2]

}

func updateOR(oldValue string, register string, replacement string) string {

	parts := strings.Split(oldValue, " ")

	if parts[0] == register {
		parts[0] = replacement
	}

	if parts[2] == register {
		parts[2] = replacement
	}

	if value1, err := strconv.ParseUint(parts[0], 10, 16); err == nil {
		if value2, err2 := strconv.ParseUint(parts[2], 10, 16); err2 == nil {
			return strconv.Itoa(int(uint16(value1) | uint16(value2)))
		}
	}

	return parts[0] + " OR " + parts[2]

}
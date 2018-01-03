package adventofcode2015

import (
	"bytes"
	"fmt"
	"strings"
)

func getStringLengths(input string) (memoryLength, codeLength int) {

	codeLength = len(input)
	count := 0
	for i := 1; i < len(input)-1; i++ {

		if input[i] == '\\' {

			if input[i+1] == '\\' {
				count++
				i++
			} else if input[i+1] == '"' {
				count++
				i++
			} else if input[i+1] == 'x' {
				count++
				i = i + 3
			}
		} else {
			count++
		}

	}

	memoryLength = count

	//fmt.Println(input)
	//fmt.Println("Memory Size:", memoryLength)
	//fmt.Println("Code Size:", codeLength)

	return memoryLength, codeLength
}

func encodeString(input string) string {

	fmt.Println(input)

	var str bytes.Buffer

	str.WriteString("\"")
	for i := 0; i < len(input); i++ {

		if input[i] == '"' {
			str.WriteString("\\\"")
		} else if input[i] == '\\' {
			str.WriteString("\\\\")
		} else {
			str.WriteByte(input[i])
		}
	}
	str.WriteString("\"")
	fmt.Println(str.String())
	return str.String()
}

func DayEightExample() {

	e1 := "\"\""

	memLength, codeLength := getStringLengths(e1)

	fmt.Println(e1)
	fmt.Println("Memory Length:", memLength)
	fmt.Println("Code Length:", codeLength)

	e2 := "\"abc\""

	memLength, codeLength = getStringLengths(e2)

	fmt.Println(e2)
	fmt.Println("Memory Length:", memLength)
	fmt.Println("Code Length:", codeLength)

	e3 := "\"aaa\\\"aaa\""

	memLength, codeLength = getStringLengths(e3)

	fmt.Println(e3)
	fmt.Println("Memory Length:", memLength)
	fmt.Println("Code Length:", codeLength)

	e4 := "\"\\x27\""

	memLength, codeLength = getStringLengths(e4)

	fmt.Println(e4)
	fmt.Println("Memory Length:", memLength)
	fmt.Println("Code Length:", codeLength)
}

func DayEightExample2() {

	e1 := "\"\""
	e2 := "\"abc\""
	e3 := "\"aaa\\\"aaa\""
	e4 := "\"\\x27\""

	stringList := []string{e1, e2, e3, e4}

	originalMemoryCount := 0
	orginalCodeCount := 0

	newMemoryCount := 0
	newCodeCount := 0

	for i := 0; i < len(stringList); i++ {
		iMemCount, iCodeCount := getStringLengths(stringList[i])
		originalMemoryCount += iMemCount
		orginalCodeCount += iCodeCount

		encodedString := encodeString(stringList[i])
		iMemCount, iCodeCount = getStringLengths(encodedString)
		newMemoryCount += iMemCount
		newCodeCount += iCodeCount
	}

	fmt.Println("Original Memory Size:", originalMemoryCount)
	fmt.Println("Original Code Size:", orginalCodeCount)

	originalDiff := orginalCodeCount - originalMemoryCount

	fmt.Println("Original Difference:", originalDiff)

	fmt.Println("New Memory Size:", newMemoryCount)
	fmt.Println("New Code Size:", newCodeCount)

	newDiff := newCodeCount - newMemoryCount

	fmt.Println("New Difference:", newDiff)

	codeDiff := newCodeCount - orginalCodeCount

	fmt.Println("Code Difference:", codeDiff)
}

func DayEightPartOne() {

	input := ReadFile("day8-input.txt")

	stringList := strings.Split(input, "\n")

	memoryCount := 0
	codeCount := 0

	for i := 0; i < len(stringList); i++ {
		iMemCount, iCodeCount := getStringLengths(stringList[i])
		memoryCount += iMemCount
		codeCount += iCodeCount
	}

	fmt.Println("Memory Size:", memoryCount)
	fmt.Println("Code Size:", codeCount)

	diff := codeCount - memoryCount

	fmt.Println("Difference:", diff)
}

func DayEightPartTwo() {

	input := ReadFile("day8-input.txt")

	stringList := strings.Split(input, "\n")

	originalMemoryCount := 0
	orginalCodeCount := 0

	newMemoryCount := 0
	newCodeCount := 0

	for i := 0; i < len(stringList); i++ {
		iMemCount, iCodeCount := getStringLengths(stringList[i])
		originalMemoryCount += iMemCount
		orginalCodeCount += iCodeCount

		encodedString := encodeString(stringList[i])
		iMemCount, iCodeCount = getStringLengths(encodedString)
		newMemoryCount += iMemCount
		newCodeCount += iCodeCount
	}

	fmt.Println("Original Memory Size:", originalMemoryCount)
	fmt.Println("Original Code Size:", orginalCodeCount)

	originalDiff := orginalCodeCount - originalMemoryCount

	fmt.Println("Original Difference:", originalDiff)

	fmt.Println("New Memory Size:", newMemoryCount)
	fmt.Println("New Code Size:", newCodeCount)

	newDiff := newCodeCount - newMemoryCount

	fmt.Println("New Difference:", newDiff)

	codeDiff := newCodeCount - orginalCodeCount

	fmt.Println("Code Difference:", codeDiff)
}

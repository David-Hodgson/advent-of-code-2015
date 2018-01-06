package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

func DaySixteenExample() {

	fmt.Println("Day 16 - Example")

}

func DaySixteenPartOne() {

	fmt.Println("Day 16 - Part One")

	mfcsamOutput := strings.Split("children: 3\ncats: 7\nsamoyeds: 2\npomeranians: 3\nakitas: 0\nvizslas: 0\ngoldfish: 5\ntrees: 3\ncars: 2\nperfumes: 1", "\n")
	mfcsamMap := make(map[string]int)

	for i := 0; i < len(mfcsamOutput); i++ {
		presentParts := strings.Split(mfcsamOutput[i], ":")

		presentName := presentParts[0]
		presentCount, _ := strconv.Atoi(presentParts[1][1:])

		mfcsamMap[presentName] = presentCount
	}

	input := ReadFile("day16-input.txt")

	sueList := strings.Split(input, "\n")

	matchName := ""
	for k := 0 ; k < len(sueList); k++ {

		sue := sueList[k]
		name := sue[0:strings.Index(sue, ":")]
		match := true 

		for present, count := range mfcsamMap {

			if strings.Contains(sue, present) {
				if !strings.Contains(sue, fmt.Sprintf("%s: %d", present, count)){
					match =  false
					break
				}
			}
		}

		if match {
			matchName = name
		}
	}

	fmt.Println("Matching Sue:", matchName)
}

func buildPresentMap(presentInput []string) map[string]int {

	mfcsamMap := make(map[string]int)

	for i := 0; i < len(presentInput); i++ {
		presentParts := strings.Split(presentInput[i], ":")

		presentName := strings.Trim(presentParts[0], " ")
		presentCount, _ := strconv.Atoi(presentParts[1][1:])

		mfcsamMap[presentName] = presentCount
	}

	return mfcsamMap
}

func DaySixteenPartTwo() {

	fmt.Println("Day 16 - Part Two")

	mfcsamOutput := strings.Split("children: 3\ncats: 7\nsamoyeds: 2\npomeranians: 3\nakitas: 0\nvizslas: 0\ngoldfish: 5\ntrees: 3\ncars: 2\nperfumes: 1", "\n")
	mfcsamMap := buildPresentMap(mfcsamOutput)

	input := ReadFile("day16-input.txt")

	sueList := strings.Split(input, "\n")

	matchName := ""
	for k := 0 ; k < len(sueList); k++ {

		sue := sueList[k]
		name := sue[0:strings.Index(sue, ":")]
		presentList := strings.Split(sue[len(name) + 2:],",")
		suePresentMap := buildPresentMap(presentList)

		match := true

		for present, count := range mfcsamMap {

			if  value, exists := suePresentMap[present] ; exists {
				if present == "cats" || present == "trees" {
					if value <= count {
						match = false
						break
					}

				} else if  present == "pomeranians" || present == "goldfish" {
					if value >= count {
						match = false
						break
					}
				} else {
					if value != count {
						match =  false
						break
					}
				}
			}
		}

		if match {
			matchName = name
		}
	}

	fmt.Println("Matching Sue:", matchName)
}

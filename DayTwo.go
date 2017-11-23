package adventofcode2015

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)


func calculatePaperRequired(l,w,h int) int {
	
	paperRequired := 2*l*w + 2*w*h + 2*h*l

	sides := [] int {l,w,h}	
	sort.Ints(sides)

	paperRequired += sides[0]*sides[1]
	return paperRequired
}

func calculateRibbonRequired(l,w,h int) int{

	sides := [] int {l,w,h}
	sort.Ints(sides)
	
	ribbonRequired := sides[0]*2 + sides[1] * 2
	ribbonRequired += l*w*h

	return ribbonRequired
}

func DayTwoPartOne() {

	input := ReadFile("day2-input.txt")

	presents := strings.Split(input, "\n")

	totalPaper := 0

	for i := 0; i<len(presents); i++ {
		if len(strings.Trim(presents[i], " ")) > 0 {

			dimensions := strings.Split(strings.Trim(presents[i],"\r"),"x")
		
			l,err := strconv.Atoi(dimensions[0])
			w,err := strconv.Atoi(dimensions[1])
			h,err := strconv.Atoi(dimensions[2])
			if err != nil {
				fmt.Println("error", err)
			}
	
			fmt.Println(dimensions)
			fmt.Println(calculatePaperRequired(l,w,h))
			totalPaper += calculatePaperRequired(l,w,h)		
		}
	}

	fmt.Println("Total Paper: ", totalPaper)
}

func DayTwoPartTwo() {


	input := ReadFile("day2-input.txt")

	presents := strings.Split(input, "\n")

	totalRibbon := 0

	for i := 0; i<len(presents); i++ {
		if len(strings.Trim(presents[i], " ")) > 0 {

			dimensions := strings.Split(strings.Trim(presents[i],"\r"),"x")
		
			l,err := strconv.Atoi(dimensions[0])
			w,err := strconv.Atoi(dimensions[1])
			h,err := strconv.Atoi(dimensions[2])
			if err != nil {
				fmt.Println("error", err)
			}
	
			fmt.Println(dimensions)
			fmt.Println(calculatePaperRequired(l,w,h))
			totalRibbon += calculateRibbonRequired(l,w,h)		
		}
	}

	fmt.Println("Total Ribbon: ", totalRibbon)
}

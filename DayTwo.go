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

func DayTwoPartOne() {

	input := ReadFile("day2-input.txt")

	presents := strings.Split(input, "\n")

	totalPaper := 0

	for i := 0; i<len(presents); i++ {
		if len(strings.Trim(presents[i], " ")) > 0 {

			dimensions := strings.Split(presents[i],"x")
		
			l,err := strconv.Atoi(dimensions[0])
			w,err := strconv.Atoi(dimensions[1])
			h,err := strconv.Atoi(dimensions[2])
			if err != nil {
				fmt.Println("error")
			}
	
			fmt.Println(dimensions)
			fmt.Println(calculatePaperRequired(l,w,h))
			totalPaper += calculatePaperRequired(l,w,h)		
		}
	}

	fmt.Println("Total Paper: ", totalPaper)
}

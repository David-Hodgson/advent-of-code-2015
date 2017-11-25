package adventofcode2015

import (
	"fmt"
	//"strconv"
	//"strings"
)

func DayThreePartOne() {

	var x int = 0
	var y int = 0
	positions := make(map[string]int)

	input := ReadFile("day3-input.txt")

	for _,rune := range input {

		switch rune { 
			case 'v':
			 y--
			 break
			case '^':
			 y++
			 break
			case '>':
			 x++
			 break
			case '<':
			 x--
			 break	
		}
		positions[fmt.Sprintf("x",x,"y",y)] = 1	
	}

	fmt.Println("Size of map ", len(positions)+1)
}

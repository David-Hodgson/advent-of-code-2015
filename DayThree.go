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



func DayThreePartTwo() {


	positions := make(map[string]int)
	positions[fmt.Sprintf("x",0,"y",0)] = 1	

	input := ReadFile("day3-input.txt")

	for i :=0; i<2;i++ {
	

		var x int = 0
		var y int = 0

		for j :=i ; j < len(input); j=j+2 {

		switch input[j] { 
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

	}
	fmt.Println("Size of map ", len(positions))
}



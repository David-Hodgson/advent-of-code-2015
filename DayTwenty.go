package adventofcode2015

import (
	"fmt"
)

func DayTwentyExample() {

	fmt.Println("Day 20 - Example")

	for i :=1; i < 10; i++ {
		fmt.Println("house ",i,":", getHouseScore(i))
	}
}

func getFactors(input int) []int {

	factors := make([]int,0)

	for i :=1; i <= input; i++ {

		if input % i == 0 {
			factors = append(factors,i)
		}
	}
	return factors
}

func getHouseScore(houseNumber int) int {

	factors := getFactors(houseNumber)
	total := 0

	for i := 0; i < len(factors); i++ {
		total += 10 * factors[i]
	}
	return total
}

func DayTwentyPartOne() {

	fmt.Println("Day 20 - Part One")
	input:= 29000000

	fmt.Println(input)

	houseNumber := 0
	for i:=input/40;i>0 ;i-- {

		score := getHouseScore(i)
		if score >= input {
			fmt.Println("House number",i,"gets at least",input,"presents")
			houseNumber = i
		}
	}

	fmt.Println("Lowest house number:", houseNumber)
}

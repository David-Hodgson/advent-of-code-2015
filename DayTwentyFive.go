package adventofcode2015

import (
	"fmt"
)

func buildGrid(size int) [][]int64 {

	grid := make([][]int64, size)

	for i:=0;i<size;i++ {
		grid[i] = make([]int64,size)
	}
	var value int64 = 20151125

	bottomValue := 0
	row,col := 0,0

	for ; ; {
		grid[col][row] = value

		//calculate next value
		value = (value * 252533) % 33554393

		//calculate next pos

		if col == size -1 && row == 0 {
			break
		}

		if row > 0 {
			row -= 1
			col += 1
		} else {
			//at the top, so
			//reset to first column
			col = 0
			bottomValue++
			row = bottomValue
		}

	}

	return grid
}

func getGridValue(targetRow,targetCol int) int64 {

	var value int64 = 20151125

	bottomValue :=1
	row,col := 1,1

	for ; ; {
		//calculate next value
		if row==(targetRow) && col==(targetCol) {
			break
		}
		value = (value * 252533) % 33554393

		if row > 1 {
			row -= 1
			col += 1
		} else {
			//at the top, so
			//reset to first column
			col = 1
			bottomValue++
			row = bottomValue
		}

	}

	return value
}

func printGrid(grid [][]int64) {

	for row :=0; row<len(grid); row++ {

		for col :=0; col<len(grid); col++ {
			if grid[col][row] > 0 {
				fmt.Print(grid[col][row], "\t")
			}
		}
		fmt.Println()
	}
}

func DayTwentyFiveExample() {

	fmt.Println("Day 25 - Example")


	grid := buildGrid(6)

	fmt.Println(grid)
	printGrid(grid)

	fmt.Println("1,1:", getGridValue(1,1))
	fmt.Println("6,6:", getGridValue(6,6))
}

func DayTwentyFivePartOne() {

	fmt.Println("Day 25 - Part One")

	targetRow := 2947
	targetCol := 3029

	value := getGridValue(targetRow,targetCol) 

	fmt.Println("Target Value:", value)
}

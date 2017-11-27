package adventofcode2015

import (
	"fmt"
)


func showGrid(grid [][]bool) {


	for i :=0 ; i< len(grid); i++ {
		for j :=0 ; j <len(grid[i]); j++ {

			if grid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}

		}

		fmt.Print("\n")
	}
}

func DaySixPartOne() {

	const gridWidth = 10
	const gridHeight = 10

	fmt.Println("Day Six")

	fmt.Println("Grid Width: ", gridWidth)
	fmt.Println("Grid Height: ", gridHeight)

	var grid [][]bool

	grid = make([][]bool,gridHeight)

	for i :=0 ; i< gridHeight ; i++ {
		grid[i] = make([]bool, gridWidth)
	}
	
	fmt.Println("grid size: ", len(grid))

	showGrid(grid)
}

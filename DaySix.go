package adventofcode2015

import (
	"fmt"
)

func setPixel(grid [][]bool, x,y int, state bool) {
	grid[y][x] = state
}

func togglePixel(grid [][]bool, x,y int) {
	grid[y][x] = !grid[y][x]
}

func turnPixelOn(grid [][]bool, x,y int, state bool) {
	setPixel(grid,x,y,true)
}

func turnPixelOff(grid [][]bool, x,y int, state bool) {
	setPixel(grid,x,y,false)
}

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

	//grid[9][5] = true
	setPixel(grid,5,9,true)
	fmt.Println("grid size: ", len(grid))

	showGrid(grid)
}

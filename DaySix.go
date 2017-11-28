package adventofcode2015

import (
	"fmt"
	"strings"
	"strconv"
)

func setPixel(grid [][]bool, x,y int, state bool) {
	grid[y][x] = state
}

func togglePixel(grid [][]bool, x,y int) {
	grid[y][x] = !grid[y][x]
}

func turnPixelOn(grid [][]bool, x,y int) {
	setPixel(grid,x,y,true)
}

func turnPixelOff(grid [][]bool, x,y int) {
	setPixel(grid,x,y,false)
}

func superBrightenPixel(grid [][]int, x,y int) {
	grid[y][x] = grid[y][x] + 2;
}

func brightenPixel(grid [][]int, x,y int) {
	grid[y][x] = grid[y][x] + 1;
}
func dimPixel(grid [][]int, x,y int) {
	if grid[y][x] >0 {
		grid[y][x] = grid[y][x] - 1
	}
}

func countPixelBrightness(grid [][]int) int {
	count := 0

	for i :=0 ; i< len(grid); i++ {
		for j :=0 ; j <len(grid[i]); j++ {
				count += grid[i][j]
		}
	}

	return count
}

func countLitPixels(grid [][]bool) int {
	count := 0

	for i :=0 ; i< len(grid); i++ {
		for j :=0 ; j <len(grid[i]); j++ {

			if grid[i][j] {
				count++	
			}
		}
	}

	return count
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

func parsePoints(command string) (minX,minY,maxX,maxY int) {

	parts := strings.Split(command, " ")
	topLeft := strings.Split(parts[0], ",")
	bottomRight := strings.Split(parts[2], ",")

	minX,_ = strconv.Atoi(topLeft[0])
	minY,_ = strconv.Atoi(topLeft[1])
	maxX,_ = strconv.Atoi(bottomRight[0])
	maxY,_ = strconv.Atoi(bottomRight[1])

	return minX,minY,maxX,maxY
}

func DaySixPartOne() {

	const gridWidth = 1000
	const gridHeight = 1000

	fmt.Println("Day Six")

	fmt.Println("Grid Width: ", gridWidth)
	fmt.Println("Grid Height: ", gridHeight)

	var grid [][]bool

	grid = make([][]bool,gridHeight)

	for i :=0 ; i< gridHeight ; i++ {
		grid[i] = make([]bool, gridWidth)
	}

	input := ReadFile("day6-input.txt")

	steps := strings.Split(input, "\n")

	for i := 0; i< len(steps); i++ {

		step := steps[i]

		//fmt.Println(step)
	
		if strings.HasPrefix(step,"toggle") {
			toggle := step[len("toggle "):]
			minX,minY,maxX,maxY := parsePoints(toggle)
			for  y := minY ; y<=maxY ; y++ {
				for x := minX ; x <= maxX ; x++ {
					togglePixel(grid,x,y)
				}
			}
		}

		if strings.HasPrefix(step,"turn on") {
			toggle := step[len("turn on "):]
			minX,minY,maxX,maxY := parsePoints(toggle)
			for  y := minY ; y<=maxY ; y++ {
				for x := minX ; x <= maxX ; x++ {
					turnPixelOn(grid,x,y)
				}
			}
		}

		if strings.HasPrefix(step,"turn off") {
			toggle := step[len("turn off "):]
			minX,minY,maxX,maxY := parsePoints(toggle)
			for  y := minY ; y<=maxY ; y++ {
				for x := minX ; x <= maxX ; x++ {
					turnPixelOff(grid,x,y)
				}
			}
		}
	}

	//showGrid(grid)
	fmt.Println("Lit pixels: ", countLitPixels(grid))
}

func DaySixPartTwo() {

	const gridWidth = 1000
	const gridHeight = 1000

	fmt.Println("Day Six - Part Two")

	fmt.Println("Grid Width: ", gridWidth)
	fmt.Println("Grid Height: ", gridHeight)

	var grid [][]int

	grid = make([][]int,gridHeight)

	for i :=0 ; i< gridHeight ; i++ {
		grid[i] = make([]int, gridWidth)
	}

	input := ReadFile("day6-input.txt")

	steps := strings.Split(input, "\n")

	for i := 0; i< len(steps); i++ {

		step := steps[i]

		//fmt.Println(step)

		if strings.HasPrefix(step,"toggle") {
			toggle := step[len("toggle "):]
			minX,minY,maxX,maxY := parsePoints(toggle)
			for  y := minY ; y<=maxY ; y++ {
				for x := minX ; x <= maxX ; x++ {
					superBrightenPixel(grid,x,y)
				}
			}
		}

		if strings.HasPrefix(step,"turn on") {
			toggle := step[len("turn on "):]
			minX,minY,maxX,maxY := parsePoints(toggle)
			for  y := minY ; y<=maxY ; y++ {
				for x := minX ; x <= maxX ; x++ {
					brightenPixel(grid,x,y)
				}
			}
		}

		if strings.HasPrefix(step,"turn off") {
			toggle := step[len("turn off "):]
			minX,minY,maxX,maxY := parsePoints(toggle)
			for  y := minY ; y<=maxY ; y++ {
				for x := minX ; x <= maxX ; x++ {
					dimPixel(grid,x,y)
				}
			}
		}
	}

	//showGrid(grid)
	fmt.Println("Brightness: ", countPixelBrightness(grid))
}

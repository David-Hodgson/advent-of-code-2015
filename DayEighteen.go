package adventofcode2015

import (
	"fmt"
	"strings"
)

func parseMatrix(input string) [][]bool {

	lines := strings.Split(input, "\n")

	matrix := make([][]bool, len(lines))

	for i := 0; i < len(lines); i++ {

		row := make([]bool, len(lines[i]))

		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '#' {
				row[j] = true
			}
		}
		matrix[i] = row
	}

	return matrix
}

func animateMatrix(inputMatrix [][]bool) [][]bool {

	outputMatrix := make([][]bool, len(inputMatrix))

	for y := 0; y < len(inputMatrix); y++ {
		outputRow := make([]bool, len(inputMatrix[y]))

		for x := 0; x < len(inputMatrix[y]); x++ {
			outputRow[x] = getNextCellState(x, y, inputMatrix)
		}
		outputMatrix[y] = outputRow
	}
	return outputMatrix
}

func getNextCellState(x, y int, matrix [][]bool) bool {
	litCount := 0

	//Row Above
	if y > 0 {
		//Above left
		if x > 0 && matrix[y-1][x-1] {
			litCount++
		}
		//above center
		if matrix[y-1][x] {
			litCount++
		}
		//above right
		if x < len(matrix[y])-1 && matrix[y-1][x+1] {
			litCount++
		}
	}

	//Same Row

	//Center left
	if x > 0 && matrix[y][x-1] {
		litCount++
	}
	//center right
	if x < len(matrix[y])-1 && matrix[y][x+1] {
		litCount++
	}

	//Row Below

	if y < len(matrix)-1 {
		//below left
		if x > 0 && matrix[y+1][x-1] {
			litCount++
		}
		//below center
		if matrix[y+1][x] {
			litCount++
		}
		//below right
		if x < len(matrix[y])-1 && matrix[y+1][x+1] {
			litCount++
		}
	}

	if matrix[y][x] {
		return litCount == 2 || litCount == 3
	} else {
		return litCount == 3
	}

}

func getLitCount(matrix [][]bool) int {

	litCount := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] {
				litCount++
			}
		}
	}

	return litCount
}

func DayEighteenExample() {

	fmt.Println("Day 18 - Example")

}

func DayEighteenPartOne() {

	fmt.Println("Day 18 - Part One")

	input := ReadFile("day18-input.txt")

	matrix := parseMatrix(input)

	for i := 0; i < 100; i++ {
		matrix = animateMatrix(matrix)
	}

	fmt.Println("Lit Count:", getLitCount(matrix))
}

func DayEighteenPartTwo() {

	fmt.Println("Day 18 - Part Two")

	input := ReadFile("day18-input.txt")

	matrix := parseMatrix(input)

	for i := 0; i < 100; i++ {
		matrix = animateMatrix(matrix)

		matrix[0][0] = true
		matrix[0][len(matrix[0])-1] = true
		matrix[len(matrix)-1][0] = true
		matrix[len(matrix)-1][len(matrix[len(matrix)-1])-1] = true
	}

	fmt.Println("Lit Count:", getLitCount(matrix))
}

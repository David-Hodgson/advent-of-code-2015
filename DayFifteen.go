package adventofcode2015

import (
	"fmt"
	"strings"
	"strconv"
)

type ingredient struct {
	name string
	capacity,durability, flavor, texture, calories int
}

func parseIngredient(input string) ingredient {

	ingredientParts := strings.Split(input, " ")

	name := ingredientParts[0][:len(ingredientParts[0])-1]
	capacity, _ := strconv.Atoi(ingredientParts[2][:len(ingredientParts[2])-1])
	durability, _ := strconv.Atoi(ingredientParts[4][:len(ingredientParts[4])-1])
	flavour, _ := strconv.Atoi(ingredientParts[6][:len(ingredientParts[6])-1])
	texture, _ := strconv.Atoi(ingredientParts[8][:len(ingredientParts[8])-1])
	calories, _ := strconv.Atoi(ingredientParts[10])

	return ingredient{name,capacity,durability,flavour,texture,calories}
}

func getRecipeScore(recipeMap map[ingredient]int) int {

	capScore := 0
	durScore := 0
	flavScore := 0
	texScore := 0

	for key,value := range recipeMap {
		capScore += value * key.capacity
		durScore += value * key.durability
		flavScore += value * key.flavor
		texScore += value * key.texture
	}

	if capScore < 0 {
		capScore = 0
	}

	if durScore < 0 {
		durScore = 0
	}

	if flavScore <0 {
		flavScore = 0
	}

	if texScore < 0 {
		texScore = 0
	}

	return capScore * durScore * flavScore * texScore
}

func getRecipeCalories(recipeMap map[ingredient]int) int {

	calorieScore := 0

	for key,value := range recipeMap {
		calorieScore += value * key.calories
	}

	return calorieScore
}

func DayFifteenExample() {

	fmt.Println("Day 15 - Example")

	input := "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8\nCinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3"

	ingredients := strings.Split(input, "\n")

	ingredientList := make([]ingredient,len(ingredients))
	for i := 0; i < len(ingredients); i++ {

		ingredient := parseIngredient(ingredients[i])
		ingredientList[i] = ingredient
	}

	fmt.Println(ingredientList)

	maxIngredientCount := 100

	maxScore := 0
	for i := 0; i <= maxIngredientCount ; i++ {
		remainingIngredientCount := maxIngredientCount - i

		for j:=0; j<= remainingIngredientCount; j++ {
			recipe := make(map[ingredient]int)

			recipe[ingredientList[0]] = i
			recipe[ingredientList[1]] = j

			score := getRecipeScore(recipe)
			//fmt.Println("i:", i, "j:", j, "Score:", score)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println("Score:", maxScore)
}


func DayFifteenPartOne() {

	fmt.Println("Day 15 - Part One")

	input := ReadFile("day15-input.txt")

	ingredients := strings.Split(input, "\n")

	ingredientList := make([]ingredient,len(ingredients))
	for i := 0; i < len(ingredients); i++ {

		ingredient := parseIngredient(ingredients[i])
		ingredientList[i] = ingredient
	}

	fmt.Println(ingredientList)

	maxIngredientCount := 100

	maxScore := 0

	for i := 0; i <= maxIngredientCount ; i++ {
		remainingIngredientCount1 := maxIngredientCount - i

		for j:=0; j<= remainingIngredientCount1; j++ {
			remainingIngredientCount2:= remainingIngredientCount1 -j

			for k := 0; k <= remainingIngredientCount2; k++ {
				remainingIngredientCount3 := remainingIngredientCount2 - k

				for l :=0; l <= remainingIngredientCount3; l++ {

					recipe := make(map[ingredient]int)

					recipe[ingredientList[0]] = i
					recipe[ingredientList[1]] = j
					recipe[ingredientList[2]] = k
					recipe[ingredientList[3]] = l

					score := getRecipeScore(recipe)
					//fmt.Println("i:", i, "j:", j, "Score:", score)
					if score > maxScore {
						maxScore = score
					}
				}
			}
		}
	}

	fmt.Println("Score:", maxScore)
}


func DayFifteenPartTwo() {

	fmt.Println("Day 15 - Part Two")

	input := ReadFile("day15-input.txt")

	ingredients := strings.Split(input, "\n")

	ingredientList := make([]ingredient,len(ingredients))
	for i := 0; i < len(ingredients); i++ {

		ingredient := parseIngredient(ingredients[i])
		ingredientList[i] = ingredient
	}

	fmt.Println(ingredientList)

	maxIngredientCount := 100

	maxScore := 0

	for i := 0; i <= maxIngredientCount ; i++ {
		remainingIngredientCount1 := maxIngredientCount - i

		for j:=0; j<= remainingIngredientCount1; j++ {
			remainingIngredientCount2:= remainingIngredientCount1 -j

			for k := 0; k <= remainingIngredientCount2; k++ {
				remainingIngredientCount3 := remainingIngredientCount2 - k

				for l :=0; l <= remainingIngredientCount3; l++ {

					recipe := make(map[ingredient]int)

					recipe[ingredientList[0]] = i
					recipe[ingredientList[1]] = j
					recipe[ingredientList[2]] = k
					recipe[ingredientList[3]] = l

					if getRecipeCalories(recipe) == 500 {
						score := getRecipeScore(recipe)
						//fmt.Println("i:", i, "j:", j, "Score:", score)
						if score > maxScore {
							maxScore = score
						}
					}
				}
			}
		}
	}

	fmt.Println("Score:", maxScore)
}

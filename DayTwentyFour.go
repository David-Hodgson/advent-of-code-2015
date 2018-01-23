package adventofcode2015

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func getBoxesTotalWeight(boxes []int) int {
	totalWeight := 0

	for i :=0; i <len(boxes); i++ {
		totalWeight += boxes[i]
	}

	return totalWeight
}

//Think max size of this set is 2248
var validComboMap map[string][]int = make(map[string][]int)

func buildBoxCombinationsToWardsTarget(currentCombo, boxes []int, currentWeight, targetWeight int) [][]int{

	combinations := make([][]int,0)

	//if len(validComboMap) == 2248 {
	//	return combinations
	//}

	sort.Ints(boxes[:])

	for i:= len(boxes)-1; i>=0; i-- {
		if currentWeight + boxes[i] > targetWeight {
			continue
		}

		remainingBoxes := make([]int,len(boxes))
		copy(remainingBoxes, boxes)
		remainingBoxes = append(remainingBoxes[:i], remainingBoxes[i+1:]...)

		newCombo := make([]int,len(currentCombo)+1)
		copy(newCombo, currentCombo)
		newCombo[len(currentCombo)] = boxes[i]

		newWeight := currentWeight + boxes[i]
		weightDiff := targetWeight - newWeight

		if newWeight == targetWeight {
			sort.Ints(newCombo)
			//fmt.Println("Found a combination:", newCombo)
			validComboMap[getMapKey(newCombo)] = newCombo
			if len(validComboMap) > 2000 {
				fmt.Println("valid combos:", len(validComboMap))
			}
			combinations = append(combinations, newCombo)
		} else if weightDiff > getBoxesTotalWeight(remainingBoxes) {
			continue
		}else {
			otherCombos := buildBoxCombinationsToWardsTarget(newCombo, remainingBoxes, newWeight, targetWeight)
			combinations = append(combinations, otherCombos...)
		}
	}

	return combinations
}

var comboCache map[int]map[string]map[string][]int = make(map[int]map[string]map[string][]int)

func getCombinationsForLength(availableBoxes []int, targetLength int) map[string][]int {
	sort.Ints(availableBoxes[:])

	if targetLengthCache,ok := comboCache[targetLength]; ok {

		if combos,ok1 := targetLengthCache[getMapKey(availableBoxes)]; ok1 {
			return combos
		}
	}
	combinationMap := make(map[string][]int)

	remainingLength := targetLength -1
	for i:=0;i<len(availableBoxes);i++ {

		if remainingLength > 0 {
			remainingBoxes := make([]int,len(availableBoxes))
			copy(remainingBoxes, availableBoxes)

			remainingBoxes = append(remainingBoxes[:i], remainingBoxes[i+1:]...)
			nextCombos := getCombinationsForLength(remainingBoxes, remainingLength)

			for _,combo := range nextCombos {
				combination := make([]int,1)
				combination[0] = availableBoxes[i]
				combination = append(combination,combo...)
				sort.Ints(combination[:])
				combinationMap[getMapKey(combination)] = combination
			}
		} else {
			combination := make([]int,1)
			combination[0] = availableBoxes[i]
			sort.Ints(combination[:])
			combinationMap[getMapKey(combination)] = combination
		}
	}

	if lengthCache,ok := comboCache[targetLength]; ok {
		lengthCache[getMapKey(availableBoxes)] = combinationMap
		comboCache[targetLength] = lengthCache

	} else {
		lengthCache = make(map[string]map[string][]int)
		lengthCache[getMapKey(availableBoxes)] = combinationMap
		comboCache[targetLength] = lengthCache
	}
	return combinationMap
}

func buildCombinations(availableBoxes []int, targetWeight int) [][]int {
	combinations := make([][]int,0)

	for i :=0; i < len(availableBoxes); i++ {
		newTargetWeight := targetWeight - availableBoxes[i]

		if newTargetWeight < 0 {
			continue
		} else if newTargetWeight == 0 {
			combo := make([]int,1)
			combo[0] = availableBoxes[i]
			combinations = append(combinations, combo)
		} else {
			remainingBoxes := make([]int,len(availableBoxes))
			copy(remainingBoxes, availableBoxes)
			remainingBoxes = append(remainingBoxes[:i], remainingBoxes[i+1:]...)
			otherCombos := buildCombinations(remainingBoxes, newTargetWeight)

			for j:=0; j<len(otherCombos); j++ {
				combo := make([]int,1)
				combo[0] = availableBoxes[i]
				combo = append(combo, otherCombos[j]...)
				combinations = append(combinations, combo)
			}
		}

	}

	return combinations

}

func stripUsedBoxes(allBoxes,usedBoxes []int) []int{
	unusedBoxes := make([]int, len(allBoxes)-len(usedBoxes))
	pos :=0
	for i:=0;i<len(allBoxes); i++ {
		used := false
		for j:=0;j<len(usedBoxes);j++ {
			if allBoxes[i]== usedBoxes[j] {
				used=true
			}

		}

		if !used {
			unusedBoxes[pos] = allBoxes[i]
			pos++
		}
	}

	return unusedBoxes
}

func getSmallestSet(set1,set2,set3 []int) (size,qe int) {

	size1 := len(set1)
	qe1 := getQE(set1)
	size2 := len(set2)
	qe2 := getQE(set2)
	size3 := len(set3)
	qe3 := getQE(set3)

	minSize := size1
	qe = qe1

	if size2 < minSize {
		minSize = size2
		qe = qe2
	} else if size2 == minSize {
		if qe2 < qe {
			qe = qe2
		}
	}

	if size3 < minSize {
		minSize = size3
		qe = qe3
	} else if size3 == minSize {
		if qe3 < qe {
			qe = qe3
		}
	}

	return minSize,qe
}

func getQE(boxes []int) int {
	qe := 1

	for i:=0; i<len(boxes); i++ {
		qe *= boxes[i]
	}

	return qe
}

func getMapKey(boxes []int) string {
	key := "["

	for i:=0;i<len(boxes);i++ {
		key += strconv.Itoa( boxes[i])
		key += ","
	}

	key += "]"

	return key
}

func reduceCombinations(combinations[][]int) map[string][]int {

	combinMap := make(map[string][]int)

	for i:=0;i<len(combinations);i++ {
		sort.Ints(combinations[i][:])
		combinMap[getMapKey(combinations[i])]=combinations[i]
	}
	return combinMap
}

func getMinQE(boxes []int, combinationTarget int) int {

	combinations := buildCombinations(boxes, combinationTarget)
	combo1Map := reduceCombinations(combinations)
	fmt.Println("First level count:", len(combo1Map))
	minQE := -1
	smallestSetSize := len(boxes)
	for _,combo1 := range combo1Map {
		set1 := combo1
		boxes2 := stripUsedBoxes(boxes, set1)

		combinations2 := buildCombinations(boxes2, combinationTarget)
		combo2Map := reduceCombinations(combinations2)
		for _, combo2 := range combo2Map {
			set2 := combo2
			boxes3 := stripUsedBoxes(boxes2,set2)

			combinations3 := buildCombinations(boxes3,combinationTarget)
			combo3Map := reduceCombinations(combinations3)
			for _,combo3 := range combo3Map {
				set3 := combo3
				size,qe := getSmallestSet(set1,set2,set3)

				if size < smallestSetSize {
					smallestSetSize = size
					minQE = qe
				} else if size == smallestSetSize {
					if minQE == -1 || qe < minQE {
						minQE = qe
					}
				}
			}
		}
	}

	return minQE
}

func doCombinationsOverlap(set1, set2 []int) bool {

	overlap := false

	fmt.Println("Comparing sets")
	fmt.Println("\tSet1:", set1)
	fmt.Println("\tSet2:", set2)

	for i :=0; i<len(set1) && !overlap; i++ {
		for j:=0;j<len(set2) && !overlap;j++ {

			if set1[i]== set2[j] {
				overlap = true
				break
			}
		}
	}

	return overlap
}

func getMinQEFromValidCombos(combinations [][]int) int {

	minQE := -1
	smallestSetSize := 28

	for i:=0; i<len(combinations);i++ {
		set1 := combinations[i]

		for j:=0; j<len(combinations); j++ {
			set2 := combinations[j]
			if i==j || doCombinationsOverlap(set1,set2){
				continue
			}

			fmt.Println("Found 2 distinct sets")
			fmt.Println(set1)
			fmt.Println(set2)

			for k:=0;k<len(combinations); k++ {
				set3 := combinations[k]
				if k==i || k==j || doCombinationsOverlap(set1,set3) || doCombinationsOverlap(set2, set3) {
					continue
				}

				fmt.Println("Found 3 distnct sets")
				fmt.Println(set3)
				size,qe := getSmallestSet(set1,set2,set3)

				fmt.Println("minSize:", size)
				fmt.Println("minQE:", qe)
				if size < smallestSetSize {
					smallestSetSize = size
					minQE = qe
				} else if size == smallestSetSize {
					if minQE == -1 || qe < minQE {
						minQE = qe
					}
				}
			}
		}
	}

	return minQE
}

func DayTwentyFourExample() {

	fmt.Println("Day 24 - Example")

	boxes := []int  {1,2,3,4,5,7,8,9,10,11}

	fmt.Println(boxes)

	totalWeight := 0

	for i :=0; i <len(boxes); i++ {
		totalWeight += boxes[i]
	}

	fmt.Println("total Weight:", totalWeight)

	combinationTarget := totalWeight / 3
	fmt.Println("Combination target:", combinationTarget)

	minQE := getMinQE(boxes,combinationTarget)
	fmt.Println("Min QE:", minQE)
}

func DayTwentyFourPartOne() {

	fmt.Println("Day 24 - Part One - Take 2")

	input := strings.Split(ReadFile("day24-input.txt"),"\n")

	boxes := make([]int, len(input))

	for i:=0; i<len(input); i++ {

		value,err := strconv.Atoi(input[i])

		if err != nil {
			fmt.Println(err)
		}
		boxes[i] = value
	}

	fmt.Println(boxes)
	fmt.Println("Number of boxes:", len(boxes))
	totalWeight := 0

	for i :=0; i <len(boxes); i++ {
		totalWeight += boxes[i]
	}

	fmt.Println("total Weight:", totalWeight)

	combinationTarget := totalWeight / 3
	fmt.Println("Combination target:", combinationTarget)

	var emptyCombo =[]int {}

	fmt.Println("Generating combinations")
	buildBoxCombinationsToWardsTarget(emptyCombo,boxes,0,combinationTarget)
	fmt.Println("Combination generatation done")
	validCombos := make([][]int,len(validComboMap))
	pos := 0
	for _,combo := range validComboMap {
		validCombos[pos] = combo
		pos++
	}

	fmt.Println("Getting minimun QE")
	minQE := getMinQEFromValidCombos(validCombos)
	fmt.Println("Min QE:", minQE)
	fmt.Println("Finished")
}

func DayTwentyFourPartOneOriginal() {

	fmt.Println("Day 24 - Part 1")

	input := strings.Split(ReadFile("day24-input.txt"),"\n")

	boxes := make([]int, len(input))

	for i:=0; i<len(input); i++ {

		value,_ := strconv.Atoi(input[i])
		boxes[i] = value
	}

	fmt.Println(boxes)

	totalWeight := 0

	for i :=0; i <len(boxes); i++ {
		totalWeight += boxes[i]
	}

	fmt.Println("total Weight:", totalWeight)

	combinationTarget := totalWeight / 3
	fmt.Println("Combination target:", combinationTarget)

	minQE := getMinQE(boxes,combinationTarget)
	fmt.Println("Min QE:", minQE)
}

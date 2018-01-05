package adventofcode2015

import (
	"fmt"
	"strconv"
	"strings"
)

type deerStats struct {
	name                           string
	speed, flyingTime, restingTime int
}

func parseReindeerStats(reindeerStats string) deerStats {

	statParts := strings.Split(reindeerStats, " ")

	name := statParts[0]
	speed, _ := strconv.Atoi(statParts[3])
	flyingTime, _ := strconv.Atoi(statParts[6])
	restingTime, _ := strconv.Atoi(statParts[13])

	return deerStats{name, speed, flyingTime, restingTime}
}

func getDistanceForTime(stats deerStats, time int) int {

	distance := 0
	cycleTime := stats.flyingTime + stats.restingTime

	fullCycleCount := time / cycleTime
	fullCycleDistance := (stats.speed * stats.flyingTime) * fullCycleCount

	timeLeft := time - (cycleTime * fullCycleCount)

	distance = fullCycleDistance

	if timeLeft >= stats.flyingTime {
		distance += (stats.speed * stats.flyingTime)
	} else {
		distance += (stats.speed * timeLeft)
	}

	return distance
}

func DayFourteenExample() {

	fmt.Println("Day 14 - Example")

	input := "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.\nDancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds."

	reindeerStats := strings.Split(input, "\n")

	maxDistance := 0
	time := 1000
	for i := 0; i < len(reindeerStats); i++ {
		stats := parseReindeerStats(reindeerStats[i])
		deerDistance := getDistanceForTime(stats, time)

		if deerDistance > maxDistance {
			maxDistance = deerDistance
		}
	}

	fmt.Println("Max Distance:", maxDistance)

}

func DayFourteenPartOne() {

	fmt.Println("Day 14 - Part One")

	input := ReadFile("day14-input.txt")

	reindeerStats := strings.Split(input, "\n")

	maxDistance := 0
	time := 2503

	for i := 0; i < len(reindeerStats); i++ {
		stats := parseReindeerStats(reindeerStats[i])
		deerDistance := getDistanceForTime(stats, time)

		if deerDistance > maxDistance {
			maxDistance = deerDistance
		}
	}

	fmt.Println("Max Distance:", maxDistance)

}

func DayFourteenPartTwo() {

	fmt.Println("Day 14 - Part 2")

	input := ReadFile("day14-input.txt")
	reindeerStats := strings.Split(input, "\n")

	time := 2503

	scoreMap := make(map[string]int)
	for x := 1; x <= time; x++ {
		maxTimeDistance := 0
		maxDeerName := make([]string, 0)
		for i := 0; i < len(reindeerStats); i++ {
			stats := parseReindeerStats(reindeerStats[i])
			deerDistance := getDistanceForTime(stats, x)
			if deerDistance > maxTimeDistance {
				maxTimeDistance = deerDistance
				maxDeerName = []string{stats.name}
			} else if deerDistance == maxTimeDistance {
				maxDeerName = append(maxDeerName, stats.name)
			}
		}

		for j := 0; j < len(maxDeerName); j++ {
			scoreMap[maxDeerName[j]] = scoreMap[maxDeerName[j]] + 1
		}
	}

	fmt.Println(scoreMap)
}

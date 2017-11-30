package adventofcode2015

import (
	"fmt"
	"strings"
	"strconv"
)

type Route struct {
	Start, End string
}

func DayNineExample() {

	fmt.Println("Day Nine Example")

	input := "London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141"

	routeStrings := strings.Split(input, "\n")

	destinations := make(map[string]bool)
	routes := make(map[Route]int)

	for i :=0 ; i< len(routeStrings) ; i++ {

		fmt.Println(routeStrings[i])

		parts := strings.Split(routeStrings[i], " ")

		startPoint := parts[0]
		endPoint := parts[2]
		distance,_ := strconv.Atoi(parts[4])

		destinations[startPoint] = true
		destinations[endPoint] = true
		fmt.Println(startPoint, endPoint, distance)

		routes[Route{startPoint, endPoint}] = distance
	}

	fmt.Println(destinations)
	fmt.Println(routes)

	distance := 0
	var localdestinations = []string {}

	for value,_ := range destinations {
		localdestinations = append(localdestinations, value)
	}

	for i :=0 ; i< len(localdestinations); i++ {


		distance = calculateDistanceToAllDestinations(localdestinations[i], localdestinations, routes, distance)
	}

	fmt.Println(distance)
}


func calculateDistanceToAllDestinations(startPoint string, destinations []string, distanceMap map[Route]int, currentDistance int ) int {

	fmt.Println("Starting at", startPoint)

	for i := 0 ; i <len(destinations); i++ {

		fmt.Println("Calculating distance to", destinations[i])

		//TODO distance could be either way
		fmt.Println("Distance:", distanceMap[Route{startPoint, destinations[i]}])

		var localdestinations = []string {}

		for value := range destinations {
			if destinations[value] != destinations[i] {
				localdestinations = append(localdestinations, destinations[value])
			}
		}

		fmt.Println(localdestinations)
	}

	return 0;
}




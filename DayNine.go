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
		routes[Route{endPoint, startPoint}] = distance
	}

	fmt.Println(destinations)
	fmt.Println(routes)

	routesCombos :=	generateRoutes(destinations)

	for j := 0 ; j < len(routesCombos) ; j++ {

		fmt.Println(routesCombos[j])
		fmt.Println(getRouteLength(routesCombos[j], routes))
	}
}

func DayNinePartOne() {

	fmt.Println("Day Nine - Part One")

	input := ReadFile("day9-input.txt")

	routeStrings := strings.Split(input, "\n")

	destinations := make(map[string]bool)
	routes := make(map[Route]int)

	for i :=0 ; i< len(routeStrings) ; i++ {

		parts := strings.Split(routeStrings[i], " ")

		startPoint := parts[0]
		endPoint := parts[2]
		distance,_ := strconv.Atoi(parts[4])

		destinations[startPoint] = true
		destinations[endPoint] = true

		routes[Route{startPoint, endPoint}] = distance
		routes[Route{endPoint, startPoint}] = distance
	}

	fmt.Println(destinations)
	fmt.Println(routes)

	routesCombos :=	generateRoutes(destinations)

	minDistance := -1
	for j := 0 ; j < len(routesCombos) ; j++ {

		dist := getRouteLength(routesCombos[j], routes)

		if dist < minDistance || minDistance == -1 {
			minDistance = dist
		}
	}

	fmt.Println("Minimum Distance:", minDistance)
}

func DayNinePartTwo() {

	fmt.Println("Day Nine - Part Two")

	input := ReadFile("day9-input.txt")

	routeStrings := strings.Split(input, "\n")

	destinations := make(map[string]bool)
	routes := make(map[Route]int)

	for i :=0 ; i< len(routeStrings) ; i++ {

		parts := strings.Split(routeStrings[i], " ")

		startPoint := parts[0]
		endPoint := parts[2]
		distance,_ := strconv.Atoi(parts[4])

		destinations[startPoint] = true
		destinations[endPoint] = true

		routes[Route{startPoint, endPoint}] = distance
		routes[Route{endPoint, startPoint}] = distance
	}

	fmt.Println(destinations)
	fmt.Println(routes)

	routesCombos :=	generateRoutes(destinations)

	minDistance := -1
	maxDistance := 0
	for j := 0 ; j < len(routesCombos) ; j++ {

		dist := getRouteLength(routesCombos[j], routes)

		if dist < minDistance || minDistance == -1 {
			minDistance = dist
		}

		if dist > maxDistance {
			maxDistance = dist
		}
	}

	fmt.Println("Minimum Distance:", minDistance)
	fmt.Println("Maximum Distance:", maxDistance)
}

func getRouteLength(route []string, distanceMap map[Route]int) int {

	distance := 0

	for i := 0; i < len(route) -1; i++ {

		route := Route{route[i], route[i+1]}

		distance += distanceMap[route]
	}
	return distance
}

func generateRoutes(destinations map[string]bool) [][]string {

	var routes [][]string 
	for destination,_ := range destinations {

		remainingDests := make(map[string]bool)

		for otherDest,_ := range destinations {
			if otherDest != destination {
				remainingDests[otherDest] = true
			}
		}

		if len(remainingDests) > 0 {

			//Keep tracking
			otherRoutes := generateRoutes(remainingDests)

			for i := 0 ; i < len(otherRoutes) ; i++ {

				var currentRoute []string
				currentRoute = append(currentRoute, destination)
				currentRoute = append(currentRoute, otherRoutes[i]...)
				routes = append(routes,currentRoute)
			}
		} else {
			var currentRoute []string
			currentRoute = append(currentRoute, destination)
			routes = append(routes,currentRoute)
		}

	}

	return routes 
}




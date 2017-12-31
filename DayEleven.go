package adventofcode2015

import (
	"fmt"
	"strings"
)

func isValidPassword(password string) bool {

	includesStraight := false
	twoPairs := false

	if strings.ContainsAny(password, "iol") {
		return false
	}

	pairOne := ""
	pairTwo := ""
	pairOneStart := 0
	pairTwoStart := 0

	for i:=0;i<len(password)-1;i++ {

		if i<len(password)-2 && password[i]+1 == password[i+1] && password[i]+2 == password[i+2] {
			includesStraight = true
		}

		if password[i] == password[i+1] {
			if pairOne == "" {
				pairOne = string(password[i])
				pairOneStart = i
			} else if pairTwo == "" {
				pairTwo = string(password[i])
				pairTwoStart = i
			} else {
				//Both are already set
				fmt.Println("Found a pair and both already set")
			}
		}
	}

	if (pairOne != pairTwo) && pairTwoStart > pairOneStart+1 {
		twoPairs = true
	}
	return includesStraight && twoPairs
}

func generateNextPassword(currentPassword string) string {


	newPassword := []rune(currentPassword)
	overflow := true 

	for i:=1; i<=len(newPassword) && overflow; i++ {

		currentChar := newPassword[len(newPassword)-i]
		currentChar = currentChar + 1

		if currentChar == 'o' || currentChar == 'i' || currentChar == 'l' {
			currentChar = currentChar +1
		}

		if currentChar == '{' {
			currentChar = 'a'
			overflow = true
		} else {
			overflow = false
		}
		newPassword[len(newPassword)-i] = currentChar 
	}
	return string(newPassword)
}

func DayElevenExample() {

	fmt.Println("Day 11 - Example")

	input := "hijklmmn"
	fmt.Println(input, ":", isValidPassword(input))

	input = "abbceffg"
	fmt.Println(input, ":", isValidPassword(input))

	input = "abbcegjk"
	fmt.Println(input, ":", isValidPassword(input))

	input = "hxbxwxba"
	fmt.Println(input, ":", isValidPassword(input))

	input = "ghjaabcc"
	fmt.Println(input, ":", isValidPassword(input))

	password := "aa"
	for i:=0; i< 126 ;i++ {
		password = generateNextPassword(password)
		fmt.Println(password)
	}
}

func DayElevenPartOne() {

	fmt.Println("Day 11 - Part One")

	currentPassword := "hxbxwxba"

	fmt.Println("Current Password:", currentPassword)

	for ; ; {

		currentPassword = generateNextPassword(currentPassword)
		fmt.Println(currentPassword)
		if isValidPassword(currentPassword) {
			break
		}
	}

	fmt.Println("New Password:", currentPassword)
}

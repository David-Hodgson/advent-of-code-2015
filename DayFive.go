package adventofcode2015

import (
	"regexp"
	"fmt"
	"strings"
)

func isSentenceNiceV1(sentence string) bool {

	nodisallowed := true
	if strings.Contains(sentence,"ab") || strings.Contains(sentence,"cd") || strings.Contains(sentence,"pq") || strings.Contains(sentence,"xy") {
		nodisallowed = false
	}

	threeVowels := false
	doubleletters := false

	if nodisallowed {
		threeVowels, _ = regexp.Match(`([aeiou]\w*){3}`, []byte(sentence));

		for i := 0; i<len(sentence)-1;i++ {

			if (sentence[i] == sentence[i+1]) {
				doubleletters = true
			}
		}
	}

	return threeVowels && doubleletters && nodisallowed
}

func isSentenceNiceV2(sentence string) bool {

	letterPairs := false
	repeats := false

	for i := 0; i<len(sentence)-2;i++ {

		if (sentence[i] == sentence[i+2]) {
			repeats = true
		}
	}

	pairs := make(map[string]int)
	for i := 0; i<len(sentence)-1;i++ {

		pair := sentence[i:i+2]

		if pos,exists := pairs[pair]; exists {
			if i > pos+1 {
				letterPairs = true
				break
			}
		} else {
			pairs[pair] = i
		}
	}

	return letterPairs && repeats
}

func DayFivePartOne() {

	input := ReadFile("day5-input.txt")

	sentences := strings.Split(input, "\n")

	niceCount := 0

	for i := 0; i<len(sentences); i++ {

		if isSentenceNiceV1(sentences[i]) {
			niceCount++
		}
	}

	fmt.Println("Nice Sentences: ",niceCount)
}

func DayFivePartTwo() {

	input := ReadFile("day5-input.txt")

	sentences := strings.Split(input, "\n")

	niceCount := 0

	for i := 0; i<len(sentences); i++ {

		if isSentenceNiceV2(sentences[i]) {
			niceCount++
		}
	}

	fmt.Println("Nice Sentences: ",niceCount)
}
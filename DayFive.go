package adventofcode2015

import (
	"regexp"
	"fmt"
	"strings"
)

func isSentenceNice(sentence string) bool {

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

func DayFivePartOne() {

	input := ReadFile("day5-input.txt")

	sentences := strings.Split(input, "\n")

	niceCount := 0

	for i := 0; i<len(sentences); i++ {

		if isSentenceNice(sentences[i]) {
			niceCount++
		}

	}

	fmt.Println("Nice Sentences: ",niceCount)
}

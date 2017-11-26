package adventofcode2015

import (
	"regexp"
	"fmt"
	"strings"
)

func DayFivePartOne() {

	input := "ugknbfddgicrmopn"

	fmt.Println(input)


	threeVowels,_ := regexp.Match(`([aeiou]\w*){3}`,[]byte(input));
	doubleletters,err := regexp.Match(`(.)\1`,[]byte(input)); 
	nodisallowed := true 

	if strings.Contains(input,"ab") || strings.Contains(input,"cd") || strings.Contains(input,"pq") || strings.Contains(input,"xy") {
		nodisallowed = false
	}


	fmt.Println(err)
	
	regex,err := regexp.Compile(`(.)\1`)
	fmt.Println(err)	
	fmt.Println(regex)	
	fmt.Println(threeVowels)
	fmt.Println(doubleletters)
	fmt.Println(nodisallowed)


}

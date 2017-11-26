package adventofcode2015

import (
	"fmt"
	"crypto/md5"
)

func DayFourPartOne() {

	input := "ckczppom"
	found := false
	value := 0

	for ; !found; {

		//fmt.Println("value: ",value)
		key := fmt.Sprintf("%s%d",input,value)
		//fmt.Println(key)
		hasher := md5.New()
		hasher.Write([]byte(key))
		firstFive := fmt.Sprintf("%x",hasher.Sum(nil))[0:5]
		//fmt.Println(firstFive)

		if firstFive == "00000" {
			found = true
			fmt.Println(key)	
		}

		value++
	}
}

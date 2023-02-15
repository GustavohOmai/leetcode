package main

import ("fmt"
		"strings")

func main() {

	s := "xix"
	
	output := romanToInt(s)
	fmt.Println(s)
	fmt.Println(output)
}

func romanToInt(s string) int {
	
	s= strings.ToUpper(s)

	valuesMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	lastNumber := 0
	actualNumber := 0
	total := 0

	for i := (len(s)-1); i >= 0; i-- {
		actualNumber = valuesMap[s[i]]

		if lastNumber > actualNumber {
			total -= actualNumber
		} else {
			total += actualNumber
		}
		lastNumber = actualNumber
	}


	return total
}

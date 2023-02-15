package main

import "fmt"

func main(){
	s:="abaabcapiusafdguyiolagsfduis"

	output:=lengthOfLongestSubstring(s)
	fmt.Println(s)
	fmt.Println(output)
}


func lengthOfLongestSubstring(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		set := make(map[byte]bool)
		length := 0
		for j := i; j < len(s); j++ {
			if set[s[j]] {
				break
			}

			set[s[j]] = true
			length++
		}

		if result < length {
			result = length
		}
	}

	return result
}
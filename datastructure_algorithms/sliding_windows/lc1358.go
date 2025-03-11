package main

import "fmt"

func numberOfSubstrings(s string) int {
	var freq [3]int
	total := 0
	l := 0

	for r := 0; r < len(s); r++ {
		freq[s[r]-'a']++

		for freq[0] > 0 && freq[1] > 0 && freq[2] > 0 {
			total += len(s) - r
			freq[s[l]-'a']--
			l++
		}
	}

	return total
}

func main() {
	s := "abcabc"
	result := numberOfSubstrings(s)
	fmt.Printf("number of substrings: %d\n", result)
}

// Count Vowels in a String

package main

import "fmt"

func countVowels(s string) int {
	count := 0

	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			count++
		}
	}

	return count
}

func main() {
	input := "hello world"
	result := countVowels(input)

	fmt.Println("Number of vowels:", result)
}

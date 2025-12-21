package main

import "fmt"

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	testStrings := []string{
		"hello",
		"Go",
		"12345",
		"Hello, 世界",
	}

	for _, s := range testStrings {
		fmt.Printf("Original: %q -> Reversed: %q\n", s, Reverse(s))
	}
}

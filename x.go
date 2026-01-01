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
	// Example usage
	original := "Hello, World!"
	reversed := Reverse(original)
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Reversed: %s\n", reversed)
}

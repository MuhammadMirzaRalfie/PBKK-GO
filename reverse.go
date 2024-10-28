package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Print("Masukkan nama (minimal 3 kata): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	words := strings.Fields(input)
	if len(words) < 3 {
		fmt.Println("Masukkan minimal 3 kata.")
		return
	}

	for _, word := range words {
		fmt.Print(reverseString(word) + " ")
	}
	fmt.Println()
}

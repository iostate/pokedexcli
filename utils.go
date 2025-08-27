package main

import "strings"

func cleanInput(text string) []string {
	// Lowercase entire string
	lowerCaseString := strings.ToLower(text)
	// Split by fields, splitting by 1 or more whitespace
	words := strings.Fields(lowerCaseString)
	return words
}

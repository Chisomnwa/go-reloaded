package functions

import (
	"strconv"
	"strings"
)

func handleHexAndBin(text string, tag string, base int) string {
	/*
	Takes in a string and searches for the word hex, and converts
	the number before it to decimal number
	*/

	// Use the parseInt() function to convert
	words := strings.Split(text, " ")

	for i := 0; i < len(words); i++ {
		// scans for (hex) and makes sure there's a word before it
		if words[i] == tag && i > 0 {
			// convert previous word from hex to decimal
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				// converts the number back to int
				words[i-1] = strconv.FormatInt(num, 10)
			}

			// remove "(hex)"
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
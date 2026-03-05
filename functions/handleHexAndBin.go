package functions

import (
	"strconv"
	"strings"
)

func handleHexAndBin(text string) string {
	words := strings.Split(text, " ")

	for i := 0; i < len(words); i++ {
		// check for hex
		if words[i] == "(hex)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}

		// check for bin
		if words[i] == "(bin)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

package functions

import (
	"strconv"
	"strings"
)

func handleBin(text string) string {

	// Use the parseInt() function to convert
	words := strings.Split(text, " ")

	for i := 0; i < len(words); i++ {
		// scans for bin and ensures there's a word before it
		if words[i] == "(bin)" && i > 0 {
					// convert previous word from binary to decimal
					num, err := strconv.ParseInt(words[i-1], 2, 64)
					if err == nil {
						// convert the number back to int
						words[i-1] = strconv.FormatInt(num, 10)
					}

					// remove "(bin)"
					words = append(words[:i], words[i+1:]...) // the thre ... unpacks the slice into individual elements
					i--
				}
			}

		return strings.Join(words, " ")
}
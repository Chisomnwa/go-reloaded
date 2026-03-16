package functions

import (
	"strconv"
	"strings"
)

func HandleCaseModifiers(text string) string {
	words := strings.Split(text, " ")

	for i := 0; i < len(words); i++ {
		// scan single word case
		if i > 0 && (words[i] == "(up)" || words[i] == "(low)" || words[i] == "(cap)") {
			switch words[i] {
			case "(up)":
				words[i-1] = strings.ToUpper(words[i-1])
			case "(low)":
				words[i-1] = strings.ToLower(words[i-1])
			case "(cap)":
				words[i-1] = Capitalize(words[i-1])
			}

			// remove (up), (low) and (cap)
			words = append(words[:i], words[i+1:]...)
			i--
		}

		// Let's treat multi word Case e.g (up N)
		if i < len(words) &&
			(strings.HasPrefix(words[i], "(up") ||
				strings.HasPrefix(words[i], "(low") ||
				strings.HasPrefix(words[i], "(cap")) {
			if i+1 >= len(words) {
				continue
			}

			// Let's extract modifier type
			modifier := strings.TrimPrefix(words[i], "(")
			modifier = strings.TrimSuffix(modifier, ",") // now we have "up", "low" or "cap"

			// Let's extract the number
			numStr := strings.TrimSuffix(words[i+1], ")")
			N, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			// Let's apply modifier to previous N words
			for j := 1; j <= N && i-j >= 0; j++ {
				switch modifier {
				case "up":
					words[i-j] = strings.ToUpper(words[i-j])
				case "low":
					words[i-j] = strings.ToLower(words[i-j])
				case "cap":
					words[i-j] = Capitalize(words[i-j])
				}
			}

			// Remove e.g "(up," "N)"
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

// Create a Capitalized function
func Capitalize(word string) string {
	if word == "" {
		return word
	}

	word = strings.ToLower(word)
	runes := []rune(word)

	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] - 32
	}
	return string(runes)
}

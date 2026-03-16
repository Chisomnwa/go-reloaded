package functions

import "strings"

func FixPunctuation(text string) string {
	words := strings.Split(text, " ")
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]
		punct := ",.!?:;"

		// -----------------------
		// Handle single quotes
		// -----------------------
		if word == "'" {
			i++

			// Collect all words until the closing quote
			var inner []string
			for i < len(words) && words[i] != "'" {
				w := words[i]

				// Strip leading punctuation and attach to previous inner word
				for len(w) > 0 && strings.ContainsRune(punct, rune(w[0])) {
					if len(inner) > 0 {
						inner[len(inner)-1] += string(w[0])
					}
					w = w[1:]
				}

				// Strip trailing punctuation
				endPunct := ""
				for len(w) > 0 && strings.ContainsRune(punct, rune(w[len(w)-1])) {
					endPunct = string(w[len(w)-1]) + endPunct
					w = w[:len(w)-1]
				}

				if w != "" {
					inner = append(inner, w+endPunct)
				} else if endPunct != "" && len(inner) > 0 {
					inner[len(inner)-1] += endPunct
				}
				i++
			}

			// Build the quoted string: attach opening quote to first inner word,
			// closing quote to last inner word — no spaces around the quotes.
			if len(inner) > 0 {
				inner[0] = "'" + inner[0]
				inner[len(inner)-1] = inner[len(inner)-1] + "'"
				result = append(result, inner...)
			}
			continue
		}

		// -----------------------
		// Handle punctuation at start
		// Example: ",but"
		// -----------------------
		for len(word) > 0 && strings.ContainsRune(punct, rune(word[0])) {
			if len(result) > 0 {
				result[len(result)-1] += string(word[0])
			}
			word = word[1:]
		}

		// -----------------------
		// Handle punctuation at end
		// Example: "raincoat."
		// -----------------------
		endPunct := ""
		for len(word) > 0 && strings.ContainsRune(punct, rune(word[len(word)-1])) {
			endPunct = string(word[len(word)-1]) + endPunct
			word = word[:len(word)-1]
		}

		// Add clean word
		if word != "" {
			result = append(result, word)
		}

		// Attach ending punctuation
		if endPunct != "" && len(result) > 0 {
			result[len(result)-1] += endPunct
		}
	}

	return strings.Join(result, " ")
}
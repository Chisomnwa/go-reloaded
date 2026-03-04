package functions

import "strings"

func fixAtoAn(text string) string {
	words := strings.Split(text, " ")

	for i := 0; i < len(words)-1; i++ {
		// Normalize to check: strip punctuation from the current word
		// so "a." or "a," still get caught
		current := strings.ToLower(strings.Trim(words[i], ",.!?:;"))

		if current == "a" {
			// Get the first letter of the next word (ignore any leading punctuation)
			next := strings.TrimLeft(words[i+1], ",.!?:;")
			if len(next) == 0 {
				continue
			}

			firstLetter := rune(strings.ToLower(next)[0])
			if strings.ContainsRune("aeiouh", firstLetter) {
				// Preserve the original casing: "a" -> "an", "A" -> "An"
				if words[i][0] == 'A' {
					words[i] = strings.Replace(words[i], "A", "An", 1)
				} else {
					words[i] = strings.Replace(words[i], "a", "an", 1)
				}
			}
		}
	}
	return strings.Join(words, " ")
}

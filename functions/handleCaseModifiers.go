package functions

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func handleCaseModifiers(text string) string {
	words := strings.Split(text, " ")

	// build a caser to work on (cap)
	caser := cases.Title(language.English)

	for i := 0; i < len(words); i++ {
		// scan words
		if words[i] == "(up)" || words[i] == "(low)" || words[i] == "cap" && i > 0 {
			modifier := words[i]

			if modifier == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
			} else if modifier == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
			} else if modifier == "(cap)" {
				words[i-1] = caser.String(words[i-1])
			}

			// remove (up), (low) and (cap)
			// I STOPPED HERE!!! TRYING TO CUT OUT THE ABOVE WORDS
			words = append(words[:1], words[i+1:])
		}
	}
}

package functions

func ProcessText(text string) string {
	text = handleHexAndBin(text, "(hex)", 16)
	text = handleHexAndBin(text, "(bin)", 2)
	// text = handleCaseModifiers(text)
	// text = fixPunctuation(text)
	// text = fixAtoAn(text)

	return text
}

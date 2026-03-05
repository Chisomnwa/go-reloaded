package functions

func ProcessText(text string) string {
	text = handleHexAndBin(text)
	// text = handleCaseModifiers(text)
	// text = fixPunctuation(text)
	// text = fixAtoAn(text)

	return text
}

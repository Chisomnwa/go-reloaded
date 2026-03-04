package functions

func ProcessText(text string) string {
	text = handleHex(text)
	text = handleBin(text)
	text = handleCaseModifiers(text)
	text = fixPunctuation(text)
	text = fixAtoAn(text)

	return text
}

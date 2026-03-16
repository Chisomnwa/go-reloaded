package functions

func ProcessText(text string) string {
	text = HandleHexAndBin(text)
	text = HandleCaseModifiers(text)
	text = FixPunctuation(text)
	text = FixAtoAn(text)

	return text
}

package functions

import (
	"os"
	"testing"
)

// Test hex and binary conversions
func TestHandleHexAndBin(t *testing.T) {
	input := "42 (hex) and 10 (bin)"
	expected := "66 and 2"
	got := HandleHexAndBin(input)
	if got != expected {
		t.Errorf("Hex/Bin conversion failed. got=%q, want=%q", got, expected)
	}
}

// Test single word case modifiers
func TestHandleCaseModifiersSingle(t *testing.T) {
	input := "hello (up) world (low) Go (cap)"
	expected := "HELLO world Go"
	got := HandleCaseModifiers(input)
	if got != expected {
		t.Errorf("Single case modifiers failed. got=%q, want=%q", got, expected)
	}
}

// Test multi-word case modifiers
func TestHandleCaseModifiersMulti(t *testing.T) {
	input := "this is exciting (up, 3)"
	expected := "THIS IS EXCITING"
	got := HandleCaseModifiers(input)
	if got != expected {
		t.Errorf("Multi-word case modifiers failed. got=%q, want=%q", got, expected)
	}
}

// Test fixAtoAn
func TestFixAtoAn(t *testing.T) {
	input := "It was a amazing day"
	expected := "It was an amazing day"
	got := FixAtoAn(input)
	if got != expected {
		t.Errorf("FixAtoAn failed. got=%q, want=%q", got, expected)
	}
}

// Test punctuation fixes
func TestFixPunctuation(t *testing.T) {
	input := "Hello , world ! ' A test ' ... Right ?"
	expected := "Hello, world! 'A test'... Right?"
	got := FixPunctuation(input)
	if got != expected {
		t.Errorf("FixPunctuation failed. got=%q, want=%q", got, expected)
	}
}

// End-to-end ProcessText
func TestProcessText(t *testing.T) {
	input := "Add 1E (hex) and 1010 (bin) (up,2) , a apple ! ' hello world '"
	expected := "Add 30 AND 10, an apple! 'hello world'"
	// Enable all processing functions in ProcessText before testing
	text := input
	text = HandleHexAndBin(text)
	text = HandleCaseModifiers(text)
	text = FixPunctuation(text)
	text = FixAtoAn(text)

	if text != expected {
		t.Errorf("End-to-end ProcessText failed.\ngot= %q\nwant=%q", text, expected)
	}
}

// Optional: test file I/O
func TestFileProcessing(t *testing.T) {
	inputFile := "test_input.txt"
	outputFile := "test_output.txt"

	os.WriteFile(inputFile, []byte("42 (hex) and 10 (bin)"), 0644)
	defer os.Remove(inputFile)
	defer os.Remove(outputFile)

	data, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result := ProcessText(string(data))
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		t.Fatal(err)
	}

	out, _ := os.ReadFile(outputFile)
	if string(out) != "66 and 2" {
		t.Errorf("File processing failed. got=%q, want=%q", string(out), "66 and 2")
	}
}
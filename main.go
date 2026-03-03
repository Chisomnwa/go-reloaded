package main

import (
	"fmt"
	"os"
	"go-reloaded/functions"
)

func main() {
	/*
		Inside her:
		- I'll validate arguments
		- Read input file
		- Call processing function
		- Write output file
	*/

	// make sure that the arguments passed into the command line isn't more than or less than 3
	if len(os.Args) != 3 {
		fmt.Println("Usage: . input.txt output.txt")
		return
	}

	// create variables for the sample input and output text files
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// read the file and handle error while reading the file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("FIle reading error", err)
		return
	}

	// turn the read data into a string so, you can process it
	result := functions.ProcessText(string(data)) // Still wrap the process function around it

	// write to the output text file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error Writing file")
	}
}

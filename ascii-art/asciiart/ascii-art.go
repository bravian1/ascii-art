package asciiart

import (
	"fmt"
	"os"
	"strings"
)

const (
	CHARACTER_HEIGHT = 8
	NUMOFASCII_XTERS = 95
)

func Run(arr []string) {
	processed, filename := processInput(arr)
	if !processed {
		fmt.Println("Usage: go run . <input_string> <filename>\n\t<filename> is optional")
		return
	}
	args := replaceSpecialcharacters(arr[1])
	args = processBackspace(args)

	if !isValidInput(args) {
		fmt.Println("[Error] Input contains unacceptable characters")
		return
	}

	input := strings.Split(args, "\\n")
	if ok, whichSpace, count := IsAllSpace(input); ok && whichSpace == "NEWLINE" {
		// print count-1 number of \n
		for i := 0; i < count-1; i++ {
			fmt.Println()
		}
		os.Exit(0)
	} else if ok, whichSpace, _ := IsAllSpace(input); ok && whichSpace == "SPACE" {
		os.Exit(0)
	} else {
		created, alphabet := CreateAlphabet(filename)
		if !created {
			fmt.Printf("Could not create the alphabet. Are you sure %s exists and is a valid ascii file?\n", filename)
			os.Exit(0)
		}

		// process input and print
		for _, word := range input {
			if word == "" {
				fmt.Println()
			} else {
				printAscii(word, alphabet)
			}
		}
	}
}

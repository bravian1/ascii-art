package asciiart

import (
	"fmt"
	"os"
	"strings"
)

//echo -e "lines\ncols" | tput -S for getting terminal width and height

const (
	CHARACTER_HEIGHT   = 8
	NUMOFASCII_XTERS   = 95
	NUMOFLINESINBANNER = 854
)

// Run processes an array of strings to generate ASCII art from a specified file.
func Run(arr []string) {
	if len(arr) < 2 {
		printUsageMessageAndExit()
		return
	}

	// var colorFlag, colorTarget, inputString string
	inputstring, flagtype, target:=checkflag(arr)
	
	processed, filename := processInput(arr)
	if !processed {
		printUsageMessageAndExit()
		return
	}

	args := replaceSpecialcharacters(inputstring)
	args = processBackspace(args)

	if valid, xter := isValidInput(args); !valid {
		fmt.Printf("[Error] Input contains unacceptable character %q\n", xter)
		return
	}

	if args == "\\n" {
		fmt.Println()
		return
	} else if args == "" {
		return
	} else {
		input := strings.Split(args, "\\n")
		created, alphabet := CreateAlphabet(filename)
		if !created {
			fmt.Printf("Could not create the alphabet. Are you sure %s exists and is a valid ascii file?\n", filename)
			os.Exit(0)
		}
		

		// Process input and print
		for _, word := range input {
			if word == "" {
				fmt.Println()
			} else {
				if flagtype =="right" || flagtype == "center" || flagtype == "left" || flagtype == "justify"{
						printAsciiJustify(word, alphabet, flagtype)
				} else if flagtype != "" {
					printAsciiWithColor(word, alphabet, colorCodes[flagtype], target)
				} else {
					printAscii(word, alphabet)
				}
			}
		}
	}
}

// printUsageMessageAndExit prints the usage message and exits
func printUsageMessageAndExit() {
	fmt.Println(`
		Usage: go run . [OPTION] [STRING]
		EX: go run . --color=<color> <letters to be colored> "something" Or go run . --align=<align> "something"
	`)
	os.Exit(1)
}

func checkflag(arr []string) (string, string, string) {
	var inputstring, flagtype, target string
	if strings.HasPrefix(arr[1], "--color=") {
		flagtype = strings.TrimPrefix(arr[1], "--color=")
		if len(arr) > 3 {
			target = arr[2]
			inputstring = arr[3]
		} else if len(arr) == 3 {
			inputstring = arr[2]
		} else {
			printUsageMessageAndExit()
		}
	} else if strings.HasPrefix(arr[1], "--align=") {
		flagtype = strings.TrimPrefix(arr[1], "--align=")
		if len(arr) == 3 {
			inputstring = arr[2]
		} else {
			printUsageMessageAndExit()
		}
	} else {
		inputstring = arr[1]
		return inputstring, "", ""
	}

	return inputstring, flagtype, target
}

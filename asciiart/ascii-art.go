package asciiart

import (
	"fmt"
	"os"
	"strings"
)



const (
	CHARACTER_HEIGHT   = 8
	NUMOFASCII_XTERS   = 95
	NUMOFLINESINBANNER = 854
)

var (
	Cyan       = "\033[36m"
	Reset      = "\033[0m"
	colorCodes = map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"orange": "\033[38;5;208m",
		"reset":   "\033[0m",
	}
)

// printAsciiWithColor prints ASCII art with a specific color.
func printAsciiWithColor(word string, alphabet map[rune][]string, color string, target string) {
	for i := 0; i < CHARACTER_HEIGHT; i++ {
		lineOutput := ""
		for _, l := range word {
			switch l {
			case '\n':
				fmt.Println()
			default:
				line := alphabet[rune(l)][i]
				if shouldColor(rune(l), target) {
					lineOutput += color + line + Reset
				} else {
					lineOutput += line
				}
			}
		}
		fmt.Println(lineOutput)
	}
}

// shouldColor determines if a letter should be colored
func shouldColor(letter rune, target string) bool {
	if target == "" {
		return true
	}
	return strings.ContainsRune(target, letter)
}


// Run processes an array of strings to generate ASCII art from a specified file.
func Run(arr []string) {
	if len(arr) < 2 {
			printUsageMessageAndExit()
			return
	}

	var colorFlag, colorTarget, inputString string
	if strings.HasPrefix(arr[1], "--color=") {
			colorFlag = strings.TrimPrefix(arr[1], "--color=")
			if len(arr) > 3 {
					colorTarget = arr[2]
					inputString = arr[3]
			} else if len(arr) == 3 {
					inputString = arr[2]
			} else {
					printUsageMessageAndExit()
			}
	} else {
			inputString = arr[1]
	}

	processed, filename := processInput(arr)
	if !processed {
			printUsageMessageAndExit()
			return
	}

	args := replaceSpecialcharacters(inputString)
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
							if colorFlag != "" {
									printAsciiWithColor(word, alphabet, colorCodes[colorFlag], colorTarget)
							} else {
									printAscii(word, alphabet)
							}
					}
			}
	}
}
// printUsageMessageAndExit prints the usage message and exits
func printUsageMessageAndExit() {
	fmt.Println(`Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <letters to be colored> "something"`)
	os.Exit(1)
}

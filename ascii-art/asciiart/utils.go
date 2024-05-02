package asciiart

import (
	"fmt"
	"strings"
)

// isValid checks if the input string contains only acceptable characters. ([32,126] & [8,10])
// The special characters newline (10), carriage return (13), and backspace (8).
// - bool: true if contains only acceptable, false otherwise.
func isValidInput(args string) bool {
	for _, ch := range args {
		if !((ch >= 32 && ch <= 126) || (ch >= 8 && ch <= 10)) {
			return false
		}
	}
	return true
}

/*
* Replace all special characters with characters that can be recognized and processed by golang
* Make them fit inside a rune
*/
func replaceSpecialcharacters(s string) string {
	replacer := strings.NewReplacer(
		"\\v", "\\n\\n\\n\\n",
		"\n", "\\n",
		"\\t", "    ",
		"\\b", "\b",
		"\\r", "\r",
		"\\a", "\a",
		"\\f", "\f",
	)
	return replacer.Replace(s)
}

// IsAllSpace checks if an array of strings consists entirely of empty strings
// It differentiates between empty spaces and newlines
// Returns:
//   - bool: entirely empty strings (true) or non-empty strings (false).
//   - string: type of whitespace detected: "SPACE" or "NEWLINE"
//   - int: if "NEWLINE" then implies how many they are
func IsAllSpace(ar []string) (bool, string, int) {
	count := 0
	for _, ch := range ar {
		if ch != "" {
			return false, "", count
		} else {
			count++
		}
	}

	if count == 1 {
		return true, "SPACE", count
	} else {
		return true, "NEWLINE", count
	}
}

// printAscii prints the ascii art representation of the given word using the given alphabet
func printAscii(word string, alphabet map[rune][]string) {
	for i := 0; i < CHARACTER_HEIGHT; i++ {
		lineOutput := ""
		for _, l := range word {
			switch l {
			case '\n':
				fmt.Println()
			default:
				lineOutput += alphabet[rune(l)][i]
			}
		}
		fmt.Println(lineOutput)
	}
}

/*
* Processes the input array and extracts whatever the program needs
* It is also responsible of dealing with optional parameters like the file names
* Exits if input does not meet the expected format
 */
func processInput(arr []string) (bool, string) {
	if len(arr) < 2 || len(arr) > 3 {
		return false, ""
	}

	filename := "standard.txt"
	if len(arr) == 3 {
		switch arr[2] {
		case "shadow", "shadow.txt":
			filename = "shadow.txt"
		case "thinkertoy", "thinkertoy.txt":
			filename = "thinkertoy.txt"
		}
	}
	return true, filename
}

/*
* Delete characters given the backspace
 */
func processBackspace(s string) string {
	temp := ""
	for index, ch := range s {
		if ch != '\b' {
			temp += string(ch)
		} else {
			temp = temp[:index-1]
		}
	}
	return temp
}

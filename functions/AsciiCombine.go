package ascii

import (
	"fmt"
	"strings"
)

// Handles newlines aspect in input string,combines all input slices' asciiArt
func AsciiCombine(input string, asciiMap map[rune][]string) ([]string, error) {
	if input == "" {
		return []string{}, nil
	}
	input = strings.Replace(input, "\r", "\\r", -1)
	input = strings.Replace(input, "\\r", "\n", -1)
	input = strings.Replace(input, "\n", "\\n", -1)

	var result []string

	words := strings.Split(input, "\\n")
	for i, char := range words {
		if words[i] == "" {
			continue
		} else {
			ascii, err := Art(char, asciiMap)
			if err != nil {
				return result, fmt.Errorf("non ascii Character")
			}
			result = append(result, ascii)
		}
	}
	return result, nil
}

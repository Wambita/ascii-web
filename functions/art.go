package ascii

import (
	"errors"
)

// combiness the line of Ascii art per slices of the input string and predefined mapping
func Art(text string, m map[rune][]string) (string, error) {
	str := ""
	// ranges through the 8 lines of the ascii art
	for i := 0; i < 8; i++ {
		// ranges through the characters in input text adding them line by line
		for _, char := range text {
			asciiArt, ok := m[char]
			if !ok {
				return "", errors.New("Non ascii Character")
			}
			str += asciiArt[i]
		}
		str += "\n"

	}
	return str, nil
}

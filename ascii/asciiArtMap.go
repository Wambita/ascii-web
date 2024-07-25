package ascii

import (
	"strings"
)

// asciiArt maps characters to their respective Ascii art 8 lines long.
func AsciiArtMap(s string) map[rune][]string {
	wordArt := strings.Split(string(s), "\n")
	char := ' '
	result := make(map[rune][]string)

	// buffer is used to for lines containing the character temporarily
	var buffer []string

	num := len(wordArt)

	for i := 0; i < num; i++ {

		line := wordArt[i]

		// if loop store the content found in the first line which is nil
		if i == 0 {
			buffer = make([]string, 0)
		} else if i%9 == 0 {
			result[char] = buffer
			buffer = make([]string, 0)
			char++
		} else {
			buffer = append(buffer, line)
		}
	}
	return result
}

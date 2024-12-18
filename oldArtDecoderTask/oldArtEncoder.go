package oldArtDecoderTask

import (
	"errors"
	"fmt"
	"strings"
)

// encodes ASCII art from decoded input
func Encode(decodedArt string) (string, error) {

	if strings.TrimSpace(decodedArt) == "" { // if empty string or spaces
		return "", errors.New(": Empty input")
	}
	var encoded strings.Builder // variable for tehe encoded string

	i := 0
	for i < len(decodedArt) {
		count, length, pattern := findLongestRepeatedPattern(decodedArt[i:])
		if count > 1 { // if more than one pattern
			count++ // keep counting "plus one" and write it as many time as count
			encoded.WriteString(fmt.Sprintf("[%d %s]", count, pattern))
			i += count * length
		} else { // if no pattern, leave untouched
			encoded.WriteByte(decodedArt[i])
			i++
		}
	}
	return encoded.String(), nil // returns encoded ASCII art
} // encode() END
///////////////////////////////////////////////////////////

// finds repetiteve patterns
func findLongestRepeatedPattern(decodedArt string) (int, int, string) {
	maxCount := 0         // max count of pattern
	maxLength := 0        //max length of pattern
	var maxPattern string // variable for the string

	for i := 0; i < len(decodedArt)/2; i++ { // iterates over decodedArt
		pattern := decodedArt[:i+1] // every round character plus one
		count := 0                  // how many same character/pattern is found
		length := len(pattern)      // length of current pattern
		// compare current pattern and substring
		for j := i + 1; j < len(decodedArt)-length+1; j += length {
			if decodedArt[j:j+length] == pattern {
				count++ // counts pattern that repeats
			} else {
				break // exit what different pattern is found
			}
		}
		if count > maxCount { // upadate if necessary
			maxCount = count
			maxLength = length
			maxPattern = pattern
		}
	} // Return the longest repeated pattern
	return maxCount, maxLength, maxPattern
} // findLongestRepeatedPattern() END
///////////////////////////////////////////////////////////
package oldArtDecoderTask

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// decodes ASCII art from encoded input
func Decode(encodedArt string) (string, error) {

	if strings.TrimSpace(encodedArt) == "" { // if empty string or spaces
		return "", errors.New(": Empty input")
	}
	// count [ and ]
	openingBrackets := strings.Count(encodedArt, "[")
	closingBrackets := strings.Count(encodedArt, "]")
	if openingBrackets != closingBrackets {
		return "", errors.New(": Brackets mismatch")
	}
	// Find errors ///////////////////////////////////////////////////////////
	re1 := regexp.MustCompile(`\[([^\[\]]*?)\]`) // anything between [ and ], but no [ ]
	matches1 := re1.FindAllStringSubmatch(encodedArt, -1)

	// Check each match for content
	for _, match1 := range matches1 {
		content := match1[1]
		// content inside [ ]
		if !strings.Contains(content, " ") { // if no space
			return "", errors.New(": Missing space")
		}
		// Check if no number
		if content != "" {
			// Convert content to integer
			_, err := strconv.Atoi(strings.Split(content, " ")[0])
			if err != nil {
				return "", errors.New(": Missing number")
			}
		}
	} // if found just [ number space ]
	re2 := regexp.MustCompile(`\[(\d+) \]`)
	if re2.MatchString(encodedArt) {
		return "", errors.New(": Missing symbol(s)")
	}
	// Find matches ///////////////////////////////////////////////////////////
	re := regexp.MustCompile(`\[(\d+) ([^[\]]+)\]`) // [ (number) space (any but [ ]) ]
	// Replace patterns
	decodedArt := replacePattern(encodedArt, re)
	return decodedArt, nil // returns decoded ASCII art and nil if no errors
} // decode() END"// decode(() END
///////////////////////////////////////////////////////////

// replacePattern replaces a matched pattern with the decoded art
func replacePattern(match string, re *regexp.Regexp) string {
	decodedArt := re.ReplaceAllStringFunc(match, func(s string) string {
		match := re.FindStringSubmatch(s)
		countStr := match[1]                 // number
		symbol := match[2]                   // characters
		count, _ := strconv.Atoi(countStr)   // number to int
		return strings.Repeat(symbol, count) // symbol as many times as count
	})
	return decodedArt // returns decoded ASCII art

} // replacePattern(() END
///////////////////////////////////////////////////////////
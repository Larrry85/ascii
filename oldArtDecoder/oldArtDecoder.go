package oldArtDecoder

import (
	//"bufio"
	"errors"
	"fmt"
	//"os"
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

/*
func main() {
	args := os.Args[1:]

	// command line arguments and usage
	if len(args) == 0 || len(args) == 1 && os.Args[1] == "-h" { // if go run . OR go run . -h
		fmt.Println("- Art tool -")
		fmt.Print("\n")
		fmt.Println("Usage: go run . <flags> [input]")
		fmt.Print("\n")
		fmt.Println("Flags:")
		fmt.Println("  -h : Prints usage menu")
		fmt.Println("  -decode [input string] : Decode one line art in shell")
		fmt.Println("  -decode-multiline ENTER [input lines]: Decode multi-line art in shell")
		fmt.Println("  -decode-file [filename] : Decode art from a file to shell")
		fmt.Println("  -encode [input string] : Encode one line art in shell")
		fmt.Println("  -encode-multiline ENTER [input lines]: Encode multi-line art in shell")
		fmt.Println("  -encode-file [filename] : Encode art from a file to shell")
		return
	}

	flag := args[0] // first argument = flag
	// six options
	switch flag {

	///////////////////////////////////////////////////////////
	// decode one line in shell
	case "-decode":
		if len(args) != 2 { // if string is missing
			fmt.Println("error: missing string")
			return
		}
		decoded, err := Decode(args[1]) // calls decode() with second argument = input
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Print("\n")
		fmt.Println("Decoded ASCII Art:")
		fmt.Print("\n")
		fmt.Println(decoded) // prints decoded input string
		fmt.Print("\n")

		// user can choose if they want to encode back or exit the program
		fmt.Println("Do you want to encode the art back to its encoded version (y) / end (any letter)?")
		var ask string
		fmt.Scanln(&ask) // user input
		if ask == "y" {
			encodedBack, err := Encode(decoded) // decoded back to encoded
			fmt.Print("\n")
			fmt.Println("Back to encoded form:")
			fmt.Print("\n") // result if no errors
			fmt.Println(encodedBack, "\n**********************************************************************")
			if err != nil {
				fmt.Println("Error encoding ASCII art:", err)
				return
			}
		} else { // Farewell to user!
			fmt.Println("Bye!\n**********************************************************************")
		}
	///////////////////////////////////////////////////////////
	// decode multiple lines in shell
	case "-decode-multiline":
		fmt.Println("Paste multi-line strings and then press Enter twice to finish:")
		fmt.Print("\n")
		var inputLines []string
		scanner := bufio.NewScanner(os.Stdin) // reads input
		for scanner.Scan() {                  // scanning input until empty line
			line := scanner.Text()
			if line == "" {
				break
			}
			inputLines = append(inputLines, line) // non empty lines into input slice
		}
		// if only Enters, no input, or empty string
		if len(inputLines) == 0 || (len(inputLines) == 1 && inputLines[0] == "") {
			fmt.Println("No input provided. Exiting...")
			return
		}

		var decodedArt strings.Builder // builds strings with Write. Works with String().

		for _, line := range inputLines { // iterate over input slice
			decoded, err := Decode(line) // calls decode() with line of input slice
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			decodedArt.WriteString(decoded) // string builder: prints decoded input
			decodedArt.WriteString("\n")    // add newline after line
		}

		fmt.Print("\n")
		fmt.Println("Decoded ASCII Art (MULTI):")
		fmt.Print("\n")
		fmt.Print(decodedArt.String()) // prints decoded input string
		fmt.Print("\n")

		fmt.Println("Do you want to encode the art back to its encoded version (y) / end (any letter)?")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			encodedBack, err := Encode(decodedArt.String())
			fmt.Print("\n")
			fmt.Println("Back to encoded form:")
			fmt.Print("\n")
			fmt.Println(encodedBack, "\n**********************************************************************")
			if err != nil {
				fmt.Println("Error encoding ASCII art:", err)
				return
			}
		} else {
			fmt.Println("Bye!\n**********************************************************************")
		}
	///////////////////////////////////////////////////////////
	// decode from file to shell
	case "-decode-file":
		if len(args) != 2 { // if text file is missing
			fmt.Println("error: missing file")
			return
		}
		filePath := args[1]            // second argument = input
		file, err := os.Open(filePath) // open text file
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		defer file.Close() // close text file!

		stat, err := file.Stat()
		if err != nil { // returns file info or prints error
			fmt.Println("error:", err)
			return
		} // if the sixe of the file is zero
		if stat.Size() == 0 {
			fmt.Println("The file is empty.")
			return
		}

		var inputLines []string
		scanner := bufio.NewScanner(file) // reads input from file
		for scanner.Scan() {              // scanning input until empty line
			line := scanner.Text()
			if line == "" {
				break
			}
			inputLines = append(inputLines, line) // non empty lines into input slice
		}
		var decodedArt strings.Builder // builds strings with Write. Works with String().

		for _, line := range inputLines { // iterate over input slice
			decoded, err := Decode(line) // calls decode() with line of input slice
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			decodedArt.WriteString(decoded) // string builder: prints decoded input
			decodedArt.WriteString("\n")    // add newline after line
		}

		fmt.Print("\n")
		fmt.Println("Decoded ASCII Art (FILE):")
		fmt.Print("\n")
		fmt.Print(decodedArt.String()) // prints decoded input string
		fmt.Print("\n")

		fmt.Println("Do you want to encode the art back to its encoded version (y) / end (any letter)?")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			encodedBack, err := Encode(decodedArt.String())

			fmt.Print("\n")
			fmt.Println("Back to encoded form:")
			fmt.Print("\n")
			fmt.Println(encodedBack, "\n**********************************************************************")

			if err != nil {
				fmt.Println("Error encoding ASCII art:", err)
				return
			}
		} else {
			fmt.Println("Bye!\n**********************************************************************")
		}
	///////////////////////////////////////////////////////////
	// encode art in shell
	case "-encode":
		if len(args) != 2 {
			fmt.Println("error: missing string")
			return
		}
		encoded, err := Encode(args[1])
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Print("\n")
		fmt.Println("Encoded ASCII Art:")
		fmt.Print("\n")
		fmt.Println(encoded)
		fmt.Print("\n")

		fmt.Println("Do you want to decode the art back to its decoded version (y) / end (any letter)?")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			decodedBack, err := Decode(encoded)
			fmt.Print("\n")
			fmt.Println("Back to decoded form:")
			fmt.Print("\n")
			fmt.Println(decodedBack, "\n**********************************************************************")
			if err != nil {
				fmt.Println("Error decoding ASCII art:", err)
				return
			}
		} else {
			fmt.Println("Bye!\n**********************************************************************")
		}
	///////////////////////////////////////////////////////////
	// encode multiple lines in shell
	case "-encode-multiline":
		fmt.Println("Paste multi-line strings and then press Enter twice to finish:")
		fmt.Print("\n")
		var inputLines []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			inputLines = append(inputLines, line)
		}
		// if only Enters, no input
		if len(inputLines) == 0 {
			fmt.Println("No input provided. Exiting...")
			return
		}

		var encodedArt strings.Builder

		for _, line := range inputLines {
			encoded, err := Encode(line)
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			encodedArt.WriteString(encoded)
			encodedArt.WriteString("\n")
		}

		fmt.Print("\n")
		fmt.Println("Encoded ASCII Art (MULTI):")
		fmt.Print("\n")
		fmt.Print(encodedArt.String())
		fmt.Print("\n")

		fmt.Println("Do you want to decode the art back to its decoded version (y) / end (any letter)?")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			decodedBack, err := Decode(encodedArt.String())
			fmt.Print("\n")
			fmt.Println("Back to decoded form:")
			fmt.Print("\n")
			fmt.Println(decodedBack, "\n**********************************************************************")
			if err != nil {
				fmt.Println("Error decoding ASCII art:", err)
				return
			}
		} else {
			fmt.Println("Bye!\n**********************************************************************")
		}
	///////////////////////////////////////////////////////////
	// encode art from file to shell
	case "-encode-file":
		if len(args) != 2 {
			fmt.Println("error: missing file")
			return
		}
		filePath := args[1]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		if stat.Size() == 0 {
			fmt.Println("The file is empty.")
			return
		}

		var inputLines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			inputLines = append(inputLines, line)
		}

		var encodedArt strings.Builder

		for _, line := range inputLines {
			encoded, err := Encode(line)
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			encodedArt.WriteString(encoded)
			encodedArt.WriteString("\n")
		}

		fmt.Print("\n")
		fmt.Println("Encoded ASCII Art (FILE):")
		fmt.Print("\n")
		fmt.Print(encodedArt.String())
		fmt.Print("\n")

		fmt.Println("Do you want to decode the art back to its decoded version (y) / end (any letter)?")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			decodedBack, err := Decode(encodedArt.String())
			fmt.Print("\n")
			fmt.Println("Back to decoded form:")
			fmt.Print("\n")
			fmt.Println(decodedBack, "\n**********************************************************************")
			if err != nil {
				fmt.Println("Error decoding ASCII art:", err)
				return
			}
		} else {
			fmt.Println("Bye!\n**********************************************************************")
		}
	///////////////////////////////////////////////////////////
	// if wrong flag
	default:
		fmt.Println("error: unknown flag")
	}
} // main() END
///////////////////////////////////////////////////////////*/
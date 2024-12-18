// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"art/oldArtDecoderTask" // decode and encode functions
	"art/server" // server.go
)

func main() {
	server.Server() // start a server

	///////////////////////////////////////////////////////////
	// oldArtDecoderTask here....
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
		decoded, err := oldArtDecoderTask.Decode(args[1]) // calls decode() with second argument = input
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
			encodedBack, err := oldArtDecoderTask.Encode(decoded) // decoded back to encoded
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
			decoded, err := oldArtDecoderTask.Decode(line) // calls decode() with line of input slice
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
			encodedBack, err := oldArtDecoderTask.Encode(decodedArt.String())
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
			decoded, err := oldArtDecoderTask.Decode(line) // calls decode() with line of input slice
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
			encodedBack, err := oldArtDecoderTask.Encode(decodedArt.String())

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
		encoded, err := oldArtDecoderTask.Encode(args[1])
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
			decodedBack, err := oldArtDecoderTask.Decode(encoded)
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
			encoded, err := oldArtDecoderTask.Encode(line)
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
			decodedBack, err := oldArtDecoderTask.Decode(encodedArt.String())
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
			encoded, err := oldArtDecoderTask.Encode(line)
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
			decodedBack, err := oldArtDecoderTask.Decode(encodedArt.String())
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
///////////////////////////////////////////////////////////
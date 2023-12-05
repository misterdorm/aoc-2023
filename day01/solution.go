// Write a Go program which reads lines of text from a file and identifies the first
// digit character and the last digit character in each line, then combining the two
// to make a 2-digit number.  There may be multiple digit characters on a line, or only
// a single digit character (in which case, that single character is both the first and
// the last digit.)  Print the sum of all the 2-digit numbers.
//
// Example input file contents:
// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
//
// Expected output (to stdout):
// 142
//
// Note: the file may be very large, so don't read the entire file into memory.
// You may assume that each line of text contains at least one digit character.
//
// Hint: see the Go standard library package "bufio" for a buffered reader type.
//

// We now need to refactor this to support digits in the strings that are spelled
// out in English.  For example, "one" would be a 1, "two" would be a 2, etc.  So
// we also need to look for those English words in addition to single digit characters.
// Note that the English words may overlap, so a string like "twone" would represent
// a first digit of 2 ("two") and a last digit of 1 ("one").

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// check if the command line arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide the input file name as a command line argument.")
		return
	}

	// get the input file name from the command line argument
	fileName := os.Args[1]

	// open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// close the file when we're done
	defer file.Close()

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)
	// keep track of the sum of all 2-digit numbers
	var sum int
	// scan each line
	for scanner.Scan() {
		// get the line
		line := scanner.Text()
		// get the first digit
		firstDigit := getFirstDigit(line)

		// get the last digit
		lastDigit := getLastDigit(line)

		// convert the digits to ints
		firstInt, _ := strconv.Atoi(firstDigit)
		lastInt, _ := strconv.Atoi(lastDigit)
		// combine them into a 2-digit number
		twoDigitNumber := (firstInt * 10) + lastInt
		// add it to the sum
		sum += twoDigitNumber
	}
	// print the sum
	fmt.Println(sum)
}

func getFirstDigit(line string) string {
	// check if the line contains a single digit character
	for i := 0; i < len(line); i++ {
		char := line[i]
		// Check if the character in this position is a digit
		if isDigit(char) {
			return string(char)
		}

		// Check if the substring starting at this position starts with
		// an English word representing a digit
		word := line[i:]
		// Define the regex pattern for English words representing digits
		pattern := "(?i)^(zero|one|two|three|four|five|six|seven|eight|nine)"
		// Compile the regex pattern
		regex := regexp.MustCompile(pattern)
		// Find the first match in the word
		match := regex.FindString(word)
		// Check if a match is found
		if match != "" {
			// Pass the match to getDigitFromWord to translate the word to a digit character
			digit := getDigitFromWord(match)
			if digit != "" {
				return digit
			}
		}
	}

	return ""
}

func getLastDigit(line string) string {
	// check if the line contains a single digit character at the end
	if len(line) == 1 && isDigit(line[0]) {
		return string(line[0])
	}

	// check if the line contains an English word representing a digit at the end
	for i := len(line) - 1; i >= 0; i-- {
		char := line[i]
		// Check if the character in this position is a digit
		if isDigit(char) {
			return string(char)
		}

		// Check if the substring ending at this position ends with
		// an English word representing a digit
		word := line[:i+1]
		// Define the regex pattern for English words representing digits
		pattern := "(?i)(zero|one|two|three|four|five|six|seven|eight|nine)$"
		// Compile the regex pattern
		regex := regexp.MustCompile(pattern)
		// Find the last match in the word
		match := regex.FindString(word)
		// Check if a match is found
		if match != "" {
			// Pass the match to getDigitFromWord to translate the word to a digit character
			digit := getDigitFromWord(match)
			if digit != "" {
				return digit
			}
		}
	}

	return ""
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func getDigitFromWord(word string) string {
	word = strings.ToLower(word) // Convert word to lowercase

	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}

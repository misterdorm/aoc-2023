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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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
		var firstDigit rune
		for _, c := range line {
			if unicode.IsDigit(rune(c)) {
				firstDigit = rune(c)
				break
			}
		}

		// get the last digit
		var lastDigit rune
		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if unicode.IsDigit(rune(c)) {
				lastDigit = rune(c)
				break
			}
		}
		// convert the digits to ints
		firstInt, _ := strconv.Atoi(string(firstDigit))
		lastInt, _ := strconv.Atoi(string(lastDigit))
		// combine them into a 2-digit number
		twoDigitNumber := (firstInt * 10) + lastInt
		// add it to the sum
		sum += twoDigitNumber
	}
	// print the sum
	fmt.Println(sum)
}

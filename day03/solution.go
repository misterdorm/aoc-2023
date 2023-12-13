// A program to read a text representation of an engine schematic diagram,
// which contains periods/dots (.), numbers (single or multiple digits), and
// other symbols.  Each number that is adjacent to a symbol (any character
// other than a period/dot, newline, or another number) represents a part
// number.  A part number is adjacent to a symbol if it is next to it on the
// same line, or if it is above or below it on the same column, or if it is
// adjacent to it diagonally.  The program will read the schematic diagram,
// identify which numbers represent part numbers, and print out a sum of all
// the part numbers.
//
// Example input:
// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var checkedLocations map[int]bool
var width int
var ratioSum int

func main() {
	// Open the file specified on the command line
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create an array of length 3 to store the most recent three lines
	var previousLines [3]string

	// Keep track of the special case at the beginning of the file
	var firstLineDone bool = false

	// Create a partNumbers map to keep track of all found part numbers
	var partNumbers []string

	// Initialize the checkedLocations map
	checkedLocations = make(map[int]bool)
	ratioSum = 0

	// Read one line from the file at a time
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// If the width variable is undefined or not set to a value, set it to the length of the first line (less the newline character)
		if width == 0 {
			width = len(line)
		}

		// Store the current line in the array
		previousLines[2] = line

		// Process the previous two lines
		if previousLines[0] != "" && previousLines[1] != "" {
			if !firstLineDone {
				// Special case for the first line of the file
				// Shift previousLines array to the right, with the first element being an empty string
				var firstLines [3]string
				firstLines[0] = ""
				firstLines[1] = previousLines[0]
				firstLines[2] = previousLines[1]
				nums := FindPartNumbers(firstLines, lineCount-2)

				// Merge the array returned by FindNumbersAt() into the partNumbers array
				partNumbers = append(partNumbers, nums...)

				firstLineDone = true
			}

			nums := FindPartNumbers(previousLines, lineCount-1)
			// Merge the array returned by FindNumbersAt() into the partNumbers array
			partNumbers = append(partNumbers, nums...)
		}

		// Shift the lines in the array
		previousLines[0] = previousLines[1]
		previousLines[1] = previousLines[2]
		lineCount++
	}

	// Special case for the last line of the file
	// Shift previousLines array to the left already done above, set last element being an empty string
	previousLines[2] = ""
	nums := FindPartNumbers(previousLines, lineCount)
	// Merge the array returned by FindNumbersAt() into the partNumbers array
	partNumbers = append(partNumbers, nums...)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Convert each key from partNumbers to an integer and add it to the sum
	var sum int = 0
	for _, numStr := range partNumbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}

	if err != nil {
		log.Fatal(err)
	}
	println("Sum of part numbers:", sum)
	println("Sum of gear ratios:", ratioSum)
}

// Find part numbers on the current line, based on context from the previous and
// next lines.  Find any symbols (characters other that periods/dots
// or numbers) on any of three lines.  If a symbol is found, check for numbers that
// are immediately adjacent to it on the same line, or above or below it on the
// previous or next line, or diagonally adjacent to it on the previous or next lines.
// Part numbers may be single or multiple digits.  It is considered to be adjacent
// to a symbol if any digit character of the part number is adjacent to the symbol.
// Return an array of part numbers found on the current line.
func FindPartNumbers(lines [3]string, lineNum int) []string {
	var allPartNumbers []string
	var oneSymPartNumbers []string

	println("-------------")
	println("Line 0:", lines[0])
	println("Line 1:", lines[1])
	println("Line 2:", lines[2])

	// Look for any symbols on any line
	for c, char := range lines[1] {
		if !isNumber(byte(char)) && char != '.' {
			println("Symbol at ", c, ":", string(char))

			// reset oneSymPartNumbers array to empty
			oneSymPartNumbers = nil

			// Look for adjacent (horizontally) numbers on the "current" (middle) line
			// On either side of the symbol in the line string, look for numbers
			// that are adjacent to it in the string. The number could be one or
			// several digits long.  If a number is found, add it to the array of
			// part numbers.
			if c > 0 {
				nums := FindNumbersAt(lines[1], c-1, lineNum)
				// Merge the array returned by FindNumbersAt() into the oneSymPartNumbers array
				oneSymPartNumbers = append(oneSymPartNumbers, nums...)
			}
			if c < len(lines[1])-1 {
				nums := FindNumbersAt(lines[1], c+1, lineNum)
				// Merge the array returned by FindNumbersAt() into the oneSymPartNumbers array
				oneSymPartNumbers = append(oneSymPartNumbers, nums...)
			}

			// Look for adjacent (vertically and diagonally) numbers on the previous and next lines
			for c1 := c - 1; c1 <= c+1; c1++ {
				if c1 >= 0 && c1 < len(lines[1]) {
					nums := FindNumbersAt(lines[0], c1, lineNum-1)
					// Merge the array returned by FindNumbersAt() into the oneSymPartNumbers array
					oneSymPartNumbers = append(oneSymPartNumbers, nums...)

					nums = FindNumbersAt(lines[2], c1, lineNum+1)
					// Merge the array returned by FindNumbersAt() into the oneSymPartNumbers array
					oneSymPartNumbers = append(oneSymPartNumbers, nums...)
				}
			}

			println("")
		}

		// If the symbol character `char` is an asterisk (*), and the number of elements in the
		// oneSymPartNumbers array is exactly 2, then multiply the two numbers together and
		// add that to a sum variable
		if char == '*' && len(oneSymPartNumbers) == 2 {
			println("Found gear symbol!")
			num1, err := strconv.Atoi(oneSymPartNumbers[0])
			if err != nil {
				log.Fatal(err)
			}
			num2, err := strconv.Atoi(oneSymPartNumbers[1])
			if err != nil {
				log.Fatal(err)
			}
			product := num1 * num2
			ratioSum += product
			println("Found gear symbol! ", num1, " * ", num2, " = ", product, " new sum: ", ratioSum)
		}

		allPartNumbers = append(allPartNumbers, oneSymPartNumbers...)
	}

	return oneSymPartNumbers
}

// isNumber checks if a character is a digit.
func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func FindNumbersAt(line string, index int, lineNum int) []string {
	var numbers []string

	c1 := index

	// Check if this location has already been checked
	if checkedLocations[lineNum*width+c1] {
		return numbers
	}

	if c1 >= 0 && c1 < len(line) && isNumber(line[c1]) {
		checkedLocations[lineNum*width+c1] = true
		partNumber := ""

		// Check for digit characters forward and backward
		for c1 := index; c1 >= 0 && isNumber(line[c1]); c1-- {
			checkedLocations[lineNum*width+c1] = true
			partNumber = string(line[c1]) + partNumber
		}
		for c1 := index + 1; c1 < len(line) && isNumber(line[c1]); c1++ {
			checkedLocations[lineNum*width+c1] = true
			partNumber += string(line[c1])
		}
		println("+++ Part number:", partNumber)
		numbers = append(numbers, partNumber)
	}

	return numbers
}

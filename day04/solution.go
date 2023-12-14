// Read in an input file that's specified on the command line.
// Each line of the file contains a card number, "Card 1", "Card 2", etc.
// followed by a colon, followed a space-separated list of numbers, followed
// by a pipe charater ("|"), followed by another space-separated list of numbers.
// Determine how many of the numbers from the first list is contained in the
// second list.  Print out the card number and the count of numbers that were
// found.

// Sample input file:
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11

package main

// Define a main function that opens the file, reads each line
// and parses the card number, and the two lists of numbers.
// The numbers need not be converted to ints, they can be strings.
// Each list of number should be stored in a slice of strings and
// the card number should be stored in a string variable.

// The main function should then call a function that takes the
// card number and the two lists of numbers as arguments and returns
// the count of numbers found in the second list.
// The main function should print out the card number and the count
// of numbers found.

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: solution <input file>")
		os.Exit(1)
	}

	// Open the file
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	sumPoints := 0

	// Read the file one line at a time
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Parse the line into the card number, and the two lists of numbers
		card, list1, list2 := parseLine(line)

		// Call a function that takes the card number and the two lists of numbers
		// as arguments and returns the count of numbers found in the second list.
		count := countMatches(list1, list2)

		// Print out the card number and the count of numbers found.
		points := int(math.Pow(2, float64(count-1)))
		sumPoints += points
		fmt.Printf("%s: %d = %d points (total %d)\n", card, count, points, sumPoints)
	}

	f.Close()
}

// Parse a line of input into the card number, and the two lists of numbers.
// Return the card number and the two lists of numbers.
func parseLine(line string) (string, []string, []string) {
	// Split the line into three parts, using the colon and the pipe characters
	parts := strings.Split(line, ":")
	card := parts[0]
	parts = strings.Split(parts[1], "|")
	list1 := strings.Fields(parts[0])
	list2 := strings.Fields(parts[1])
	return card, list1, list2
}

// Take the card number and the two lists of numbers as arguments and return
// the count of numbers found in the second list.
func countMatches(list1 []string, list2 []string) int {
	count := 0
	for _, num1 := range list1 {
		for _, num2 := range list2 {
			if num1 == num2 {
				count++
			}
		}
	}
	return count
}

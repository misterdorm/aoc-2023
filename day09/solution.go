// The input file (specified on the command line) contains a list of integers
// on each line.  Read those into an array of slices of ints.
// Then, for each slice, calculate the sum of each pair of integers in the
// slice.  If the sum of any pair is equal to 2020, then print the product of
// those two integers.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check if the input file name is provided as a command line argument.
	if len(os.Args) < 2 {
		fmt.Println("Input file name is missing.")
		os.Exit(1)
	}

	// Get the input file name from the command line arguments.
	inputFile := os.Args[1]

	// Read the input file into a slice of slices of ints.
	input, err := readInput(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	// Loop over all the slices of ints in the input slice.
	next := 0
	for _, slice := range input {
		fmt.Printf("slice: %v\n", slice)
		// Calculate the difference between each integer in the slice.
		// recursively until all the differences are zero.

		next += extrapolateNextValue(slice)
		fmt.Printf("\n")
	}

	fmt.Printf("sum of nexts: %v\n", next)

}

func extrapolateNextValue(slice []int) int {

	var next int

	if sliceDifference(slice) == nil {
		// return last integer in the slice
		next = slice[len(slice)-1]
	} else {
		// calculate the next integer in the slice
		next = slice[len(slice)-1] + extrapolateNextValue(sliceDifference(slice))
	}

	fmt.Printf("next: %v\n", next)
	return next

}

// sliceDifference function that takes a slide of ints and returns another
// slice that contains the difference between each integer in the input slice
// therefore the returned slice will have one less element than the input slice.
// returns nil if all the differences are zero
func sliceDifference(slice []int) []int {
	var diff []int
	zeroes := true

	for i := 1; i < len(slice); i++ {
		diff = append(diff, slice[i]-slice[i-1])
		if slice[i]-slice[i-1] != 0 {
			zeroes = false
		}
	}

	fmt.Printf("diff: %v\n", diff)

	if zeroes {
		return nil
	} else {
		return diff
	}
}

// readInput reads the input file into a slice of slices of ints.  Each line
// of the file is a slice of ints.
func readInput(filename string) ([][]int, error) {
	// Open the input file.
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read each line of the file into a slice of slices of ints.
	var input [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ") // Split the line string on spaces
		var slice []int
		for _, str := range words {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			slice = append(slice, num)
		}
		input = append(input, slice)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

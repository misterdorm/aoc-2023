// As you try to figure out what to do next, you notice a poster on a wall near the ferry dock. "Boat races! Open to the public! Grand prize is an all-expenses-paid trip to Desert Island!" That must be where the sand comes from! Best of all, the boat races are starting in just a few minutes.
//
// You manage to sign up as a competitor in the boat races just in time. The organizer explains that it's not really a traditional race - instead, you will get a fixed amount of time during which your boat has to travel as far as it can, and you win if your boat goes the farthest.
//
// As part of signing up, you get a sheet of paper (your puzzle input) that lists the time allowed for each race and also the best distance ever recorded in that race. To guarantee you win the grand prize, you need to make sure you go farther in each race than the current record holder.
//
// The organizer brings you over to the area where the boat races are held. The boats are much smaller than you expected - they're actually toy boats, each with a big button on top. Holding down the button charges the boat, and releasing the button allows the boat to move. Boats move faster if their button was held longer, but time spent holding the button counts against the total race time. You can only hold the button at the start of the race, and boats don't move until the button is released.
//
// For example:
//
// Time:      7  15   30
// Distance:  9  40  200
// This document describes three races:
//
// The first race lasts 7 milliseconds. The record distance in this race is 9 millimeters.
// The second race lasts 15 milliseconds. The record distance in this race is 40 millimeters.
// The third race lasts 30 milliseconds. The record distance in this race is 200 millimeters.
// Your toy boat has a starting speed of zero millimeters per millisecond. For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by one millimeter per millisecond.
//
// So, because the first race lasts 7 milliseconds, you only have a few options:
//
// Don't hold the button at all (that is, hold it for 0 milliseconds) at the start of the race. The boat won't move; it will have traveled 0 millimeters by the end of the race.
// Hold the button for 1 millisecond at the start of the race. Then, the boat will travel at a speed of 1 millimeter per millisecond for 6 milliseconds, reaching a total distance traveled of 6 millimeters.
// Hold the button for 2 milliseconds, giving the boat a speed of 2 millimeters per millisecond. It will then get 5 milliseconds to move, reaching a total distance of 10 millimeters.
// Hold the button for 3 milliseconds. After its remaining 4 milliseconds of travel time, the boat will have gone 12 millimeters.
// Hold the button for 4 milliseconds. After its remaining 3 milliseconds of travel time, the boat will have gone 12 millimeters.
// Hold the button for 5 milliseconds, causing the boat to travel a total of 10 millimeters.
// Hold the button for 6 milliseconds, causing the boat to travel a total of 6 millimeters.
// Hold the button for 7 milliseconds. That's the entire duration of the race. You never let go of the button. The boat can't move until you let go of the button. Please make sure you let go of the button so the boat gets to move. 0 millimeters.
// Since the current record for this race is 9 millimeters, there are actually 4 different ways you could win: you could hold the button for 2, 3, 4, or 5 milliseconds at the start of the race.
//
// In the second race, you could hold the button for at least 4 milliseconds and at most 11 milliseconds and beat the record, a total of 8 different ways to win.
//
// In the third race, you could hold the button for at least 11 milliseconds and no more than 19 milliseconds and still beat the record, a total of 9 ways you could win.
//
// To see how much margin of error you have, determine the number of ways you can beat the record in each race; in this example, if you multiply these values together, you get 288 (4 * 8 * 9).
//
// Determine the number of ways you could beat the record in each race. What do you get if you multiply these numbers together?

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// main function which calls readInputFile() to read the input file and then
// calls calculateNumberOfWaysToWin() to calculate the number of ways to win
func main() {
	races := readInputFile()

	// For each race returned from readInputFile(), calculate the number of ways
	// to win and multiply them together
	totalNumberOfWaysToWin := 1

	for _, race := range races {
		totalNumberOfWaysToWin *= calculateNumberOfWaysToWin(race[0], race[1])
	}

	fmt.Printf("+++ totalNumberOfWaysToWin: %d\n", totalNumberOfWaysToWin)

}

// Function that calculates the minimum and maximum number of milliseconds to
// hold the button such that the boat will travel farther than the record of
// that race.  Inputs are the race dureation and the record distance.  Outputs
// are the minimum and maximum number of milliseconds to hold the button.
func calculateNumberOfWaysToWin(raceDuration, recordDistance int) int {
	fmt.Printf("+++ raceDuration: %d, recordDistance: %d\n", raceDuration, recordDistance)
	// Start at 1ms (0ms is guraranteed to not beat the record), and increment
	// until we hit the lower limit of the "hold time" such that we'll beat the
	// record
	minHoldTime := 1
	for {
		if calculateDistanceTraveled(minHoldTime, raceDuration) > recordDistance {
			break
		}
		minHoldTime++
	}

	// Do the same thing, starting at raceDuration - 1 (holding for the entire
	// race duration is guraranteed to not beat the record), and decrement until
	// we hit the upper limit of the "hold time" such that we'll beat the record
	maxHoldTime := raceDuration - 1
	for {
		if calculateDistanceTraveled(maxHoldTime, raceDuration) > recordDistance {
			break
		}
		maxHoldTime--
	}

	// Return the number of ways to win (the difference between the max and min hold times)
	fmt.Printf("+++ minHoldTime: %d, maxHoldTime: %d, ways to win: %d\n", minHoldTime, maxHoldTime, maxHoldTime-minHoldTime+1)
	return maxHoldTime - minHoldTime + 1

}

// Function that calculates the distance traveled by the boat given the hold time and race duration
func calculateDistanceTraveled(holdTime, raceDuration int) int {
	speed := holdTime                     // Boat's speed increases by 1 millimeter per millisecond
	travelTime := raceDuration - holdTime // Remaining time for the boat to travel
	distance := speed * travelTime        // Distance traveled = speed * time
	return distance
}

// Function that reads the input file, given on the command line, and returns
// an array of arrays of ints, where the first element of each array is the
// time duration of the race (first line in the input file) and the second
// element is the record distance for that race (second line in the input file)
//
// The input file looks like this:
// Time:      7  15   30
// Distance:  9  40  200
//
// Where the time duration of the first race is 7, and the record distance for
// the first race is 9.
func readInputFile() [][]int {
	// Read the input file given on the command line
	inputFileName := os.Args[1]
	inputFile, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Println(err)
	}

	// Split the input file into lines
	inputLines := strings.Split(string(inputFile), "\n")

	// Split each line into words
	inputWords := [][]string{}
	for _, line := range inputLines {
		inputWords = append(inputWords, strings.Fields(line))
	}

	// Convert the input words into ints
	inputInts := [][]int{}
	for _, words := range inputWords {
		inputInts = append(inputInts, []int{})
		for _, word := range words {
			if i, err := strconv.Atoi(word); err == nil {
				inputInts[len(inputInts)-1] = append(inputInts[len(inputInts)-1], i)
			}
		}
	}

	// Take the frist two elements of the inputInts array and put the first element
	// of each array into a separate array containing two elements
	races := [][]int{}

	for i := 0; i < len(inputInts[0]); i++ {
		races = append(races, []int{inputInts[0][i], inputInts[1][i]})
	}

	return races
}

// Given and input file, specified on the command line, containing a list of
// game results.  The game consists of drawing a random number of colored
// cubes (red, green, or blue) from a bag several times.
// Each line in the file contains a game ID "Game #:" followe by a semi-colon
// separated list of drawing results.  Each result is a comma-separated list
// of how many of each colors were drawn.  For example:
//
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
//
// Note that the results of each draw may contain any number of colors, and
// they are not necessarily in any particular order.
//
// This code needs to determine which games are possible, given that we know
// the bag contains only 12 red cubes, 13 green cubes, and 14 blue cubes.  The
// final output of the program should be the sum of the game IDs of all the
// possible games

// Start with defining a struct to hold the game results and writing a function to
// parse the results of a single game (a single line from the input file), including
// the game ID.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameResult struct {
	GameID  int
	Results []Results
}

type Results struct {
	Red   int
	Green int
	Blue  int
}

func parseGameResult(line string) (GameResult, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return GameResult{}, fmt.Errorf("invalid game result format: %s", line)
	}

	gameID, err := strconv.Atoi(strings.TrimSpace(parts[0][5:]))
	if err != nil {
		return GameResult{}, fmt.Errorf("invalid game ID: %s", parts[0][5:])
	}

	results := strings.Split(strings.TrimSpace(parts[1]), ";")
	for i := range results {
		results[i] = strings.TrimSpace(results[i])
	}

	gameResults := GameResult{
		GameID:  gameID,
		Results: make([]Results, len(results)),
	}

	for i, result := range results {
		red, green, blue, err := parseDrawing(result)
		if err != nil {
			return GameResult{}, err
		}

		gameResults.Results[i] = Results{
			Red:   red,
			Green: green,
			Blue:  blue,
		}
	}

	return gameResults, nil
}

func parseDrawing(drawing string) (int, int, int, error) {
	colors := strings.Split(drawing, ",")
	red, green, blue := 0, 0, 0

	for _, color := range colors {
		color = strings.TrimSpace(color)
		parts := strings.Split(color, " ")
		if len(parts) != 2 {
			return 0, 0, 0, fmt.Errorf("invalid color format: %s", color)
		}

		count, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid count: %s", parts[0])
		}

		switch parts[1] {
		case "red":
			red += count
		case "green":
			green += count
		case "blue":
			blue += count
		default:
			return 0, 0, 0, fmt.Errorf("invalid color: %s", parts[1])
		}
	}

	return red, green, blue, nil
}

// Write the main function which opens the input file (as specified on the command line),
// reads each line, and calls parseGameResult to parse the results of each game.
// Print an error if there is a problem parsing the line.
// For each game, call the function to determine if the game is possible or not.
// Print the game ID and whether the game is possible or not.
// Finally, print the sum of the game IDs of all possible games.

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameSum := 0 // Initialize gameSum variable

	for scanner.Scan() {
		gameResult, err := parseGameResult(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		possible := isPossible(gameResult)
		fmt.Printf("Game %d: %t\n", gameResult.GameID, possible)

		if possible { // Increment gameSum if game is possible
			gameSum += gameResult.GameID
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sum of possible game IDs:", gameSum) // Print the sum of possible game IDs
}

// Write a function to determine if a game is possible or not.  The function
// should take a GameResult as input and return a bool indicating if the game
// is possible or not.
func isPossible(gameResult GameResult) bool {
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, result := range gameResult.Results {
		if result.Red > bag["red"] ||
			result.Green > bag["green"] ||
			result.Blue > bag["blue"] {
			return false
		}
	}

	return true
}

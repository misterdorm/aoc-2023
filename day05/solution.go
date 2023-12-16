// The almanac (your puzzle input) lists all of the seeds that need to be planted. It also lists what type of soil to use with each kind of seed, what type of fertilizer to use with each kind of soil, what type of water to use with each kind of fertilizer, and so on. Every type of seed, soil, fertilizer and so on is identified with a number, but numbers are reused by each category - that is, soil 123 and fertilizer 123 aren't necessarily related to each other.

// For example:

// seeds: 79 14 55 13

// seed-to-soil map:
// 50 98 2
// 52 50 48

// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15

// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4

// water-to-light map:
// 88 18 7
// 18 25 70

// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13

// temperature-to-humidity map:
// 0 69 1
// 1 0 69

// humidity-to-location map:
// 60 56 37
// 56 93 4
// The almanac starts by listing which seeds need to be planted: seeds 79, 14, 55, and 13.

// The rest of the almanac contains a list of maps which describe how to convert numbers from a source category into numbers in a destination category. That is, the section that starts with seed-to-soil map: describes how to convert a seed number (the source) to a soil number (the destination). This lets the gardener and his team know which soil to use with which seeds, which water to use with which fertilizer, and so on.

// Rather than list every source number and its corresponding destination number one by one, the maps describe entire ranges of numbers that can be converted. Each line within a map contains three numbers: the destination range start, the source range start, and the range length.

// Consider again the example seed-to-soil map:

// 50 98 2
// 52 50 48
// The first line has a destination range start of 50, a source range start of 98, and a range length of 2. This line means that the source range starts at 98 and contains two values: 98 and 99. The destination range is the same length, but it starts at 50, so its two values are 50 and 51. With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 corresponds to soil number 51.

// The second line means that the source range starts at 50 and contains 48 values: 50, 51, ..., 96, 97. This corresponds to a destination range starting at 52 and also containing 48 values: 52, 53, ..., 98, 99. So, seed number 53 corresponds to soil number 55.

// Any source numbers that aren't mapped correspond to the same destination number. So, seed number 10 corresponds to soil number 10.

// So, the entire list of seed numbers and their corresponding soil numbers looks like this:

// seed  soil
// 0     0
// 1     1
// ...   ...
// 48    48
// 49    49
// 50    52
// 51    53
// ...   ...
// 96    98
// 97    99
// 98    50
// 99    51
// With this map, you can look up the soil number required for each initial seed number:

// Seed number 79 corresponds to soil number 81.
// Seed number 14 corresponds to soil number 14.
// Seed number 55 corresponds to soil number 57.
// Seed number 13 corresponds to soil number 13.
// The gardener and his team want to get started as soon as possible, so they'd like to know the closest location that needs a seed. Using these maps, find the lowest location number that corresponds to any of the initial seeds. To do this, you'll need to convert each seed number through other categories until you can find its corresponding location number. In this example, the corresponding types are:

// Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
// Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
// Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
// Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.
// So, the lowest location number in this example is 35.

// What is the lowest location number that corresponds to any of the initial seed numbers?

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Conversion map struct containing a destination range start, a source range start, and a range length.
type ConversionMap struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

// Create a struct containing an array of ConversionMap structs
// for each type of conversion (seed-to-soil, soil-to-fertilizer, etc.)
type ConversionMaps struct {
	SeedToSoilMap         []ConversionMap
	SoilToFertilizerMap   []ConversionMap
	FertilizerToWaterMap  []ConversionMap
	WaterToLightMap       []ConversionMap
	LightToTemperatureMap []ConversionMap
	TemperatureToHumidity []ConversionMap
	HumidityToLocationMap []ConversionMap
}

// Main function that opens the input file (specified on the command line),
// calls another function for reading the first line of seed numbers, then
// calls another function for reading the conversion maps, until the end of the file

func main() {
	// Read the input file, and check for errors
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Read the first line of seed numbers
	seedNumbers := readSeedNumbers(scanner)

	fmt.Printf("%v\n", seedNumbers)

	// Read the conversion maps
	conversionMaps := readConversionMaps(scanner)

	// Print out the contents of each element of the conversionMaps struct
	fmt.Printf("%v\n", conversionMaps.SeedToSoilMap)
	fmt.Printf("%v\n", conversionMaps.SoilToFertilizerMap)
	fmt.Printf("%v\n", conversionMaps.FertilizerToWaterMap)
	fmt.Printf("%v\n", conversionMaps.WaterToLightMap)
	fmt.Printf("%v\n", conversionMaps.LightToTemperatureMap)
	fmt.Printf("%v\n", conversionMaps.TemperatureToHumidity)
	fmt.Printf("%v\n", conversionMaps.HumidityToLocationMap)

	// Calculate the lowest location number
	lowestLocationNumber := calculateLowestLocationNumber(seedNumbers, conversionMaps)

	// Print the lowest location number
	fmt.Println(lowestLocationNumber)

}

// Function that takes a number (source) and an array of ConversionMap structs,
// checking convertNumber for each conversion map, returning the resulting
// destination number if the source number is found in the conversion map.
// If the source number is not found in any conversion map, the function
// returns the source number.

func convertNumberArray(source int, conversionMaps []ConversionMap) int {
	// Loop through the conversion maps
	for _, conversionMap := range conversionMaps {
		// Check the conversion map for the source number
		destination := convertNumber(source, conversionMap.SourceRangeStart, conversionMap.DestinationRangeStart, conversionMap.RangeLength)

		// If the source number is found in the conversion map, return the destination number
		if destination != source {
			return destination
		}
	}

	// If the source number is not found in any conversion map, return the source number
	return source
}

// Function that takes an array of seed numbers and a ConversionMaps struct,
// and returns the lowest location number that corresponds to any of the seed numbers.

func calculateLowestLocationNumber(seedNumbers []int, conversionMaps *ConversionMaps) int {
	// Create an empty array of integers
	var locationNumbers []int

	// Loop through the seed numbers
	for _, seedNumber := range seedNumbers {
		// Convert the seed number to a location number
		locationNumber := convertNumberArray(seedNumber, conversionMaps.SeedToSoilMap)
		locationNumber = convertNumberArray(locationNumber, conversionMaps.SoilToFertilizerMap)
		locationNumber = convertNumberArray(locationNumber, conversionMaps.FertilizerToWaterMap)
		locationNumber = convertNumberArray(locationNumber, conversionMaps.WaterToLightMap)
		locationNumber = convertNumberArray(locationNumber, conversionMaps.LightToTemperatureMap)
		locationNumber = convertNumberArray(locationNumber, conversionMaps.TemperatureToHumidity)
		locationNumber = convertNumberArray(locationNumber, conversionMaps.HumidityToLocationMap)

		// Add the location number to the array of location numbers
		locationNumbers = append(locationNumbers, locationNumber)
	}

	// Find the lowest location number
	lowestLocationNumber := locationNumbers[0]
	for _, locationNumber := range locationNumbers {
		if locationNumber < lowestLocationNumber {
			lowestLocationNumber = locationNumber
		}
	}

	// Return the lowest location number
	return lowestLocationNumber
}

// Function that takes a number (source), source range start number,
// destination range start number, and range length, and returns the
// corresponding number (destination) in the destination range.  If
// the source number is not in the source range, the function returns
// the source number.

func convertNumber(source int, sourceRangeStart int, destinationRangeStart int, rangeLength int) int {
	// If the source number is not in the source range, return the source number
	if source < sourceRangeStart || source >= sourceRangeStart+rangeLength {
		return source
	}

	// Calculate the destination number
	destination := destinationRangeStart + (source - sourceRangeStart)

	// Return the destination number
	return destination
}

// Function for reading the first line of seed numbers,
// and returning an array of integers
// The line format looks like this:
// seeds: 79 14 55 13

func readSeedNumbers(scanner *bufio.Scanner) []int {
	// Create an empty array of integers
	var seedNumbers []int

	// Read the first line of the file
	scanner.Scan()
	line := scanner.Text()

	// Split the line into an array of strings
	words := strings.Split(line, " ")

	// Loop through the words
	for _, word := range words {
		// If the word is a number, convert it to an integer and add it to the array
		if number, err := strconv.Atoi(word); err == nil {
			seedNumbers = append(seedNumbers, number)
		}
	}

	// Return the array of integers
	return seedNumbers
}

// Function for reading the conversion maps,
// and returning a ConversionMaps struct
// The format of the maps looks like this:
// seed-to-soil map:
// 50 98 2
// 52 50 48

func readConversionMaps(scanner *bufio.Scanner) *ConversionMaps {
	// Create an array of conversion map
	var conversionMap ConversionMap

	// Create an empty conversionMaps struct
	var conversionMaps ConversionMaps

	var mapName string

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line matches a regular expression of the form "word-to-word map:"
		if words := regexp.MustCompile(`^(\w+-to-\w+) map:$`).FindStringSubmatch(line); words != nil {
			mapName = words[1]
		}

		// If the line starts with a number, split the line into an array of strings
		// and convert the strings to integers, then add them to the conversion map
		if line != "" && isDigit(line[0]) {
			words := strings.Split(line, " ")
			conversionMap.DestinationRangeStart, _ = strconv.Atoi(words[0])
			conversionMap.SourceRangeStart, _ = strconv.Atoi(words[1])
			conversionMap.RangeLength, _ = strconv.Atoi(words[2])

			// Determine which element of the conversionMaps struct to add the conversion map to,
			// based on the mapName string
			switch mapName {
			case "seed-to-soil":
				conversionMaps.SeedToSoilMap = append(conversionMaps.SeedToSoilMap, conversionMap)
			case "soil-to-fertilizer":
				conversionMaps.SoilToFertilizerMap = append(conversionMaps.SoilToFertilizerMap, conversionMap)
			case "fertilizer-to-water":
				conversionMaps.FertilizerToWaterMap = append(conversionMaps.FertilizerToWaterMap, conversionMap)
			case "water-to-light":
				conversionMaps.WaterToLightMap = append(conversionMaps.WaterToLightMap, conversionMap)
			case "light-to-temperature":
				conversionMaps.LightToTemperatureMap = append(conversionMaps.LightToTemperatureMap, conversionMap)
			case "temperature-to-humidity":
				conversionMaps.TemperatureToHumidity = append(conversionMaps.TemperatureToHumidity, conversionMap)
			case "humidity-to-location":
				conversionMaps.HumidityToLocationMap = append(conversionMaps.HumidityToLocationMap, conversionMap)
			}
		}
	}

	// Return the array of conversion maps
	return &conversionMaps
}

// As isDigit function for checking if a character is a digit
func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

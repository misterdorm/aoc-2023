// Implement a function for reading in a text file (given on the command line),
// that contains lines with a string followed by a number (separated by a space).
// Example:
//
// 32T3K 765
// T55J5 684
// KK677 28
// KTJJT 220
// QQQJA 483
//
// The function should return an array of Hand structs, where each struct
// contains the string and the number from the file. The array does not need to be sorted
//
// The function should return an error if the file cannot be read, or if the
// file contains invalid data.
//
// The file is guaranteed to be well-formed, and the number will always be
// between 0 and 9999.
//
// The program should be able to handle files with an arbitrary number of lines.
//
// You can use the following command to run the program:
//
// go run solution.go input.txt
//
// The test cases can be run with:
//
// go test
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run solution.go input.txt")
		os.Exit(1)
	}

	hands, err := readHands(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Sort the hands
	hands = sortHands(hands)
	fmt.Printf("----------------- Sorted --------------------\n")

	// Loop through the hands and call determineHandType for each one
	rank := 1
	total := 0
	for _, hand := range hands {
		fmt.Printf("%s: %d, %d\n", hand.cards, determineHandType(hand.cards), rank)
		total += hand.bid * rank
		rank++
	}

	fmt.Printf("Total: %d\n", total)

}

func readHands(filename string) ([]Hand, error) {
	var hands []Hand

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return hands, err
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return hands, fmt.Errorf("invalid line: %s", line)
		}

		num, err := strconv.Atoi(fields[1])
		if err != nil {
			return hands, fmt.Errorf("invalid number: %s", fields[1])
		}

		hands = append(hands, Hand{cards: fields[0], bid: num})
	}

	return hands, nil
}

// Implement a sortHands function that sorts the hands based on the cards in it.
// A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order, where A is the highest and 2 is the lowest.
//
// Every hand is exactly one type. From strongest to weakest, they are:
//
// 1. Five of a kind, where all five cards have the same label: AAAAA
// 2. Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// 3. Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// 4. Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// 5. Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// 6. One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// 7. High card, where all cards' labels are distinct: 23456
// Hands are primarily ordered based on type; for example, every full house is stronger than any three of a kind.
//
// If two hands have the same type, a second ordering rule takes effect. Start by comparing the first card in each hand. If these cards are different, the hand with the stronger first card is considered stronger. If the first card in each hand have the same label, however, then move on to considering the second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.
//
// So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger. Similarly, 77888 and 77788 are both a full house, but 77888 is stronger because its third card is stronger (and both hands have the same first and second card).
//
// Finally, if two hands have the same type and all five cards have the same label, the hands are considered equal. This applies to all hand types, including five of a kind.

// Define a type for a slice of Hand objects
type Hands []Hand

// Implement the sort.Interface methods for the Hands type

// Len returns the number of elements in the collection
func (h Hands) Len() int {
	return len(h)
}

// Less reports whether the element with index i should sort before the element with index j
func (h Hands) Less(i, j int) bool {
	// Here you need to implement your custom comparison logic
	// For example, you might compare the 'cards' or 'bid' fields of the Hand objects
	hand_i := determineHandType(h[i].cards)
	hand_j := determineHandType(h[j].cards)

	if hand_i == hand_j {
		return comparePositionCards(h[i].cards, h[j].cards)
	} else {
		return hand_i > hand_j
	}
}

// Swap swaps the elements with indexes i and j
func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func sortHands(hands []Hand) []Hand {
	sort.Sort(Hands(hands))
	return hands
}

func comparePositionCards(cards1 string, cards2 string) bool {
	// Compare each card in each hand, starting with the first card
	// If the cards are different, return the hand with the higher card
	// If the cards are the same, continue with the next card
	// If all cards are the same, return 0

	cardValues := make(map[rune]int)
	cardValues['A'] = 13
	cardValues['K'] = 12
	cardValues['Q'] = 11
	cardValues['J'] = 0
	cardValues['T'] = 9
	cardValues['9'] = 8
	cardValues['8'] = 7
	cardValues['7'] = 6
	cardValues['6'] = 5
	cardValues['5'] = 4
	cardValues['4'] = 3
	cardValues['3'] = 2
	cardValues['2'] = 1

	for i := 0; i < 5; i++ {
		if rune(cards1[i]) != rune(cards2[i]) {
			return cardValues[rune(cards1[i])] < cardValues[rune(cards2[i])]
		}
	}

	return true
}

// Implement a function that determines the type of hand from the list of cards in a string
func determineHandType(cards string) int {
	// See how many times the most frequent character appears
	maxCount, nextMaxCount := countMostFrequentChars(cards)

	switch maxCount {
	case 5:
		return 1
	case 4:
		return 2
	case 3:
		// If the most frequent character appears 3 times, it could be a three of a kind or a full house
		// If the second most frequent character appears 2 times, it's a full house
		// Otherwise, it's a three of a kind
		if nextMaxCount == 2 {
			return 3
		} else {
			return 4
		}
	case 2:
		// If the most frequent character appears 2 times, it could be a two pair or a one pair
		// If the second most frequent character appears 2 times, it's a two pair
		// Otherwise, it's a one pair
		if nextMaxCount == 2 {
			return 5
		} else {
			return 6
		}
	case 1:
		// If the most frequent character appears 1 time, it's a high card
		return 7
	}

	return 0
}

// Implement a function that searches a string for the characters
// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2 and returns the number of times
// that the character appearing the most times appears in the string
func countMostFrequentChars(cards string) (int, int) {
	chars := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	maxCount := 0
	nextMaxCount := 0
	most := rune(' ')
	nextMost := rune(' ')

	// Special case:  If the hand is JJJJJ, convert all the Js to As
	if strings.Count(cards, "J") == 5 {
		cards = strings.Replace(cards, "J", "A", -1)
	}

	jokerCount := strings.Count(cards, "J")

	for _, char := range chars {
		count := strings.Count(cards, string(char))
		if count > maxCount {
			maxCount = count
			most = char
		}
	}

	for _, char := range chars {
		if char != most {
			count := strings.Count(cards, string(char))
			if count > nextMaxCount {
				nextMaxCount = count
				nextMost = char
			}
		}
	}

	// Convert all the jokers to the most frequent character
	// unless the most frequent character is "J", in which case
	// convert all the jokers to the next most frequent character
	// and recalculate the counts
	if jokerCount > 0 {
		if most == 'J' {
			cards = strings.Replace(cards, "J", string(nextMost), -1)
		} else {
			cards = strings.Replace(cards, "J", string(most), -1)
		}
		maxCount, nextMaxCount = countMostFrequentChars(cards)
	}

	return maxCount, nextMaxCount
}

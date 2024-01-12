// The input file, given on the command line, contains on line of
// a string of "R" and "L" characters, which represent right and
// left directions for following a map of nodes, which follows
// in the file.
//
// The list of 3-character nodes are listed one per line, with this format:
// <node name> = (<left node name>, <right node name>)
// For example:
// HQV = (LSD, NCQ)
// TLQ = (VLQ, KVB)
// LND = (BFJ, LGF)
//
// The first node in the list is the starting node.  From there, we need
// to follow the directions in the input string to traverse through the
// nodes until we reach node named "ZZZ".
//
// If all the directions are followed, but we have not yet reached node "ZZZ",
// then the directions should be repeated starting from the beginning, until
// we do reach "ZZZ".

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Node represents a node in the map
type Node struct {
	Name      string
	LeftNode  string
	RightNode string
}

// Main function which opens the input file, reads the first line containing
// the list of directions, keeping the filehandle open for later reading of
// the node list.
func main() {
	// Open the input file
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening input file %s\n", os.Args[1])
		os.Exit(1)
	}
	defer inputFile.Close()

	// Create a bufio.Reader object to read the input file
	reader := bufio.NewReader(inputFile)

	// Read the first line of the input file, which contains the directions
	// to follow to traverse the nodes.
	directions, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input file %s\n", os.Args[1])
		os.Exit(1)
	}
	// Remove the trailing newline from the directions string
	directions = directions[:len(directions)-1]

	// Create a map of nodes, indexed by name.
	nodeMap := make(map[string]Node)

	fmt.Printf("Directions: %s\n", directions)

	// Read the list of nodes from the input file, into a slice of strings.
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		nodeName, leftNode, rightNode := parseNode(line)
		// print the three variables from the above line
		fmt.Printf("Node: %s, Left: %s, Right: %s\n", nodeName, leftNode, rightNode)

		// if nodeName blank, then skip this line
		if nodeName == "" {
			continue
		}

		nodeMap[nodeName] = Node{nodeName, leftNode, rightNode}
	}

	// Find all starting nodes, those that end in "A"
	currentNodes := make([]string, 0)
	for _, node := range nodeMap {
		if strings.HasSuffix(node.Name, "A") {
			currentNodes = append(currentNodes, node.Name)
		}
	}

	// Create a slice of ints to keep track of the number of steps taken for each starting node
	stepCount := make([]int, len(currentNodes))
	for i := range stepCount {
		stepCount[i] = 0
	}

	fmt.Printf("Starting nodes: %v\n", currentNodes)

	// Follow the directions, starting from all nodes that end in "A"
	// until all paths reach an node ending in "Z"

	// Iterate across all current nodes
	for i := 0; i < len(currentNodes); i++ {

		found := false
		for !found {
			for _, direction := range directions {

				// Get the Node struct for the current node.
				node := nodeMap[currentNodes[i]]

				// fmt.Printf("%02d: %s", i, currentNodes[i])

				// If the direction is "R", then follow the right node.
				if direction == 'R' {
					currentNodes[i] = node.RightNode
				} else {
					// Otherwise, follow the left node.
					currentNodes[i] = node.LeftNode
				}

				// fmt.Printf(" %s -> %s, %d\n", string(direction), currentNodes[i], stepCount[i])

				// Increment the number of steps taken.
				stepCount[i]++

				if strings.HasSuffix(currentNodes[i], "Z") {
					found = true
					break
				}

			}
		}

		fmt.Printf("Steps for path %d: %d\n", i, stepCount[i])
	}

	// Find the least common multiple of all the integers in stepCount
	lcm := lcmOfSlice(stepCount)
	fmt.Printf("Least common multiple: %d\n", lcm)
}

// parseNode parses a line from the input file, which contains a node name,
// and the names of the left and right nodes.
func parseNode(node string) (string, string, string) {
	re := regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)`)
	matches := re.FindStringSubmatch(node)
	if len(matches) != 4 {
		return "", "", ""
	}
	return matches[1], matches[2], matches[3]
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcmOfSlice(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = lcm(result, number)
	}
	return result
}

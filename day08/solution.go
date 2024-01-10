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
	// The first node is the starting node.
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

	// Follow the directions, starting at the first node.
	steps := 0
	currentNode := "AAA"

	for currentNode != "ZZZ" {
		for _, direction := range directions {
			// Get the Node struct for the current node.
			node := nodeMap[currentNode]

			fmt.Print(currentNode)

			// If the direction is "R", then follow the right node.
			if direction == 'R' {
				currentNode = node.RightNode
			} else {
				// Otherwise, follow the left node.
				currentNode = node.LeftNode
			}

			// Increment the number of steps taken.
			steps++

			fmt.Printf(" %s -> %s, %d\n", string(direction), currentNode, steps)

			// If we have reached the "ZZZ" node, then we are done.
			if currentNode == "ZZZ" {
				break
			}
		}
	}

	// Print the name of the node we ended up at.
	fmt.Println(steps)
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

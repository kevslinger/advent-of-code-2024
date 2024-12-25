package day25

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

type trieNode struct {
	count    int
	children []*trieNode
}

func RunDay25(path string) {
	answer, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", answer)
	}
}

func part1(file io.Reader) (int, error) {
	lines := readLines(file)
	keys, locks := parseKeysAndLocks(lines)
	keyTrie := convertSlicesToTrie(keys)
	return numFittingPairs(keyTrie, locks), nil
}

func readLines(file io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseKeysAndLocks(lines []string) ([][]int, [][]int) {
	var keys, locks [][]int
	for idx := 0; idx < len(lines); idx += 8 {
		// Key
		if lines[idx] == "....." {
			keys = append(keys, parseKeyOrLock(lines[idx:idx+8], '#'))
		} else if lines[idx] == "#####" {
			// Lock
			locks = append(locks, parseKeyOrLock(lines[idx:idx+8], '.'))
		}
	}
	return keys, locks
}

func parseKeyOrLock(lines []string, targetChar byte) []int {
	vals := make([]int, 5)
	for col := 0; col < len(lines[0]); col++ {
		row := 0
		for row < len(lines) && lines[row][col] != targetChar {
			row++
		}
		if targetChar == '#' {
			vals[col] = 6 - row
		} else {
			vals[col] = row - 1
		}
	}
	return vals
}

func convertSlicesToTrie(slices [][]int) *trieNode {
	root := &trieNode{children: make([]*trieNode, 6)}
	for _, slice := range slices {
		curNodes := []*trieNode{root}
		for _, elem := range slice {
			curNodesLen := len(curNodes)
			for steps := 0; steps < curNodesLen; steps++ {
				curNode := curNodes[0]
				curNodes = curNodes[1:]
				for val := elem; val <= 5; val++ {
					if curNode.children[val] == nil {
						curNode.children[val] = &trieNode{children: make([]*trieNode, 6)}
					}
					curNodes = append(curNodes, curNode.children[val])
				}
			}
		}
		for _, node := range curNodes {
			node.count++
		}
	}
	return root
}

func numFittingPairs(keyTrie *trieNode, locks [][]int) int {
	numPairs := 0
	for _, lock := range locks {
		curNode := keyTrie
		for _, val := range lock {
			if curNode.children[5-val] == nil {
				curNode = nil
				break
			}
			curNode = curNode.children[5-val]
		}
		if curNode != nil {
			numPairs += curNode.count
		}
	}
	return numPairs
}

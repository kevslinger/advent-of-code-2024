package day23

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay23(path string) {
	answer, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", answer)
	}
}

type computer struct {
	val              string
	connections      map[string]struct{}
	connectionsSlice []computer
}

func part1(file io.Reader) (int, error) {
	lines := readLines(file)
	computerNetwork := parseComputers(lines)
	return findThreeGroups(computerNetwork), nil
}

func readLines(file io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseComputers(lines []string) map[string]computer {
	computers := make(map[string]computer)
	for _, line := range lines {
		split := strings.Split(line, "-")
		computer1, computer2 := split[0], split[1]
		if _, ok := computers[computer1]; !ok {
			computers[computer1] = computer{val: computer1, connections: make(map[string]struct{}), connectionsSlice: make([]computer, 0)}
		}
		if _, ok := computers[computer2]; !ok {
			computers[computer2] = computer{val: computer2, connections: make(map[string]struct{})}
		}
		computer := computers[computer1]
		computer.connections[computer2] = struct{}{}
		computer.connectionsSlice = append(computer.connectionsSlice, computers[computer2])
		computers[computer1] = computer

		computer = computers[computer2]
		computer.connections[computer1] = struct{}{}
		computer.connectionsSlice = append(computer.connectionsSlice, computers[computer1])
		computers[computer2] = computer
	}
	return computers
}

func findThreeGroups(network map[string]computer) int {
	numGroups := 0
	visited := make(map[string]struct{})
	for _, computer := range network {
		if _, ok := visited[computer.val]; ok {
			continue
		}
		visited[computer.val] = struct{}{}
		for idx1, edge1 := range computer.connectionsSlice {
			if _, ok := visited[edge1.val]; ok {
				continue
			}
			for idx2 := idx1 + 1; idx2 < len(computer.connectionsSlice); idx2++ {
				edge2 := computer.connectionsSlice[idx2]
				if _, ok := visited[edge2.val]; ok {
					continue
				}
				if _, ok := edge1.connections[edge2.val]; ok {
					if strings.HasPrefix(computer.val, "t") || strings.HasPrefix(edge1.val, "t") || strings.HasPrefix(edge2.val, "t") {
						numGroups++
						fmt.Println(computer.val, edge1.val, edge2.val)
					}
				}
			}
		}
	}
	return numGroups
}

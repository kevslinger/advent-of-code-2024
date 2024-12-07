package day6

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay6(path string) {
	answer, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", answer)
	}
}

func part1(file io.Reader) (int, error) {
	grid := createGrid(file)
	startX, startY := locateGuard(grid)
	return simulateGuard(grid, startX, startY), nil
}

func createGrid(file io.Reader) []string {
	grid := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func locateGuard(grid []string) (int, int) {
	for rowIdx, row := range grid {
		for colIdx, val := range row {
			if val == '^' {
				return rowIdx, colIdx
			}
		}
	}
	return -1, -1
}

func simulateGuard(grid []string, x int, y int) int {
	// Start direction is north
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	curDirectionIdx := 0
	numPositionsVisited := 0
	visited := make(map[int]map[int]bool)
	for x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) {
		if _, ok := visited[x]; !ok {
			visited[x] = make(map[int]bool)
		}
		if !visited[x][y] {
			numPositionsVisited++
			visited[x][y] = true
		}
		// Check if moving in the current direction would lead to an obstacle
		newX := x + directions[curDirectionIdx][0]
		newY := y + directions[curDirectionIdx][1]
		if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[newX]) {
			break
		}
		if grid[newX][newY] == '#' {
			curDirectionIdx = (curDirectionIdx + 1) % 4
		} else {
			x = newX
			y = newY
		}
	}
	return numPositionsVisited
}

package day4

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay4(path string) {
	answer, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", answer)
	}

	answer, err = runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error in Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 2 is %d\n", answer)
	}
}

func part1(file io.Reader) (int, error) {
	xmasCount := 0
	grid := readGrid(file)
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for colIdx := 0; colIdx < len(grid[rowIdx]); colIdx++ {
			xmasCount += countXmasFromSpot(grid, rowIdx, colIdx)
		}
	}
	return xmasCount, nil
}

func part2(file io.Reader) (int, error) {
	masCount := 0
	grid := readGrid(file)
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for colIdx := 0; colIdx < len(grid[rowIdx]); colIdx++ {
			masCount += countMasFromSpot(grid, rowIdx, colIdx)
		}
	}
	return masCount, nil
}

func readGrid(file io.Reader) []string {
	grid := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	return grid
}

func countXmasFromSpot(grid []string, rowIdx int, colIdx int) int {
	if grid[rowIdx][colIdx] != 'X' {
		return 0
	}
	xmas := "XMAS"
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}
	xmasCount := 0
	for _, direction := range directions {
		newRowIdx := rowIdx
		newColIdx := colIdx
		xmasIdx := 0
		for newRowIdx >= 0 && newRowIdx < len(grid) && newColIdx >= 0 && newColIdx < len(grid[newRowIdx]) {
			if grid[newRowIdx][newColIdx] != xmas[xmasIdx] {
				break
			}
			xmasIdx++
			if xmasIdx >= len(xmas) {
				xmasCount++
				break
			}
			newRowIdx += direction[0]
			newColIdx += direction[1]
		}
	}
	return xmasCount
}

func countMasFromSpot(grid []string, rowIdx int, colIdx int) int {
	if rowIdx == 0 || rowIdx == len(grid)-1 || colIdx == 0 || colIdx == len(grid[rowIdx])-1 || grid[rowIdx][colIdx] != 'A' {
		return 0
	}

	topLeft := grid[rowIdx-1][colIdx-1]
	bottomRight := grid[rowIdx+1][colIdx+1]
	if topLeft == 'S' {
		if bottomRight != 'M' {
			return 0
		}
	} else if topLeft == 'M' {
		if bottomRight != 'S' {
			return 0
		}
	} else {
		return 0
	}

	topRight := grid[rowIdx-1][colIdx+1]
	bottomLeft := grid[rowIdx+1][colIdx-1]
	if topRight == 'S' {
		if bottomLeft != 'M' {
			return 0
		}
	} else if topRight == 'M' {
		if bottomLeft != 'S' {
			return 0
		}
	} else {
		return 0
	}
	return 1
}

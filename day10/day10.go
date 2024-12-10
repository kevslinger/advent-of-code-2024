package day10

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

type coord struct {
	x int
	y int
}

func (c coord) Add(x coord) coord {
	return coord{c.x + x.x, c.y + x.y}
}

var directions []coord = []coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func RunDay10(path string) {
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
	lines := parseFile(file)
	hill := parseHill(lines)
	startingCoords := findStartingCoordinates(hill)
	totalScore := 0
	for _, coord := range startingCoords {
		accessiblePeaks := findAccessiblePeaks(coord, hill)
		totalScore += len(accessiblePeaks)
	}
	return totalScore, nil
}

func part2(file io.Reader) (int, error) {
	lines := parseFile(file)
	hill := parseHill(lines)
	startingCoords := findStartingCoordinates(hill)
	totalScore := 0
	for _, coord := range startingCoords {
		uniqueTrails := findAccessiblePeaks(coord, hill)
		for _, count := range uniqueTrails {
			totalScore += count
		}
	}
	return totalScore, nil
}

func parseFile(file io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseHill(lines []string) [][]int {
	hill := make([][]int, len(lines))
	for linesIdx, line := range lines {
		lineInts := make([]int, len(line))
		for lineIdx, val := range line {
			lineInts[lineIdx] = int(val - '0')
		}
		hill[linesIdx] = lineInts
	}
	return hill
}

func findStartingCoordinates(hill [][]int) []coord {
	coords := make([]coord, 0)
	for x, row := range hill {
		for y, val := range row {
			if val == 0 {
				coords = append(coords, coord{x, y})
			}
		}
	}
	return coords
}

func findAccessiblePeaks(startingPoint coord, hill [][]int) map[coord]int {
	accessiblePeaks := make(map[coord]int)
	if hill[startingPoint.x][startingPoint.y] == 9 {
		accessiblePeaks[startingPoint] = 1
		return accessiblePeaks
	}
	for _, direction := range directions {
		newPoint := startingPoint.Add(direction)
		if newPoint.x < 0 || newPoint.x >= len(hill) || newPoint.y < 0 || newPoint.y >= len(hill[newPoint.x]) || hill[startingPoint.x][startingPoint.y]+1 != hill[newPoint.x][newPoint.y] {
			continue
		}
		newAccessiblePeaks := findAccessiblePeaks(newPoint, hill)
		for peak, count := range newAccessiblePeaks {
			accessiblePeaks[peak] += count
		}
	}
	return accessiblePeaks
}

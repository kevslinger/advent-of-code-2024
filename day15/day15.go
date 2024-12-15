package day15

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

const (
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'
	empty = '.'
	wall  = '#'
	robot = '@'
	box   = 'O'
)

var (
	upCoord    = coord{-1, 0}
	downCoord  = coord{1, 0}
	leftCoord  = coord{0, -1}
	rightCoord = coord{0, 1}
)

type coord struct {
	x int
	y int
}

func (c coord) Add(b coord) coord {
	return coord{x: c.x + b.x, y: c.y + b.y}
}

func RunDay15(path string) {
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
	lines := parseLines(file)
	mapLines := make([][]rune, 0)
	linesIdx := 0
	line := lines[linesIdx]
	for linesIdx < len(lines)-1 && len(line) > 0 {
		mapLines = append(mapLines, []rune(line))
		linesIdx++
		line = lines[linesIdx]
	}
	instructions := strings.Join(lines[linesIdx+1:], "")
	robotCoord, err := findRobotPosition(mapLines)
	if err != nil {
		return -1, err
	}
	mapLines = simulateRobot(mapLines, robotCoord, instructions)
	return sumGpsCoords(mapLines), nil
}

func part2(file io.Reader) (int, error) {
	return -1, nil
}

func parseLines(file io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func findRobotPosition(mapLines [][]rune) (coord, error) {
	for rowIdx, row := range mapLines {
		for colIdx, val := range row {
			if val == robot {
				return coord{x: rowIdx, y: colIdx}, nil
			}
		}
	}
	return coord{x: -1, y: -1}, fmt.Errorf("robot not found")
}

func simulateRobot(mapLines [][]rune, robotCoord coord, instructions string) [][]rune {
	var instructionCoord coord
	for _, instruction := range instructions {
		switch instruction {
		case up:
			instructionCoord = upCoord
		case down:
			instructionCoord = downCoord
		case left:
			instructionCoord = leftCoord
		case right:
			instructionCoord = rightCoord
		}
		nextPosition := robotCoord.Add(instructionCoord)
		switch mapLines[nextPosition.x][nextPosition.y] {
		case empty:
			mapLines[robotCoord.x][robotCoord.y] = empty
			mapLines[nextPosition.x][nextPosition.y] = robot
			robotCoord = nextPosition
		case box:
			newBoxPosition, err := findNewBoxPosition(mapLines, nextPosition, instructionCoord)
			if err == nil {
				mapLines[robotCoord.x][robotCoord.y] = empty
				mapLines[newBoxPosition.x][newBoxPosition.y] = box
				mapLines[nextPosition.x][nextPosition.y] = robot
				robotCoord = nextPosition
			}
		}
	}
	return mapLines
}

func findNewBoxPosition(mapLines [][]rune, boxCoord coord, directionCoord coord) (coord, error) {
	for mapLines[boxCoord.x][boxCoord.y] != empty {
		switch mapLines[boxCoord.x][boxCoord.y] {
		case wall:
			return coord{x: -1, y: -1}, fmt.Errorf("could not move box")
		case box:
			boxCoord = boxCoord.Add(directionCoord)
		}
	}
	return boxCoord, nil
}

func sumGpsCoords(mapLines [][]rune) int {
	sumCoords := 0
	for rowIdx, row := range mapLines {
		for colIdx, val := range row {
			if val == box {
				sumCoords += 100*rowIdx + colIdx
			}
		}
	}
	return sumCoords
}

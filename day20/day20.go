package day20

import (
	"bufio"
	"fmt"
	"io"
	"math"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

const (
	wall  = '#'
	empty = '.'
	start = 'S'
	end   = 'E'
)

var directions = []coord{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}

type coord struct {
	x int
	y int
}

func (c coord) add(x coord) coord {
	return coord{x: c.x + x.x, y: c.y + x.y}
}

func (c coord) isOpposite(x coord) bool {
	return c.x*-1 == x.x && c.y*-1 == x.y
}

func RunDay20(path string) {
	answer, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", answer)
	}
}

func part1(file io.Reader) (int, error) {
	lines := parseLines(file)
	grid := makeGrid(lines)
	startingPosition, err := findStartingPosition(grid)
	if err != nil {
		return -1, err
	}
	fastestTimeWithoutCheating, cheatTimes := runThroughGrid(grid, startingPosition)
	return cheatTimesSavingAtLeastXPicos(cheatTimes, fastestTimeWithoutCheating, 100), nil
}

func part2(file io.Reader) (int, error) {
	return -1, nil
}

func parseLines(file io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func makeGrid(lines []string) [][]rune {
	grid := make([][]rune, 0, len(lines))
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func findStartingPosition(grid [][]rune) (coord, error) {
	for x, row := range grid {
		for y, val := range row {
			if val == start {
				return coord{x: x, y: y}, nil
			}
		}
	}
	return coord{}, fmt.Errorf("start not found")
}

type raceState struct {
	pos           coord
	steps         int
	timesCheated  int
	lastDirection coord
}

func runThroughGrid(grid [][]rune, startingPosition coord) (int, []int) {
	fastestTimeWithoutCheating := math.MaxInt
	var cheatTimes []int
	var queue []raceState
	queue = append(queue, raceState{pos: startingPosition, steps: 0, timesCheated: 0, lastDirection: coord{}})
	for len(queue) > 0 {
		queueLen := len(queue)
		for times := 0; times < queueLen; times++ {
			curState := queue[0]
			// Pop from queue
			queue = queue[1:]
			// Check if we won
			if grid[curState.pos.x][curState.pos.y] == end {
				if curState.timesCheated > 0 {
					cheatTimes = append(cheatTimes, curState.steps)
				} else {
					fastestTimeWithoutCheating = min(fastestTimeWithoutCheating, curState.steps)
				}
				continue
			}
			// Can only cheat max of 2 picoseconds
			if grid[curState.pos.x][curState.pos.y] == wall {
				curState.timesCheated += 1
			}
			if curState.timesCheated > 2 {
				continue
			}
			// Find where we can move
			for _, direction := range directions {
				// Don't go back the opposite way
				if curState.lastDirection.isOpposite(direction) {
					continue
				}
				nextPos := curState.pos.add(direction)
				// Out of bounds check
				if nextPos.x < 0 || nextPos.x >= len(grid) || nextPos.y < 0 || nextPos.y >= len(grid[nextPos.x]) {
					continue
				}
				if grid[nextPos.x][nextPos.y] == wall {
					queue = append(queue, raceState{pos: nextPos, steps: curState.steps + 1, timesCheated: curState.timesCheated + 1, lastDirection: direction})
				} else {
					queue = append(queue, raceState{pos: nextPos, steps: curState.steps + 1, timesCheated: curState.timesCheated, lastDirection: direction})
				}

			}
		}
	}
	return fastestTimeWithoutCheating, cheatTimes
}

func cheatTimesSavingAtLeastXPicos(cheatTimes []int, fastestWithoutCheating int, numPicosToSave int) int {
	numCheats := 0
	for _, cheatTime := range cheatTimes {
		if fastestWithoutCheating-cheatTime >= numPicosToSave {
			numCheats++
		}
	}
	return numCheats
}

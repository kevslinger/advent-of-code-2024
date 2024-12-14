package day14

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

const (
	gridWidth  = 101
	gridHeight = 103
)

func RunDay14(path string) {
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
	robots, err := parseRobots(lines)
	if err != nil {
		return -1, err
	}
	for stepNum := 0; stepNum < 100; stepNum++ {
		simulateRobotStep(robots)
	}
	quadrantTotals := countRobotsInQuadrants(robots)
	quadrantProduct := 1
	for _, quadrantVal := range quadrantTotals {
		quadrantProduct *= quadrantVal
	}
	return quadrantProduct, nil
}

func part2(file io.Reader) (int, error) {
	lines := parseLines(file)
	robots, err := parseRobots(lines)
	if err != nil {
		return -1, err
	}
	for stepNum := 1; stepNum <= 1000; stepNum++ {
		fmt.Println(stepNum)
		simulateRobotStep(robots)
		robotGrid := plotRobots(robots)
		if isChristmasTree(robotGrid) {
			return stepNum, nil
		}
	}
	return -1, nil
}

type robot struct {
	pX int
	pY int
	vX int
	vY int
}

func parseLines(file io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseRobots(lines []string) ([]*robot, error) {
	numRegex := regexp.MustCompile(`-?[0-9]+`)
	robots := make([]*robot, 0)
	for _, line := range lines {
		nums := numRegex.FindAllString(line, -1)
		pX, err := strconv.Atoi(nums[0])
		if err != nil {
			return robots, err
		}
		pY, err := strconv.Atoi(nums[1])
		if err != nil {
			return robots, err
		}
		vX, err := strconv.Atoi(nums[2])
		if err != nil {
			return robots, err
		}
		vY, err := strconv.Atoi(nums[3])
		if err != nil {
			return robots, err
		}
		robots = append(robots, &robot{pX: pX, pY: pY, vX: vX, vY: vY})
	}
	return robots, nil
}

func simulateRobotStep(robots []*robot) {
	for _, robot := range robots {
		robot.pX = (robot.pX + robot.vX) % gridWidth
		if robot.pX < 0 {
			robot.pX = gridWidth + robot.pX
		}
		robot.pY = (robot.pY + robot.vY) % gridHeight
		if robot.pY < 0 {
			robot.pY = gridHeight + robot.pY
		}
	}
}

func countRobotsInQuadrants(robots []*robot) [4]int {
	var quadrantCounts [4]int
	for _, robot := range robots {
		if robot.pX < gridWidth/2 {
			if robot.pY < gridHeight/2 {
				quadrantCounts[0] += 1
			} else if robot.pY > gridHeight/2 {
				quadrantCounts[1] += 1
			}
		} else if robot.pX > gridWidth/2 {
			if robot.pY < gridHeight/2 {
				quadrantCounts[2] += 1
			} else if robot.pY > gridHeight/2 {
				quadrantCounts[3] += 1
			}
		}
	}
	return quadrantCounts
}

func plotRobots(robots []*robot) [][]int {
	grid := make([][]int, gridHeight)
	for idx := 0; idx < len(grid); idx++ {
		grid[idx] = make([]int, gridWidth)
	}
	for _, robot := range robots {
		grid[robot.pY][robot.pX] = 1
	}
	return grid
}

func isChristmasTree(grid [][]int) bool {
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for colIdx := 0; colIdx < len(grid[rowIdx]); colIdx++ {
			if grid[rowIdx][colIdx] == 1 && topsChristmasTree(grid, rowIdx, colIdx) {
				return true
			}
			if grid[rowIdx][colIdx] == 1 && bottomsChristmasTree(grid, rowIdx, colIdx) {
				fmt.Println(rowIdx, colIdx)
				fmt.Println(grid[rowIdx])
				return true
			}
		}
	}
	return false
}

func topsChristmasTree(grid [][]int, rowIdx int, colIdx int) bool {
	separator := 0
	// Keep going down until you can't
	for rowIdx < len(grid) && colIdx-separator >= 0 && colIdx+separator < len(grid[rowIdx]) && grid[rowIdx][colIdx-separator] == 1 && grid[rowIdx][colIdx+separator] == 1 {
		rowIdx++
		separator++
	}
	if separator > 2 {
		fmt.Println(rowIdx, separator)
		fmt.Println(grid[rowIdx-3])
		fmt.Println(grid[rowIdx-2])
		fmt.Println(grid[rowIdx-1])
		fmt.Println(grid[rowIdx])
	}

	// Need to go back up 1 row
	rowIdx--
	separator--
	// see if the bottom is a full row
	for cIdx := colIdx - separator; cIdx <= colIdx+separator; cIdx++ {
		if grid[rowIdx][colIdx] != 1 {
			return false
		}
	}
	return separator > 2
}

func bottomsChristmasTree(grid [][]int, rowIdx int, colIdx int) bool {
	separator := 0
	for colIdx-separator >= 0 && colIdx+separator < len(grid[rowIdx]) && grid[rowIdx][colIdx-separator] == 1 && grid[rowIdx][colIdx+separator] == 1 {
		separator++
	}
	return separator > 5
}

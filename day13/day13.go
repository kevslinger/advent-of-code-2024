package day13

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay13(path string) {
	answer, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", answer)
	}
}

type claw struct {
	a     button
	b     button
	prize coord
}

type button struct {
	loc  coord
	cost int
}

type coord struct {
	x int
	y int
}

func part1(file io.Reader) (int, error) {
	lines := parseFile(file)
	claws, err := parseClaws(lines)
	if err != nil {
		return -1, err
	}
	totalTokens := 0
	for _, claw := range claws {
		totalTokens += getTokensForPrize(claw)
	}
	return totalTokens, nil
}

func parseFile(file io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseClaws(lines []string) ([]claw, error) {
	claws := make([]claw, 0)
	numRegex := regexp.MustCompile("[0-9]+")
	for idx := 0; idx < len(lines); idx += 4 {
		aButton, err := parseButton(numRegex, lines[idx], 3)
		if err != nil {
			return claws, err
		}
		bButton, err := parseButton(numRegex, lines[idx+1], 1)
		if err != nil {
			return claws, err
		}
		prize, err := parseCoord(numRegex, lines[idx+2])
		if err != nil {
			return claws, err
		}
		claws = append(claws, claw{aButton, bButton, prize})
	}
	return claws, nil
}

func parseButton(numRegex *regexp.Regexp, coordsLine string, cost int) (button, error) {
	coord, err := parseCoord(numRegex, coordsLine)
	return button{loc: coord, cost: cost}, err
}

func parseCoord(numRegex *regexp.Regexp, coordsLine string) (coord, error) {
	coords := numRegex.FindAllString(coordsLine, -1)
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return coord{}, err
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return coord{}, err
	}
	return coord{x: x, y: y}, nil
}

func getTokensForPrize(c claw) int {
	mminimumTokens := math.MaxInt
	for bPresses := 0; bPresses <= 100; bPresses++ {
		xTargetRemaining := c.prize.x - c.b.loc.x*bPresses
		yTargetRemaining := c.prize.y - c.b.loc.y*bPresses
		aPresses := calcPressesForTarget(c.a, xTargetRemaining, yTargetRemaining)
		if aPresses >= 0 {
			mminimumTokens = min(mminimumTokens, bPresses*c.b.cost+aPresses*c.a.cost)
		}
	}
	if mminimumTokens == math.MaxInt {
		return 0
	}
	return mminimumTokens
}

func calcPressesForTarget(b button, targetX int, targetY int) int {
	if targetX%b.loc.x == 0 && targetY%b.loc.y == 0 && targetX/b.loc.x == targetY/b.loc.y {
		return targetX / b.loc.x
	}
	return -1
}

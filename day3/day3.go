package day3

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

var (
	mulRegex = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	numRegex = regexp.MustCompile("[0-9]+")
)

func RunDay3(path string) {
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
	totalSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := mulRegex.FindAllString(line, -1)
		for _, match := range matches {
			product, err := calculateProduct(numRegex.FindAllString(match, 2))
			if err != nil {
				return -1, err
			}
			totalSum += product
		}
	}
	return totalSum, nil
}

func part2(file io.Reader) (int, error) {
	totalSum := 0
	doExp := regexp.MustCompile(`do\(\)`)
	dontExp := regexp.MustCompile(`don't\(\)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mulMatches := mulRegex.FindAllStringIndex(line, -1)
		doMatches := doExp.FindAllStringIndex(line, -1)
		dontMatches := dontExp.FindAllStringIndex(line, -1)
		isEnabled := true
		doPtr := 0
		dontPtr := 0
		for _, match := range mulMatches {
			startingIdx := match[0]
			// Advance the pointers of the do and don't indices until they reach the next mul(...)
			for doPtr < len(doMatches) && startingIdx > doMatches[doPtr][0] {
				doPtr++
			}
			for dontPtr < len(dontMatches) && startingIdx > dontMatches[dontPtr][0] {
				dontPtr++
			}
			doPtr = max(0, doPtr-1)
			dontPtr = max(0, dontPtr-1)
			if startingIdx > dontMatches[dontPtr][0] && startingIdx > doMatches[doPtr][0] {
				if doMatches[doPtr][0] > dontMatches[dontPtr][0] {
					isEnabled = true
				} else {
					isEnabled = false
				}
			} else if startingIdx > dontMatches[dontPtr][0] {
				isEnabled = false
			}

			if isEnabled {
				product, err := calculateProduct(numRegex.FindAllString(line[match[0]:match[1]], 2))
				if err != nil {
					return -1, err
				}
				totalSum += product
			}
		}
	}
	return totalSum, nil
}

func calculateProduct(nums []string) (int, error) {
	num1, err := strconv.Atoi(nums[0])
	if err != nil {
		return -1, err
	}
	num2, err := strconv.Atoi(nums[1])
	if err != nil {
		return -1, err
	}
	return num1 * num2, nil
}

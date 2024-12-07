package day7

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay7(path string) {
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
	equations, err := parseEquations(lines)
	if err != nil {
		return -1, err
	}
	totalCalibration := 0
	for _, equation := range equations {
		if canBeValidEquation(equation[0], equation[1], equation[2:]) {
			totalCalibration += equation[0]
		}
	}
	return totalCalibration, nil
}

func part2(file io.Reader) (int, error) {
	lines := parseFile(file)
	equations, err := parseEquations(lines)
	if err != nil {
		return -1, err
	}
	totalCalibration := 0
	for _, equation := range equations {
		if canBeValidEquationPart2(equation[0], equation[1], equation[2:]) {
			totalCalibration += equation[0]
		}
	}
	return totalCalibration, nil
}

func parseFile(file io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseEquations(lines []string) ([][]int, error) {
	equations := make([][]int, 0)
	numExpr := regexp.MustCompile(`[0-9]+`)
	for _, line := range lines {
		equation := make([]int, 0)
		numbers := numExpr.FindAllString(line, -1)
		for _, number := range numbers {
			intNum, err := strconv.Atoi(number)
			if err != nil {
				return equations, err
			}
			equation = append(equation, intNum)
		}
		equations = append(equations, equation)
	}
	return equations, nil
}

func canBeValidEquation(targetSum int, currentSum int, equation []int) bool {
	if len(equation) == 0 {
		return targetSum == currentSum
	}
	// Can choose multiplication or addition
	return canBeValidEquation(targetSum, currentSum+equation[0], equation[1:]) || canBeValidEquation(targetSum, currentSum*equation[0], equation[1:])
}

func canBeValidEquationPart2(targetSum int, currentSum int, equation []int) bool {
	if len(equation) == 0 {
		return targetSum == currentSum
	}
	concatenatedNum, _ := strconv.Atoi(strconv.Itoa(currentSum) + strconv.Itoa(equation[0]))
	// Can choose multiplication or addition or concatenation
	return canBeValidEquationPart2(targetSum, currentSum+equation[0], equation[1:]) || canBeValidEquationPart2(targetSum, currentSum*equation[0], equation[1:]) || canBeValidEquationPart2(targetSum, concatenatedNum, equation[1:])
}

package day1

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay1(path string) {
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
	leftNums := make([]int, 0)
	rightNums := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		leftNums = append(leftNums, nums[0])
		rightNums = append(rightNums, nums[1])
	}
	slices.Sort(leftNums)
	slices.Sort(rightNums)

	totalSum := 0
	for idx := 0; idx < len(leftNums); idx++ {
		totalSum += int(math.Abs(float64(leftNums[idx] - rightNums[idx])))
	}
	return totalSum, nil
}

func part2(file io.Reader) (int, error) {
	leftNums := make([]int, 0)
	rightNumsCounts := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		leftNums = append(leftNums, nums[0])
		if _, ok := rightNumsCounts[nums[1]]; !ok {
			rightNumsCounts[nums[1]] = 0
		}
		rightNumsCounts[nums[1]]++
	}

	totalSum := 0
	for _, num := range leftNums {
		totalSum += num * rightNumsCounts[num]
	}
	return totalSum, nil
}

func parseLine(line string) ([]int, error) {
	nums := strings.Split(line, " ")
	leftNum, err := strconv.Atoi(nums[0])
	if err != nil {
		return make([]int, 0), err
	}
	// Sort out spacing issues with input text
	rightNumIdx := 1
	rightNum, err := strconv.Atoi(nums[rightNumIdx])
	for err != nil && rightNumIdx < len(nums) {
		rightNumIdx++
		rightNum, err = strconv.Atoi(nums[rightNumIdx])
	}
	if err != nil {
		return make([]int, 0), err
	}
	return []int{leftNum, rightNum}, nil
}

package day5

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay5(path string) {
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
	ordering := make(map[int]map[int]bool)
	scanner := bufio.NewScanner(file)
	var err error
	for scanner.Scan() {
		if len(ordering) == 0 {
			ordering, err = parseOrdering(scanner)
			if err != nil {
				return -1, err
			}
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		update, err := parseUpdate(line)
		if err != nil {
			return -1, err
		}
		if incorrectIndex, _ := findIncorrectIndices(update, ordering); incorrectIndex == -1 {
			totalSum += update[len(update)/2]
		}
	}
	return totalSum, nil
}

func part2(file io.Reader) (int, error) {
	totalSum := 0
	ordering := make(map[int]map[int]bool)
	scanner := bufio.NewScanner(file)
	var err error
	for scanner.Scan() {
		if len(ordering) == 0 {
			ordering, err = parseOrdering(scanner)
			if err != nil {
				fmt.Println("error is non-null")
				return -1, err
			}
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		update, err := parseUpdate(line)
		if err != nil {
			return -1, err
		}
		incorrectIndex, spotToPlace := findIncorrectIndices(update, ordering)
		if incorrectIndex != -1 {
			for incorrectIndex != -1 {
				update = slices.Insert(update, spotToPlace+1, update[incorrectIndex])
				update = slices.Delete(update, incorrectIndex, incorrectIndex+1)
				incorrectIndex, spotToPlace = findIncorrectIndices(update, ordering)
			}
			totalSum += update[len(update)/2]
		}
	}
	return totalSum, nil
}

func parseOrdering(scanner *bufio.Scanner) (map[int]map[int]bool, error) {
	ordering := make(map[int]map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		nums := strings.Split(line, "|")
		num0, err := strconv.Atoi(nums[0])
		if err != nil {
			return ordering, err
		}
		num1, err := strconv.Atoi(nums[1])
		if err != nil {
			return ordering, err
		}
		if _, ok := ordering[num1]; !ok {
			ordering[num1] = make(map[int]bool)
		}
		ordering[num1][num0] = true
	}
	return ordering, nil
}

func parseUpdate(line string) ([]int, error) {
	updateString := strings.Split(line, ",")
	update := make([]int, len(updateString))
	for idx, num := range updateString {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			return update, err
		}
		update[idx] = intNum
	}
	return update, nil
}

func findIncorrectIndices(update []int, ordering map[int]map[int]bool) (int, int) {
	firstIncorrectIndex := -1
	secondIncorrectIndex := -1
	for idx, num := range update {
		for idx2 := idx + 1; idx2 < len(update); idx2++ {
			if _, ok := ordering[num]; ok {
				if ordering[num][update[idx2]] {
					firstIncorrectIndex = idx
					secondIncorrectIndex = idx2
				}
			}
		}
		if firstIncorrectIndex != -1 {
			return firstIncorrectIndex, secondIncorrectIndex
		}
	}
	return firstIncorrectIndex, secondIncorrectIndex
}

// We need a way to gather rules
// map[int]map[int]bool
// maps after number to a set of before number(s)
// Then as we go through the update, for each number we need to check the numbers after it to see if any break the rules

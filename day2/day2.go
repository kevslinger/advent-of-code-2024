package day2

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

func RunDay2(path string) {
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
	safeReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report, err := formulateReport(scanner.Text())
		if err != nil {
			return -1, err
		}
		if isSafeReport(report) {
			safeReports++
		}
	}
	return safeReports, nil
}

func part2(file io.Reader) (int, error) {
	safeReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report, err := formulateReport(scanner.Text())
		if err != nil {
			return -1, err
		}
		// Check if the whole report is safe
		if isSafeReport(report) {
			safeReports++
			continue
		}
		// Cut off an element from the report and see if it's safe
		for i := 0; i < len(report); i++ {
			cutReport := slices.Clone(report)
			if isSafeReport(append(cutReport[:i], cutReport[i+1:]...)) {
				safeReports++
				break
			}
		}
	}
	return safeReports, nil
}

func formulateReport(line string) ([]int, error) {
	reportStr := strings.Split(line, " ")
	report := make([]int, len(reportStr))
	for idx, val := range reportStr {
		num, err := strconv.Atoi(val)
		if err != nil {
			return make([]int, 0), nil
		}
		report[idx] = num
	}
	return report, nil
}

func isSafeReport(report []int) bool {
	isIncreasing := true
	isDecreasing := true
	for idx := range report {
		if idx > 0 {
			diff := math.Abs(float64(report[idx] - report[idx-1]))
			if diff < 1 || diff > 3 {
				return false
			}
			if report[idx] <= report[idx-1] {
				isIncreasing = false
			}
			if report[idx] >= report[idx-1] {
				isDecreasing = false
			}
		}
		if !isIncreasing && !isDecreasing {
			return false
		}
	}
	return isIncreasing || isDecreasing
}

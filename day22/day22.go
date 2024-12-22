package day22

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

const simulationSteps = 2000

func RunDay22(path string) {
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
	lines := readLines(file)
	secretNumbers, err := getSecretNumbers(lines)
	if err != nil {
		return -1, err
	}
	finalSecretNumbers := simulateBuyers(secretNumbers)
	return getSecretNumberSum(finalSecretNumbers), nil
}

func part2(file io.Reader) (int, error) {
	lines := readLines(file)
	secretNumbers, err := getSecretNumbers(lines)
	if err != nil {
		return -1, err
	}
	bananaSequencesPerBuyer := simulateBuyersPart2(secretNumbers)
	return getHighestProfit(bananaSequencesPerBuyer), nil
}

func readLines(file io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getSecretNumbers(lines []string) ([]int, error) {
	var secretNumbers []int
	for _, line := range lines {
		secretNumber, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		secretNumbers = append(secretNumbers, secretNumber)
	}
	return secretNumbers, nil
}

func simulateBuyers(secretNumbers []int) []int {
	finalSecretNumbers := make([]int, len(secretNumbers))
	for idx, secretNumber := range secretNumbers {
		finalSecretNumbers[idx] = simulateBuyer(secretNumber)
	}
	return finalSecretNumbers
}

func simulateBuyersPart2(secretNumbers []int) []map[[4]int]int {
	bananaSequencesPerBuyer := make([]map[[4]int]int, len(secretNumbers))
	for idx, secretNumber := range secretNumbers {
		bananaSequencesPerBuyer[idx] = simulateBuyerPart2(secretNumber)
	}
	return bananaSequencesPerBuyer
}

func simulateBuyer(secretNumber int) int {
	for step := 0; step < simulationSteps; step++ {
		secretNumber = simulateBuyerOneStep(secretNumber)
	}
	return secretNumber
}

func simulateBuyerPart2(secretNumber int) map[[4]int]int {
	bananaSequences := make(map[[4]int]int)
	curPrice := secretNumber % 10
	var priceChanges [4]int
	for step := 0; step < simulationSteps; step++ {
		secretNumber = simulateBuyerOneStep(secretNumber)
		nextPrice := secretNumber % 10
		priceChange := nextPrice - curPrice
		priceChanges = [4]int{priceChanges[1], priceChanges[2], priceChanges[3], priceChange}
		if step >= 3 {
			if _, ok := bananaSequences[priceChanges]; !ok {
				bananaSequences[priceChanges] = nextPrice
			}
		}
		curPrice = nextPrice
	}
	return bananaSequences
}

func simulateBuyerOneStep(secretNumber int) int {
	secretNumber = firstStep(secretNumber)
	secretNumber = secondStep(secretNumber)
	return thirdStep(secretNumber)
}

func firstStep(secretNumber int) int {
	result := secretNumber * 64
	mixed := mixSecretNumber(secretNumber, result)
	return pruneSecretNumber(mixed)
}

func secondStep(secretNumber int) int {
	result := secretNumber / 32
	mixed := mixSecretNumber(secretNumber, result)
	return pruneSecretNumber(mixed)
}

func thirdStep(secretNumber int) int {
	result := secretNumber * 2048
	mixed := mixSecretNumber(secretNumber, result)
	return pruneSecretNumber(mixed)
}

func mixSecretNumber(secretNumber int, result int) int {
	return secretNumber ^ result
}

func pruneSecretNumber(secretNumber int) int {
	return secretNumber % 16777216
}

func getSecretNumberSum(secretNumbers []int) int {
	sum := 0
	for _, num := range secretNumbers {
		sum += num
	}
	return sum
}

func getHighestProfit(priceChangesPerBuyer []map[[4]int]int) int {
	visited := make(map[[4]int]struct{})
	highestProfit := 0
	for _, buyerPriceChanges := range priceChangesPerBuyer {
		for priceChange := range buyerPriceChanges {
			if _, ok := visited[priceChange]; !ok {
				highestProfit = max(highestProfit, getProfitFromPriceSequence(priceChangesPerBuyer, priceChange))
			}
			visited[priceChange] = struct{}{}
		}
	}
	return highestProfit
}

func getProfitFromPriceSequence(priceChangesPerBuyer []map[[4]int]int, priceChange [4]int) int {
	totalProfit := 0
	for _, buyer := range priceChangesPerBuyer {
		totalProfit += buyer[priceChange]
	}
	return totalProfit
}

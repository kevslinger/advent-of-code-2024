package day11

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"strings"
	"sync"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

type stone struct {
	val       *big.Int
	nextStone *stone
}

func (s *stone) String() string {
	if s.nextStone == nil {
		return s.val.String()
	}
	return s.val.String() + "," + s.nextStone.String()
}

func RunDay11(path string) {
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
	stones := parseStones(lines)
	simulateStones(stones, 25)
	numStones := 0
	for stones != nil {
		numStones++
		stones = stones.nextStone
	}
	return numStones, nil
}

func part2(file io.Reader) (int, error) {
	lines := parseFile(file)
	stones := parseStones(lines)
	simulateStones(stones, 75)
	numStones := 0
	for stones != nil {
		numStones++
		stones = stones.nextStone
	}
	return numStones, nil
}

func parseFile(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func parseStones(lines []string) *stone {
	val, _ := new(big.Int).SetString(lines[0], 0)
	headStone := &stone{val: val, nextStone: nil}
	prevStone := headStone
	for _, valStr := range lines[1:] {
		val, _ = new(big.Int).SetString(valStr, 0)
		stone := &stone{val: val, nextStone: nil}
		prevStone.nextStone = stone
		prevStone = stone
	}
	return headStone
}

func simulateStones(headStone *stone, numBlinks int) {
	for idx := 0; idx < numBlinks; idx++ {
		fmt.Println(idx)
		curStone := headStone
		wg := sync.WaitGroup{}
		for curStone != nil {
			wg.Add(1)
			go func(s *stone) {
				defer wg.Done()
				if s.val.Cmp(new(big.Int)) == 0 {
					s.val.Add(s.val, big.NewInt(1))
				} else if len(s.val.String())%2 == 0 {
					rightStr := s.val.String()[len(s.val.String())/2:]
					for len(rightStr) > 1 && rightStr[0] == '0' {
						rightStr = rightStr[1:]
					}
					rightVal, _ := new(big.Int).SetString(rightStr, 0)
					newStone := &stone{val: rightVal, nextStone: s.nextStone}
					s.nextStone = newStone
					leftStr := s.val.String()[:len(s.val.String())/2]
					leftVal, _ := new(big.Int).SetString(leftStr, 0)
					s.val = leftVal
				} else {
					s.val.Mul(s.val, big.NewInt(2024))
				}
			}(curStone)
			curStone = curStone.nextStone
		}
		wg.Wait()
	}
}

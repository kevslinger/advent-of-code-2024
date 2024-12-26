package day24

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

const (
	and = "AND"
	or  = "OR"
	xor = "XOR"
)

func RunDay24(path string) {
	answer, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", answer)
	}
}

func part1(file io.Reader) (int, error) {
	lines := readLines(file)
	wires, gates, err := parseWiresAndGates(lines)
	if err != nil {
		return -1, err
	}
	computeAllWireVals(wires, gates)
	return convertZValueBinaryToDecimal(wires), nil
}

func readLines(file io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseWiresAndGates(lines []string) (map[string]int, [][]string, error) {
	wires := make(map[string]int)
	gates := make([][]string, 0)
	linesIdx := 0
	for linesIdx < len(lines) && len(lines[linesIdx]) > 0 {
		wireParts := strings.Split(lines[linesIdx], ": ")
		initialVal, err := strconv.Atoi(wireParts[1])
		if err != nil {
			return wires, gates, err
		}
		wires[wireParts[0]] = initialVal
		linesIdx++
	}
	linesIdx++
	for linesIdx < len(lines) {
		gateParts := strings.Split(lines[linesIdx], " ")
		gate := []string{gateParts[0], gateParts[1], gateParts[2], gateParts[4]}
		gates = append(gates, gate)
		linesIdx++
	}
	return wires, gates, nil
}

func computeAllWireVals(wires map[string]int, gates [][]string) {
	seen := make(map[int]struct{})
	for len(gates) > len(seen) {
		for idx, gate := range gates {
			if _, ok := seen[idx]; ok {
				continue
			}
			val1, ok1 := wires[gate[0]]
			val2, ok2 := wires[gate[2]]
			if !ok1 || !ok2 {
				continue
			}
			op := gate[1]
			switch op {
			case and:
				wires[gate[3]] = val1 & val2
			case or:
				wires[gate[3]] = val1 | val2
			case xor:
				wires[gate[3]] = val1 ^ val2
			}
			seen[idx] = struct{}{}
		}
	}
}

func convertZValueBinaryToDecimal(wires map[string]int) int {
	res := 0
	zIdx := 0
	for {
		zIdxStr := strconv.Itoa(zIdx)
		if zIdx < 10 {
			zIdxStr = "0" + zIdxStr
		}
		zIdxStr = "z" + zIdxStr
		if wireVal, ok := wires[zIdxStr]; ok {
			res += wireVal << zIdx
		} else {
			return res
		}
		zIdx++
	}
}

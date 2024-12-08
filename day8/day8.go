package day8

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay8(path string) {
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
	m := parseMap(file)
	plantedAntinodes := plantAntinodes(m, plantAntinodePart1)
	return len(plantedAntinodes), nil
}

func part2(file io.Reader) (int, error) {
	m := parseMap(file)
	plantedAntinodes := plantAntinodes(m, plantAntinodePart2)
	return len(plantedAntinodes), nil
}

type Map struct {
	grid  []string
	nodes map[rune][]Coord
}

type Coord struct {
	x int
	y int
}

func parseMap(file io.Reader) Map {
	m := Map{nodes: make(map[rune][]Coord)}
	grid := make([]string, 0)
	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		for y, val := range line {
			if val != '.' {
				if _, ok := m.nodes[val]; !ok {
					m.nodes[val] = make([]Coord, 0)
				}
				m.nodes[val] = append(m.nodes[val], Coord{x, y})
			}
		}
		x++
	}
	m.grid = grid
	return m
}

func plantAntinodes(m Map, plantAntinodeFunc func(map[Coord]bool, Coord, Coord, int, int)) map[Coord]bool {
	numRows, numCols := len(m.grid), len(m.grid[0])
	plantedAntinodes := make(map[Coord]bool)
	for _, coords := range m.nodes {
		for idx1, coord1 := range coords {
			for idx2, coord2 := range coords {
				if idx1 == idx2 {
					continue
				}
				plantAntinodeFunc(plantedAntinodes, coord1, coord2, numRows, numCols)
				plantAntinodeFunc(plantedAntinodes, coord2, coord1, numRows, numCols)
			}
		}
	}
	return plantedAntinodes
}

func plantAntinodePart1(plantedAntinodes map[Coord]bool, coord1 Coord, coord2 Coord, numRows int, numCols int) {
	newAntenna := Coord{2*coord1.x - coord2.x, 2*coord1.y - coord2.y}
	if newAntenna.x >= 0 && newAntenna.x < numRows && newAntenna.y >= 0 && newAntenna.y < numCols {
		plantedAntinodes[newAntenna] = true
	}
}

func plantAntinodePart2(plantedAntinodes map[Coord]bool, coord1 Coord, coord2 Coord, numRows int, numCols int) {
	factorFirstCoord := 1
	factorSecondCoord := 0
	newAntenna := Coord{factorFirstCoord*coord1.x - factorSecondCoord*coord2.x, factorFirstCoord*coord1.y - factorSecondCoord*coord2.y}
	for newAntenna.x >= 0 && newAntenna.x < numRows && newAntenna.y >= 0 && newAntenna.y < numCols {
		plantedAntinodes[newAntenna] = true
		factorFirstCoord++
		factorSecondCoord++
		newAntenna.x = factorFirstCoord*coord1.x - factorSecondCoord*coord2.x
		newAntenna.y = factorFirstCoord*coord1.y - factorSecondCoord*coord2.y
	}
}

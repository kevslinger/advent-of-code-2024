package day9

import (
	"bufio"
	"fmt"
	"io"
	"math/big"

	"github.com/kevslinger/advent-of-code-2024/runner"
)

func RunDay9(path string) {
	answer, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", answer)
	}
}

func part1(file io.Reader) (int, error) {
	disk := readDisk(file)
	diskSlice := convertDiskToSlice(disk)
	diskSlice = compressDiskSlice(diskSlice)
	return int(computeChecksum(diskSlice).Int64()), nil
}

func readDisk(file io.Reader) string {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func convertDiskToSlice(disk string) []int {
	diskSlice := make([]int, 0)
	id := 0
	for diskIdx, char := range disk {
		val := int(char - '0')
		// File or free space
		if diskIdx%2 == 0 {
			for times := 0; times < val; times++ {
				diskSlice = append(diskSlice, id)
			}
			id++
		} else {
			for times := 0; times < val; times++ {
				diskSlice = append(diskSlice, -1)
			}
		}
	}
	return diskSlice
}

func compressDiskSlice(diskSlice []int) []int {
	leftPtr, rightPtr := 0, len(diskSlice)-1
	for leftPtr < rightPtr {
		if diskSlice[leftPtr] == -1 {
			for diskSlice[rightPtr] == -1 {
				rightPtr--
			}
			if leftPtr >= rightPtr {
				return diskSlice
			}
			diskSlice[leftPtr] = diskSlice[rightPtr]
			diskSlice[rightPtr] = -1
		}
		leftPtr++
	}
	return diskSlice
}

func computeChecksum(diskSlice []int) *big.Int {
	totalChecksum := new(big.Int)
	for idx, val := range diskSlice {
		if val == -1 {
			return totalChecksum
		}
		totalChecksum.Add(totalChecksum, big.NewInt(int64(idx*val)))
	}
	return totalChecksum
}

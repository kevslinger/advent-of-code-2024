package main

import (
	"fmt"

	"github.com/kevslinger/advent-of-code-2024/day1"
	"github.com/kevslinger/advent-of-code-2024/day2"
	"github.com/kevslinger/advent-of-code-2024/day3"
)

func main() {
	day1.RunDay1(GetInputDay(1))
	day2.RunDay2(GetInputDay(2))
	day3.RunDay3(GetInputDay(3))
}

func GetInputDay(day int) string {
	return fmt.Sprintf("./day%d/data/day%d.txt", day, day)
}

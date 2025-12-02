package main

import (
	"aoc-in-go/2025/lib"
	"math"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type move struct {
	direction string
	distance  int
}

func run(runPart2 bool, input string) any {
	if runPart2 {
		return part2(input)
	}

	return part1(input)
}

func part1(input string) any {
	lines := lib.SplitLines(input)
	moves := make([]move, len(lines))

	for i, line := range lines {
		distance, _ := strconv.Atoi(line[1:])
		moves[i] = move{
			direction: string(line[0]),
			distance:  distance,
		}
	}

	count := 0
	curr := 50

	for _, move := range moves {
		var next int

		if move.direction == "L" {
			next = curr - move.distance
		} else {
			next = curr + move.distance
		}

		next = next % 100
		if next < 0 {
			next = int(math.Abs(float64(100 + next)))
		} else if next > 99 {
			next = int(math.Abs(float64(100 - next)))
		}

		if next == 0 {
			count += 1
		}

		curr = next
	}

	return count
}

func part2(input string) any {
	lines := lib.SplitLines(input)
	moves := make([]move, len(lines))

	for i, line := range lines {
		distance, _ := strconv.Atoi(line[1:])
		moves[i] = move{
			direction: string(line[0]),
			distance:  distance,
		}
	}

	count := 0
	curr := 50

	for _, move := range moves {
		var next int

		fullRotations := move.distance / 100
		count += fullRotations

		if move.direction == "L" {
			next = curr - (move.distance % 100)
		} else {
			next = curr + (move.distance % 100)
		}

		if curr != 0 && next < 0 {
			count++
		} else if next > 99 {
			count++
		} else if curr != 0 && next == 0 {
			count++
		}

		realLocation := next
		if next < 0 {
			realLocation = int(math.Abs(float64(100 + next)))
		} else if next > 99 {
			realLocation = next % 100
		}

		curr = realLocation
	}

	return count
}

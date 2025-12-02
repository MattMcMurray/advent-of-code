package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type move struct {
	direction string
	distance  int
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	lines := SplitLines(input)
	moves := make([]move, len(lines))

	for i, line := range lines {
		fmt.Println(i, line)
		distance, _ := strconv.Atoi(line[1:])
		moves[i] = move{
			direction: string(line[0]),
			distance:  distance,
		}
	}

	count := 0
	curr := 50

	fmt.Println("Starting at: ", curr)

	for _, move := range moves {
		var next int

		if move.direction == "L" {
			next = curr - move.distance
		} else {
			next = curr + move.distance
		}

		fmt.Printf("Moved %d spaces to the %s\n", move.distance, move.direction)
		fmt.Printf("\tLanded on: %d\n", next)

		next = next % 100
		if next < 0 {
			next = int(math.Abs(float64(100 + next)))
		} else if next > 99 {
			next = int(math.Abs(float64(100 - next)))
		}
		fmt.Printf("\tWhich equates to: %d\n", next)

		if next == 0 {
			count += 1
		}

		curr = next
	}

	return count
}

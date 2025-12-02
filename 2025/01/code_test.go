package main

import "testing"

func TestPart2BaseCase(t *testing.T) {
	result := part2("L68")
	if result != 1 {
		t.Errorf("%d does not equal %d", result, 1)
	}
}

func TestPart2FullRotationsLeft(t *testing.T) {
	result := part2("L200")
	if result != 2 {
		t.Errorf("Full rotations miscalculated: %d", result)
	}
}

func TestPart2FullRotationsRight(t *testing.T) {
	result := part2("R200")
	if result != 2 {
		t.Errorf("Full rotations miscalculated: %d", result)
	}
}

func TestLandingOnZero(t *testing.T) {
	result := part2("L50\nL50\nL50")
	if result != 2 {
		t.Errorf("Miscalculation when landing on 0: %d", result)
	}
}

func TestPart2Example(t *testing.T) {
	input := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	result := part2(input)
	if result != 6 {
		t.Errorf("Result does not match expected sample: %d", result)
	}
}

func TestPart2DanceOnZero(t *testing.T) {
	input := "L50\nR1\nL1\nR1\nL1"
	result := part2(input)
	expected := 3
	if result != expected {
		t.Errorf("Result does not match expected value: %d %d", result, expected)
	}
}

func TestPart2LargeValue(t *testing.T) {
	input := "L500"
	result := part2(input)
	expected := 5
	if result != expected {
		t.Errorf("Result does not match expected value: %d %d", result, expected)
	}
}

func TestPart2FullRotationPlusZeroLanding(t *testing.T) {
	input := "L150\nR150"
	result := part2(input)
	expected := 3
	if result != expected {
		t.Errorf("Result does not match expected value: %d %d", result, expected)
	}
}

func TestPart2MultipleZeroLanding(t *testing.T) {
	input := "R50\nL100\nR100\nL100"
	result := part2(input)
	expected := 4
	if result != expected {
		t.Errorf("Result does not match expected value: %d %d", result, expected)
	}
}

func TestPart2MultipleZeroLandingInverse(t *testing.T) {
	input := "L50\nR100\nL100\nR100"
	result := part2(input)
	expected := 4
	if result != expected {
		t.Errorf("Result does not match expected value: %d %d", result, expected)
	}
}

func TestPart2MultipleZeroLandingLarge(t *testing.T) {
	input := "L50\nR1000"
	result := part2(input)
	expected := 11
	if result != expected {
		t.Errorf("Result does not match expected value: %d %d", result, expected)
	}
}

func TestPart2EdgeCase(t *testing.T) {
	input := "L48\nL25"
	result := part2(input)
	expected := 1
	if result != expected {
		t.Errorf("Result does not match expected value: %d %d", result, expected)
	}
}

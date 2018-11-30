package main

import (
	"testing"
)

func TestSumTurnWhenTurningOffGridOf1FromOff(t *testing.T) {
	lights := [][]int{{0}}
	count := LocalCount(lights, Line{command: TurnOff, startX: 0, startY: 0, endX: 0, endY: 0})
	if count != 0 {
		t.Fatal("count should be 0")
	}
}

func TestSumTurnWhenTurningOffGridOf1FromOn(t *testing.T) {
	lights := [][]int{{1}}
	count := LocalCount(lights, Line{command: TurnOff, startX: 0, startY: 0, endX: 0, endY: 0})
	if count != -1 {
		t.Fatal("count should be -1")
	}
}

func TestSumTurnWhenTurningOffGridOf2FromAllOn(t *testing.T) {
	lights := [][]int{{1, 1}, {1, 1}}
	count := LocalCount(lights, Line{command: TurnOff, startX: 0, startY: 0, endX: 1, endY: 1})
	if count != -4 {
		t.Fatalf("count should be -4 but is %d", count)
	}
}

func TestSumTurnWhenToggleGridOf2From3LightsOn(t *testing.T) {
	lights := [][]int{{1, 1}, {1, 0}}
	count := LocalCount(lights, Line{command: TurnToggle, startX: 0, startY: 0, endX: 1, endY: 1})
	if count != -2 {
		t.Fatalf("count should be -2 but is %d", count)
	}
}

package main

import "testing"

func testCoordinateParsing(t *testing.T, expectedCoordinate, actualCoordinate int) {
	if expectedCoordinate != actualCoordinate {
		t.Fatal("coordinate is not right")
	}
}

func TestStartXCoordinateTurnOn(t *testing.T) {
	testCoordinateParsing(t, 489, Parse("turn on 489,959 through 759,964").startX)
}

func TestStartXCoordinateTurnOffInput(t *testing.T) {
	testCoordinateParsing(t, 820, Parse("turn off 820,516 through 871,914").startX)
}

func TestStartXCoordinateTurnToggleInput(t *testing.T) {
	testCoordinateParsing(t, 756, Parse("toggle 756,965 through 812,992").startX)
}

func TestEndXCoordinateTurnOn(t *testing.T) {
	testCoordinateParsing(t, 759, Parse("turn on 489,959 through 759,964").endX)
}

func TestStartYCoordinateTurnOffInput(t *testing.T) {
	testCoordinateParsing(t, 516, Parse("turn off 820,516 through 871,914").startY)
}

func TestEndYCoordinateTurnToggleInput(t *testing.T) {
	testCoordinateParsing(t, 992, Parse("toggle 756,965 through 812,992").endY)
}

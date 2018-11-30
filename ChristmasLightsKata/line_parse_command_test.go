package main

import "testing"

func testCommandParsing(t *testing.T, input string, initialLightValue, expectedLightValue int) {
	var offLightActual = initialLightValue
	Parse(input).command(&offLightActual)
	if expectedLightValue != offLightActual {
		t.Fatal("command is not right")
	}
}

func TestTurnOnCommand(t *testing.T) {
	testCommandParsing(t, "turn on 489,959 through 759,964", 0, 1)
}

func TestTurnOffCommand(t *testing.T) {
	testCommandParsing(t, "turn off 820,516 through 871,914", 1, 0)
}

func TestToggleCommandStartOff(t *testing.T) {
	testCommandParsing(t, "toggle 756,965 through 812,992", 0, 1)
}
func TestToggleCommandStartOn(t *testing.T) {
	testCommandParsing(t, "toggle 756,965 through 812,992", 1, 0)
}

package main

import (
	"testing"
)

func TestTurnOffFrom0nLightWorks(t *testing.T) {
	light := [][]int{{1}}
	ToggleLights(&light, Line{command: TurnOff, startX: 0, startY: 0, endX: 0, endY: 0})
	if light[0][0] != 0 {
		t.Fatal("Light is still not off")
	}
}

func TestTurnOnFrom0Works(t *testing.T) {
	light := [][]int{{0}}
	ToggleLights(&light, Line{command: TurnOn, startX: 0, startY: 0, endX: 0, endY: 0})
	if light[0][0] != 1 {
		t.Fatal("Light is still not on")
	}
}

func TestToggleLightsFromOnWorks(t *testing.T) {
	light := [][]int{{1}}
	ToggleLights(&light, Line{command: TurnToggle, startX: 0, startY: 0, endX: 0, endY: 0})
	if light[0][0] != 0 {
		t.Fatal("Light is still not off")
	}
}

func TestToggleLightsFromOffWorks(t *testing.T) {
	light := [][]int{{0}}
	ToggleLights(&light, Line{command: TurnToggle, startX: 0, startY: 0, endX: 0, endY: 0})
	if light[0][0] != 1 {
		t.Fatal("Light is still not on")
	}
}

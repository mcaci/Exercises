package main

import (
	"testing"
)

func TestHueOffFrom0nLightWorks(t *testing.T) {
	light := [][]int{{1}}
	ToggleLights(&light, Line{command: HueOff, startX: 0, startY: 0, endX: 0, endY: 0})
	if light[0][0] != 0 {
		t.Fatal("Light is still not off")
	}
}
func TestHueOffFrom0ffLightWorks(t *testing.T) {
	light := [][]int{{0}}
	ToggleLights(&light, Line{command: HueOff, startX: 0, startY: 0, endX: 0, endY: 0})
	if light[0][0] != 0 {
		t.Fatal("Light is still not off")
	}
}

func TestHueOnFrom0Works(t *testing.T) {
	light := [][]int{{0}}
	ToggleLights(&light, Line{command: HueOn, startX: 0, startY: 0, endX: 0, endY: 0})
	if light[0][0] != 1 {
		t.Fatal("Light is still not on")
	}
}

func TestHueToggleLightsFromOffWorks(t *testing.T) {
	light := [][]int{{0}}
	ToggleLights(&light, Line{command: HueToggle, startX: 0, startY: 0, endX: 0, endY: 0})
	if light[0][0] != 2 {
		t.Fatal("Light is still not on")
	}
}

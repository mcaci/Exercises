package main

import "math"

var TurnOn = func(light *int) { (*light) = 1 }
var TurnOff = func(light *int) { (*light) = 0 }
var TurnToggle = func(light *int) { (*light) ^= 1 }

var HueOn = func(light *int) { (*light)++ }
var HueOff = func(light *int) { (*light) = int(math.Max(0, float64((*light)-1))) }
var HueToggle = func(light *int) { (*light) += 2 }

func ToggleLights(lights *[][]int, line Line) {
	for i := line.startX; i <= line.endX; i++ {
		for j := line.startY; j <= line.endY; j++ {
			line.command(&(*lights)[i][j])
		}
	}
}

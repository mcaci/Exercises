package main

import (
	"strconv"
	"strings"
)

type Line struct {
	command                    func(light *int)
	startX, startY, endX, endY int
}

func extractCommand(line string, isHue bool) func(*int) {
	var command func(*int)
	if !isHue {
		switch {
		case strings.HasPrefix(line, "turn on"):
			command = TurnOn
		case strings.HasPrefix(line, "turn off"):
			command = TurnOff
		case strings.HasPrefix(line, "toggle"):
			command = TurnToggle
		}
	} else {
		switch {
		case strings.HasPrefix(line, "turn on"):
			command = HueOn
		case strings.HasPrefix(line, "turn off"):
			command = HueOff
		case strings.HasPrefix(line, "toggle"):
			command = HueToggle
		}
	}
	return command
}

func Parse(line string, isHue bool) Line {
	command := extractCommand(line, isHue)

	eraseNotNum := func(r rune) rune {
		if r < '0' || r > '9' {
			return ' '
		}
		return r
	}
	numLine := strings.Map(eraseNotNum, line)
	nums := strings.Fields(numLine)
	startX, _ := strconv.Atoi(nums[0])
	startY, _ := strconv.Atoi(nums[1])
	endX, _ := strconv.Atoi(nums[2])
	endY, _ := strconv.Atoi(nums[3])

	return Line{command: command, startX: startX, startY: startY, endX: endX, endY: endY}
}

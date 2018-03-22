package dontpanictest

import . "dontpanic"

type dontPanicTestData struct {
	clonePos	int
	expectedDir	string
}

var elevators = []Elevator {
	Elevator{0, 5},
	Elevator{1, 3},
}
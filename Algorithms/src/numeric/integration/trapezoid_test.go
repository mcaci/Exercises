package integration

import "testing"

func TestTrapezoidWithOnePartition(t *testing.T) {
	area := Area(Trapezoid, testFunction, 0, 2, 1)
	var expectedArea float64 = 3
	verify(t, area, expectedArea)
}

func TestTrapezoidWithTwoPartitions(t *testing.T) {
	area := Area(Trapezoid, testFunction, 0, 2, 2)
	var expectedArea float64 = 3
	verify(t, area, expectedArea)
}
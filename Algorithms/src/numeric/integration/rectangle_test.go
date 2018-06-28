package integration

import "testing"

func TestRectangleWithOnePartition(t *testing.T) {
	area := Area(Rectangle, testFunction, 0, 2, 1)
	var expectedArea float64 = 4
	verify(t, area, expectedArea)
}

func TestRectangleWithTwoPartitions(t *testing.T) {
	area := Area(Rectangle, testFunction, 0, 2, 2)
	var expectedArea float64 = 3.5
	verify(t, area, expectedArea)
}

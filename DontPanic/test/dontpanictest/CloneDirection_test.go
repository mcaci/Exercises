package dontpanictest

import "testing"
import . "dontpanic"

func TestCloneShouldGoLeft(t *testing.T) {
	data := dataSetup(10, "LEFT")
	testDirection(data, t)
}

func TestCloneShouldGoRight(t *testing.T) {
	data := dataSetup(1 , "RIGHT")
	testDirection(data, t)
}

func (data dontPanicTestData) directionForElevator() string {
	elevator := elevators[1]
	return InWhichDirectionTheCloneShouldGo(data.clonePos, elevator.ElevatorPos)
}

func dataSetup(clonePosition int, expectedDirection string) dontPanicTestData {
	return dontPanicTestData {
		clonePos: clonePosition,
		expectedDir: expectedDirection,
	}
}

func testDirection(data dontPanicTestData, t *testing.T) {
	directionToTake := data.directionForElevator()
	if directionToTake != data.expectedDir {
		t.Fatalf("Expected %s but got %s", data.expectedDir, directionToTake)
	}
}
package dontpanic

type Clone struct {
	CloneFloor, ClonePos int
	Direction string
}

func (marvin *Clone) EvaluateAction(elevators []Elevator) string { 
	if marvin.hasCloneArrived() || marvin.isCloneAtAnElevator(elevators) {
		return "WAIT"
	} else {
		return marvin.action(elevators)
	}
}

func (marvin *Clone) action(elevators []Elevator) string {
	marvinExpectedDir := InWhichDirectionTheCloneShouldGo(marvin.ClonePos, elevators[marvin.CloneFloor].ElevatorPos)
	if marvinExpectedDir == marvin.Direction {
		return "WAIT"	
	} else {
		return "BLOCK"
	}
}

func (marvin *Clone) hasCloneArrived() bool {
	return marvin.CloneFloor < 0
}

func (marvin *Clone) isCloneAtAnElevator(elevators []Elevator) bool {
	return marvin.ClonePos == elevators[marvin.CloneFloor].ElevatorPos
}
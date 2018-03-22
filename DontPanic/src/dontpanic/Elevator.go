package dontpanic

type Elevator struct {
    ElevatorFloor, ElevatorPos int
}

type Elevators []Elevator

func (elevator Elevators) Len() int {
    return len(elevator)
}
func (elevator Elevators) Swap(i, j int) {
    elevator[i], elevator[j] = elevator[j], elevator[i]
}
func (elevator Elevators) Less(i, j int) bool {
    return elevator[i].ElevatorFloor < elevator[j].ElevatorFloor
}
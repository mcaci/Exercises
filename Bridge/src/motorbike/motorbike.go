package motorbike

type motorbike struct {
	Position, Speed, maxSpeed int
}

func newMotorbike(position, speed int) *motorbike {
	return &motorbike{Position : position, Speed : speed}
}

func (m *motorbike) nextCommand(b *bridge) string {
	command := "UNKNWON"
	if m.Position == 0 && m.Speed == 0 && b.RoadLength == 6 && b.GapLength == 4 {
		command = "SPEED"
	}
	if m.Position == 3 && m.Speed == 5 && b.RoadLength == 6 && b.GapLength == 4 {
		command = "WAIT"
	}
	if m.Position == 5 && m.Speed == 5 && b.RoadLength == 6 && b.GapLength == 4 {
		command = "JUMP"
	}
	if m.Position == 11 && m.Speed == 5 && b.RoadLength == 6 && b.GapLength == 4 {
		command = "SLOW"
	}
	if m.Position == 11 && m.Speed == 0 && b.RoadLength == 6 && b.GapLength == 4 {
		command = "WAIT"
	}
	return command
}

// func (m *motorbike) nextCommand2(b *bridge) string {
// 	command := "UNKNWON"
// 	if m.Position == 0 && m.Speed == 0 && b.RoadLength == 6 && b.GapLength == 4 {
// 		command = "SPEED"
// 	}
// 	if m.Position == 3 && m.Speed == 5 && b.RoadLength == 6 && b.GapLength == 4 {
// 		command = "WAIT"
// 	}
// 	if m.Position == 5 && m.Speed == 5 && b.RoadLength == 6 && b.GapLength == 4 {
// 		command = "JUMP"
// 	}
// 	if m.Position == 11 && m.Speed == 5 && b.RoadLength == 6 && b.GapLength == 4 {
// 		command = "SLOW"
// 	}
// 	if m.Position == 11 && m.Speed == 0 && b.RoadLength == 6 && b.GapLength == 4 {
// 		command = "WAIT"
// 	}
// 	return command
// }

func (m *motorbike) SetMaxSpeed(gapLength int) {
	m.maxSpeed = gapLength + 1
}

func (m *motorbike) GetMaxSpeed() int {
	return m.maxSpeed
}
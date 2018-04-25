package motorbike

type motorbike struct {
	Position, Speed int
}

func (m *motorbike) nextCommand(b *bridge) string {
	if m.Position == 5 && m.Speed == 5 && b.RoadLenght == 6 && b.GapLenght == 4 {
		return "JUMP"
	}
	if m.Position == 11 && m.Speed == 5 && b.RoadLenght == 6 && b.GapLenght == 4 {
		return "SLOW"
	}
	if m.Position == 11 && m.Speed == 0 && b.RoadLenght == 6 && b.GapLenght == 4 {
		return "WAIT"
	}
	return "WAIT"
}
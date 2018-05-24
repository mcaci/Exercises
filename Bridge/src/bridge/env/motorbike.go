package env

type Motorbike struct {
	Position, Speed, maxSpeed int
}

func NewMotorbike(position, speed int) *Motorbike {
	return &Motorbike{Position: position, Speed: speed}
}

func (m *Motorbike) SetMaxSpeed(gapLength int) {
	m.maxSpeed = gapLength + 1
}

func (m *Motorbike) GetMaxSpeed() int {
	return m.maxSpeed
}

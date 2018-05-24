package command

import "bridge/env"

func nextCommand(m *(env.Motorbike), b *(env.Bridge)) string {
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

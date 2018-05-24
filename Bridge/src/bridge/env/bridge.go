package env

type Bridge struct {
	RoadLength, GapLength, PlatformLength int
}

func NewBridge(road, gap, platform int) *Bridge {
	return &Bridge{RoadLength: road, GapLength: gap, PlatformLength: platform}
}

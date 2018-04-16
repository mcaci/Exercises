package path

func BuildGolfCourseWithPaths(paths [](*Path), golfMap []string) []string {
	for _, path := range paths {
		for _, direction := range *(path.StepSequence) {
			line := golfMap[path.Start.X()]
			golfMap[path.Start.X()] = line[:path.Start.Y()] + direction + line[path.Start.Y()+1:]
			path.Start.Move(direction)
		}
	}
	return golfMap
}

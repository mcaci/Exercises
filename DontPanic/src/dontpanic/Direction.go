package dontpanic

func InWhichDirectionTheCloneShouldGo(clonePos, elevatorPos int) string {
	if clonePos < elevatorPos {
		return "RIGHT"
	} else {
		return "LEFT"
	}
}
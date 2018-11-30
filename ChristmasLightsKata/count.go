package main

func LocalCount(cells [][]int, line Line) int {
	oldCount := countOnes(cells)
	ToggleLights(&cells, line)
	newCount := countOnes(cells)
	return newCount - oldCount
}

func countOnes(cells [][]int) int {
	count := 0
	for _, rows := range cells {
		for _, cell := range rows {
			count += cell
		}
	}
	return count
}

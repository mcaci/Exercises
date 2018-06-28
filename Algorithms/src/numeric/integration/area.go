package integration

func Area(
	areaFunction func(func(...float64) float64, float64, float64) float64,
	curveFunction func(...float64) float64,
	start, end float64,
	numberOfPartition uint) float64 {
		dx := (end - start) / float64(numberOfPartition)
		total_area := 0.0
		for x := start + dx; x <= end; x += dx {
			area := areaFunction(curveFunction, x, dx)
			total_area += area
		}
		return total_area
}

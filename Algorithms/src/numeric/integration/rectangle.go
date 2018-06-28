package integration

func Rectangle(testFunction func(parameters ...float64) float64, x, dx float64) float64 {
	return testFunction(x) * dx
}
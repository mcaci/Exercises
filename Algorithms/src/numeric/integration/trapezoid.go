package integration

func Trapezoid(testFunction func(parameters ...float64) float64, x, dx float64) float64 {
	return (testFunction(x - dx) + testFunction(x)) / 2 * dx
}

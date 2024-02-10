package calculator

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}
func multiply(a, b float64) float64 {
	return a * b
}
func divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

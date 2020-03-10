package math

func Calc(operation string) func(a, b float64) float64 {
	var temp func(a, b float64) float64
	switch operation {
	case "-":
		temp = func(a, b float64) (total float64) { return a - b }
	case "+":
		temp = func(a, b float64) (total float64) { return a + b }
	case "*":
		temp = func(a, b float64) (total float64) { return a * b }
	case "/":
		temp = func(a, b float64) (total float64) { return a / b }
	}

	return temp
}
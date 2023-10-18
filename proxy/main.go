package proxy

func calc(x, y int, op Operator) int {
	return op.calculate(x, y)
}

func main() {
	loggedAdder := NewLoggedAdder(&Adder{})
	calc(10, 5, loggedAdder)

	loggedSubtracter := NewLoggedSubtracter(&Subtracter{})
	calc(10, 5, loggedSubtracter)
}

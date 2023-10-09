package calculator

type Operator interface {
	calculate(x, y int) int
}

type Adder struct{}

func (a *Adder) calculate(x, y int) int {
	return x + y
}

type Subtractor struct{}

func (s *Subtractor) calculate(x, y int) int {
	return x - y
}

type Multiplier struct{}

func (m *Multiplier) calculate(x, y int) int {
	return x * y
}

type Divider struct{}

func (d *Divider) calculate(x, y int) int {
	return x / y
}

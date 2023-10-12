package adapter

type Multiplier interface {
	calc(x, y int) int
}

type Adapter struct {
	adder Adder
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (a Adapter) calc(x, y int) int {
	pos := (x >= 0 && y >= 0) || (x <= 0 && y <= 0)

	var sum int
	for i := 0; i < abs(y); i++ {
		sum += abs(x)
	}

	if !pos {
		return -sum
	}
	return sum
}

type Adder struct{}

func (a Adder) calc(x, y int) int {
	return x + y
}

package strategy

import (
	"fmt"

	"github.com/JunNishimura/design-pattern-with-go/strategy/calculator"
)

func main() {
	x := 10
	y := 5

	add := calculator.NewContext(&calculator.Adder{}).Calculate(x, y)
	sub := calculator.NewContext(&calculator.Subtractor{}).Calculate(x, y)
	mul := calculator.NewContext(&calculator.Multiplier{}).Calculate(x, y)
	div := calculator.NewContext(&calculator.Divider{}).Calculate(x, y)
	fmt.Println("add: ", add)
	fmt.Println("sub: ", sub)
	fmt.Println("mul: ", mul)
	fmt.Println("div: ", div)
}

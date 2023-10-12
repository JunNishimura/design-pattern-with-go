package adapter

import "fmt"

func main() {
	mult := Adapter{adder: Adder{}}
	fmt.Printf("3 * 5 = %d\n", mult.calc(3, 5))
	fmt.Printf("3 * 1 = %d\n", mult.calc(3, 1))
	fmt.Printf("3 * 0 = %d\n", mult.calc(3, 0))
	fmt.Printf("3 * -1 = %d\n", mult.calc(3, -1))
	fmt.Printf("-3 * -1 = %d\n", mult.calc(-3, -1))
}

package singleton

import (
	"fmt"

	"github.com/JunNishimura/design-pattern-with-go/singleton/counter"
)

func main() {
	c1 := counter.New()
	fmt.Println("counter: ", c1.Value())
	c1.Increment()
	c1.Increment()
	fmt.Println("counter: ", c1.Value())
	c2 := counter.New()
	c2.Decrement()
	fmt.Println("counter: ", c2.Value())
	fmt.Println("counter: ", c1.Value())
	fmt.Println("c1 == c2: ", c1 == c2)
}

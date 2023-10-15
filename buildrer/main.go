package builder

import (
	"fmt"

	"github.com/JunNishimura/design-pattern-with-go/buildrer/animal"
)

func main() {
	catBuilder := animal.NewCatBuilder()
	director := animal.NewDirector(catBuilder)
	director.BuildCat()
	fmt.Printf("%+v\n", catBuilder.Get())

	dogBuilder := animal.NewDogBuilder()
	director = animal.NewDirector(dogBuilder)
	director.BuildDog()
	fmt.Printf("%+v\n", dogBuilder.Get())
}

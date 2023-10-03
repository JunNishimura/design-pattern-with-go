package counter

var instance *Counter

type Counter struct {
	Value int
}

func New() *Counter {
	if instance == nil {
		instance = new(Counter)
	}
	return instance
}

func (c *Counter) Increment() {
	c.Value++
}

func (c *Counter) Decrement() {
	c.Value--
}

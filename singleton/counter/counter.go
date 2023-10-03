package counter

var instance *Counter

type Counter struct {
	value int
}

func New() *Counter {
	if instance == nil {
		instance = new(Counter)
	}
	return instance
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Decrement() {
	c.value--
}

package calculator

type Context struct {
	strategy Computer
}

func NewContext(c Computer) *Context {
	return &Context{
		strategy: c,
	}
}

func (c *Context) Calculate(x, y int) int {
	return c.strategy.calculate(x, y)
}

package calculator

type Context struct {
	strategy Operator
}

func NewContext(c Operator) *Context {
	return &Context{
		strategy: c,
	}
}

func (c *Context) Calculate(x, y int) int {
	return c.strategy.calculate(x, y)
}

package proxy

import "log"

type Operator interface {
	calculate(x, y int) int
}

type Adder struct{}

func (a *Adder) calculate(x, y int) int {
	return x + y
}

type LoggedAdder struct {
	adder *Adder
}

func NewLoggedAdder(adder *Adder) *LoggedAdder {
	return &LoggedAdder{adder: adder}
}

func (a *LoggedAdder) calculate(x, y int) int {
	result := a.adder.calculate(x, y)
	log.Printf("%d + %d = %d\n", x, y, result)
	return result
}

type Subtracter struct{}

func (s *Subtracter) calculate(x, y int) int {
	return x - y
}

type LoggedSubtracter struct {
	subtracter *Subtracter
}

func NewLoggedSubtracter(subtracter *Subtracter) *LoggedSubtracter {
	return &LoggedSubtracter{subtracter: subtracter}
}

func (s *LoggedSubtracter) calculate(x, y int) int {
	result := s.subtracter.calculate(x, y)
	log.Printf("%d - %d = %d\n", x, y, result)
	return result
}

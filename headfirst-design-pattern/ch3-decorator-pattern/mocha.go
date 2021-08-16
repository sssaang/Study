package main

type Mocha struct {
	beverage Beverage
}

func (m *Mocha) description() string {
	return m.beverage.description() + ", Mocha"
}

func (m *Mocha) cost() float32 {
	return m.beverage.cost() + .2
}
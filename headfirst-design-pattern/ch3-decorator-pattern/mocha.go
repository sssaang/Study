package main

type Mocha struct {
	beverage Beverage
}

func (m *Mocha) decription() string {
	return m.beverage.decription() + ", Mocha"
}

func (m *Mocha) cost() float32 {
	return m.beverage.cost() + .2
}
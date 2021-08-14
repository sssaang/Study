package main

type Milk struct {
	beverage Beverage
}

func (m *Milk) description() string {
	return m.beverage.description() + ", Milk"
}

func (m *Milk) cost() float32 {
	return m.beverage.cost() + .4
}
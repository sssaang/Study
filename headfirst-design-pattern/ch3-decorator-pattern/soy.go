package main

type Soy struct {
	beverage Beverage
}

func (s *Soy) description() string {
	return s.beverage.description() + ", Soy"
}

func (s *Soy) cost() float32 {
	return s.beverage.cost() + .2
}
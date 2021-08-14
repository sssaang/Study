package main

type Whip struct {
	beverage Beverage
}

func (w *Whip) description() string {
	return w.beverage.description() + ", Whip"
}

func (w *Whip) cost() float32 {
	return w.beverage.cost() + .2
}
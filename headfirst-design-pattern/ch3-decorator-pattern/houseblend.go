package main

type Houseblend struct {}

func (hb *Houseblend) description() string {
	return "House Blend coffee"
}

func (hb *Houseblend) cost() float32 {
	return 0.80
} 
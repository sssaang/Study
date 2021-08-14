package main

type Espresso struct {}

func (ep *Espresso) description() string {
	return "Espresso"
}

func (ep *Espresso) cost() float32 {
	return 1.12
} 
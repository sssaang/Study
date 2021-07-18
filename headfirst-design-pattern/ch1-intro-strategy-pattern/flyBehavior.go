package main

import "fmt"

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct {}

func (fww *FlyWithWings) fly() {
	fmt.Println("I'm Flying!")
}

type FlyNoWay struct {}

func (fnw *FlyNoWay) fly() {
	fmt.Println("I can't fly... 🥺")
}
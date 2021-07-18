package main

import "fmt"

func main() {
	mallardDuck := newMallardDuck()

	fmt.Println("Mallard duck performs")
	mallardDuck.performFly()
	mallardDuck.performQuack()

	fmt.Println("Model duck performs")
	modelDuck := newModelDuck()
	modelDuck.performFly()
	modelDuck.performQuack()

	mallardDuck.setFlyBehavior(&FlyNoWay{})
	mallardDuck.setQuackBehavior(&MuteQuack{})

	fmt.Println("dynamically changed Mallard duck performs")
	mallardDuck.performFly()
	mallardDuck.performQuack()
}
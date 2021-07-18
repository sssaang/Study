package main

import (
	"fmt"
)

type Duck struct {
	display func()
	fb FlyBehavior
	qb QuackBehavior
}

func (d *Duck) performFly() {
	d.fb.fly()
}
func (d *Duck) performQuack() {
	d.qb.quack()
}

func (d *Duck) setFlyBehavior(fb FlyBehavior) {
	d.fb = fb
}

func (d *Duck) setQuackBehavior(qb QuackBehavior) {
	d.qb = qb
}

type mallardDuck struct {
	*Duck
}

func newMallardDuck() *mallardDuck {
	// implementing the abstract method
	d := &Duck{
		 display: func() {
				fmt.Println("I’m a real Mallard duck")
		 },
		 fb: &FlyWithWings{},
		 qb: &Quack{},
	}

	return &mallardDuck{d}
}

type modelDuck struct {
	*Duck
}

func newModelDuck() *modelDuck {
	// implementing the abstract method
	d := &Duck{
		 display: func() {
				fmt.Println("I’m a real Model duck")
			},
			fb: &FlyNoWay{},
			qb: &MuteQuack{},
	}

	return &modelDuck{d}
}
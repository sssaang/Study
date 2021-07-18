package main

import "fmt"

type QuackBehavior interface {
	quack()
}

type Quack struct{}
func (q *Quack) quack() {
	fmt.Println("Quack! 🐥")
}

type Squeak struct{}
func (sq *Squeak) quack() {
	fmt.Println("Squeak!")
}

type MuteQuack struct{}
func (mq *MuteQuack) quack() {
	fmt.Print("Slience...")
}
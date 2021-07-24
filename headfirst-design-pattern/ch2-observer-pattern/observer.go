package main

type Observer interface {
	update(temperature float32, humidity float32, pressure float32)
	display()
}
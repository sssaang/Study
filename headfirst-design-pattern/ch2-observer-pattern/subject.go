package main

type Subject interface {
	addObserver(obs *Observer)
	removeObserver(obs *Observer)
	notifyObserver()
	setChanged()
}
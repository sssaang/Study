package main

import "fmt"

type CurrentConditionsDisplay struct {
	temperature float32
	humidity float32
  pressure float32
	WeatherStation *WeatherStation
}

func newCurrentConditionsDisplay() *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{}
}

func (ccd *CurrentConditionsDisplay) update(temperature float32, humidity float32, pressure float32) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.pressure = pressure
	ccd.display()
}

func (ccd CurrentConditionsDisplay) display() {
	fmt.Printf("Current conditions: Temperature:%.2f, Humidity:%.2f and Pressure:%.2f\n", ccd.temperature, ccd.humidity, ccd.pressure)
}
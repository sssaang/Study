package main

import "fmt"

type CurrentConditionsDisplay struct {
	temperature float32
	humidity float32
  pressure float32
	weatherStation *WeatherStation
}

func newCurrentConditionsDisplay(ws *WeatherStation) *CurrentConditionsDisplay {
	ccd := &CurrentConditionsDisplay{
		weatherStation: ws,
	}
	ws.addObserver(ccd)
	return ccd
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
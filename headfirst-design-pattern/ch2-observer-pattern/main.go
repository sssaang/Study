package main

func main() {
	ws := newWeatherStation(80.0, 65.0, 40.0)
	newCurrentConditionsDisplay(ws)

	ws.setChanged()
	ws.notifyObserver()

	ws.setMeasurements(85.0, 15.0, 440.0)

	newCurrentConditionsDisplay(ws)
	newCurrentConditionsDisplay(ws)

	ws.setMeasurements(1.0, 1.0, 1.0)
}
package main

type WeatherStation struct {
	changed bool
	temperature float32
	humidity float32
  pressure float32
	observers []Observer
}

func newWeatherStation(temperature float32, humidity float32, pressure float32) *WeatherStation {
	return &WeatherStation{
		changed: false,
		temperature: temperature,
		humidity: humidity,
		pressure: pressure,
		observers: []Observer{},
	}
}

func (ws *WeatherStation) setMeasurements(temperature float32, humidity float32, pressure float32) {
	ws.temperature = temperature
  ws.humidity = humidity
  ws.pressure = pressure
	ws.setChanged()
	ws.notifyObserver()
}

func (ws *WeatherStation) addObserver(obsToAdd Observer) {
	ws.observers = append(ws.observers, obsToAdd)
}

func (ws *WeatherStation) removeObserver(obsToRemove Observer) {
	for i, obs := range ws.observers {
		if obs == obsToRemove {
			ws.observers = append(ws.observers[:i], ws.observers[i+1:]...)
		}
	}
}

func (ws *WeatherStation) notifyObserver() {
	if ws.changed {
		for _, obs := range ws.observers {
			obs.update(ws.temperature, ws.humidity ,ws.pressure)
		}
	}
}

func (ws *WeatherStation) setChanged() {
	ws.changed = true
}
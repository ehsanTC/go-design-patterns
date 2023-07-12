package observer

import (
	"math/rand"

	"github.com/ehsanTC/sliceExt"
)

type WeatherStation struct {
	humidityObservers  []HumidityObserver
	windSpeedObservers []WindSpeedObserver
}

func (w *WeatherStation) AddHumidityObserver(observer HumidityObserver) HumidityObserver {
	if w.humidityObservers == nil {
		w.humidityObservers = make([]HumidityObserver, 0)
	}

	sliceExt.Add(&w.humidityObservers, observer)
	return observer
}

func (w *WeatherStation) RemoveHumidityObserver(observer HumidityObserver) HumidityObserver {
	if w.humidityObservers == nil {
		return nil
	}

	sliceExt.Remove(&w.humidityObservers, observer)
	return observer
}

func (w *WeatherStation) AddWindSpeedObserver(observer WindSpeedObserver) WindSpeedObserver {
	if w.windSpeedObservers == nil {
		w.windSpeedObservers = make([]WindSpeedObserver, 0)
	}

	sliceExt.Add(&w.windSpeedObservers, observer)
	return observer
}

func (w *WeatherStation) RemoveWindSpeedObserver(observer WindSpeedObserver) WindSpeedObserver {
	if w.windSpeedObservers == nil {
		return nil
	}

	sliceExt.Remove(&w.windSpeedObservers, observer)
	return observer
}

func (w *WeatherStation) MeasureHumidity() int {
	humidity := rand.Intn(100)

	// Inform observers
	if w.humidityObservers != nil {
		for _, observer := range w.humidityObservers {
			observer.HumidityMeasured(humidity)
		}
	}

	return humidity
}

func (w *WeatherStation) MeasureWindSpeed() int {
	max := 75 // mph - hurricane
	speed := rand.Intn(max)

	// Inform observers
	if w.windSpeedObservers != nil {
		for _, observer := range w.windSpeedObservers {
			observer.WindSpeedMeasured(speed)
		}
	}

	return speed
}

package observer

import (
	"fmt"
	"time"
)

type Observer struct {
	number int
}

func (o *Observer) Number() int {
	return o.number
}

func (o *Observer) SetNumber(n int) {
	o.number = n
}

func (o *Observer) Name() string {
	return "Observer"
}

func (o *Observer) Work() {
	farmer := &Farmer{lowerHumidity: 15, UpperHumidity: 60, workingState: true}
	roofWorker := &RoofWorker{windSpeedThreshold: 10, workingState: true}
	weatherStation := &WeatherStation{}

	weatherStation.AddHumidityObserver(farmer)
	weatherStation.AddWindSpeedObserver(roofWorker)

	fmt.Println("\nThe proper humidity for farmer is between 15% to 60%")
	fmt.Println("The proper wind speed for roof worker is blow 10 mph\n")

	fmt.Println("\nMeasuring wind speed and humidity at 6 am")
	fmt.Println("humidity: ", weatherStation.MeasureHumidity(), "%")
	fmt.Println("wind speed: ", weatherStation.MeasureWindSpeed(), "mph")
	time.Sleep(time.Second * 2)

	fmt.Println("\n\nMeasuring wind speed and humidity at 8:00 am")
	fmt.Println("humidity: ", weatherStation.MeasureHumidity(), "%")
	fmt.Println("wind speed: ", weatherStation.MeasureWindSpeed(), "mph")
	time.Sleep(time.Second * 2)

	fmt.Println("\n\nMeasuring wind speed and humidity at 12:00 am")
	fmt.Println("humidity: ", weatherStation.MeasureHumidity(), "%")
	fmt.Println("wind speed: ", weatherStation.MeasureWindSpeed(), "mph")
	time.Sleep(time.Second * 2)

	fmt.Println("\n\nMeasuring wind speed and humidity at 14:00 pm")
	fmt.Println("humidity: ", weatherStation.MeasureHumidity(), "%")
	fmt.Println("wind speed: ", weatherStation.MeasureWindSpeed(), "mph")
	time.Sleep(time.Second * 2)

	fmt.Println("\n\nMeasuring wind speed and humidity at 16:00 pm")
	fmt.Println("humidity: ", weatherStation.MeasureHumidity(), "%")
	fmt.Println("wind speed: ", weatherStation.MeasureWindSpeed(), "mph")
	time.Sleep(time.Second * 2)

	fmt.Println("\n\nMeasuring wind speed and humidity at 18:00 pm")
	fmt.Println("humidity: ", weatherStation.MeasureHumidity(), "%")
	fmt.Println("wind speed: ", weatherStation.MeasureWindSpeed(), "mph")
}

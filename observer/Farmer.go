package observer

import "fmt"

type Farmer struct {
	lowerHumidity int
	UpperHumidity int

	workingState bool
}

func (f *Farmer) HumidityMeasured(humidity int) {
	if humidity >= f.lowerHumidity && humidity <= f.UpperHumidity {
		if f.workingState == true {
			fmt.Println("Farmer is planting wheat.")
		} else {
			fmt.Println("Farmer started to planting wheat again.")
			f.workingState = true
		}
	} else {
		if f.workingState {
			var weatherType string
			if humidity < f.lowerHumidity {
				weatherType = "dried"
			} else {
				weatherType = "wet"
			}

			fmt.Println(fmt.Sprintf("Farmer can not plant in %s weather. Stopped farming!", weatherType))
			f.workingState = false
		} else {
			fmt.Println("Weather is still too wet. Waiting!")
		}
	}
}

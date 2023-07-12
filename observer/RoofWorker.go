package observer

import "fmt"

type RoofWorker struct {
	windSpeedThreshold int

	workingState bool
}

func (r *RoofWorker) WindSpeedMeasured(speed int) {
	if speed < r.windSpeedThreshold {
		if r.workingState {
			fmt.Println("Roof worker continues working.")
		} else {
			fmt.Println("Roof worker started working!")
			r.workingState = true
		}
	} else {
		if r.workingState {
			fmt.Println("Roof worker stopped working. Wind speed is very high!")
			r.workingState = false
		} else {
			fmt.Println("Roof worker still Waiting. The wind speed is very high.")
		}
	}
}

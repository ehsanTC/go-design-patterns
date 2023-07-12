package observer

type HumidityObserver interface {
	HumidityMeasured(int)
}

type WindSpeedObserver interface {
	WindSpeedMeasured(int)
}

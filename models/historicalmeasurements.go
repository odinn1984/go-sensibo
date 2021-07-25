package models

type HistoricalMeasurements struct {
	Temperature []struct {
		Time  string
		Value float64
	}
	Humidity []struct {
		Time  string
		Value float64
	}
}

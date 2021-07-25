package models

type ACStateData struct {
	On                bool
	Mode              string
	FanLevel          string
	TargetTemperature float64
	TemperatureUnit   string
	Swing             string
}

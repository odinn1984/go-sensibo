// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// ACStateData hold information about the state of the AC.
type ACStateData struct {
	On                bool   `json:"on"`
	Mode              string `json:"mode"`
	FanLevel          string `json:"fanLevel"`
	TargetTemperature int    `json:"targetTemperature"`
	TemperatureUnit   string `json:"temperatureUnit"`
	Swing             string `json:"swing"`
}

// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// ACStateData hold information about the state of the AC.
type ACStateData struct {
	On                bool
	Mode              string
	FanLevel          string
	TargetTemperature float64
	TemperatureUnit   string
	Swing             string
}

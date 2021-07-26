// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// Information about historical measurements.
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

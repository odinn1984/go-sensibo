// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// DeviceTimer holds information about the timer on the device.
type DeviceTimer struct {
	ID                     string
	IsEnabled              bool
	ACState                ACStateData
	CausedBy               CausedBy
	CreateTime             string
	CreateTimeSecondsAgo   int
	LastScheduledInstances []struct {
		Type                 string
		TargetTime           string
		TargetTimeSecondsAgo int
		Status               string
		ScheduleID           string
		LastExecutions       []string
	}
	TargetTime               string
	TargetTimeSecondsFromNow int
}

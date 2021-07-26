// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// DeviceSchedule holds information about schedule on the device.
type DeviceSchedule struct {
	ID        string
	IsEnabled bool
	ACState   struct {
		ACStateData
		Extra struct {
			Scheduler struct {
				ClimateReact         bool `json:"climate_react"`
				Motion               string
				On                   bool
				ClimateReactSettings ClimateReact `json:"climate_react_settings"`
				PureBoost            string       `json:"pure_boost"`
			}
		}
		HorizontalSwing string
		Light           string
	}
	CausedBy               CausedBy
	CreateTime             string
	CreateTimeSecondsAgo   int
	RecurringDays          []string
	TargetTimeLocal        string
	TimeZone               string
	PodUID                 string
	NextTime               string
	NextTimeSecondsFromNow int
}

// CreateDeviceSchedulePayload is the payload type for CreateDeviceSchedule method.
type CreateDeviceSchedulePayload struct {
	TargetTimeLocal string
	TimeZone        string
	ACState         ACStateData
	RecurringDays   []string
}

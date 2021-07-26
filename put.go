// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sensibo

import (
	"bytes"
	"fmt"

	"github.com/odinn1984/go-sensibo/models"
)

// Set the device timer.
//
// This function allows us to set the device time which will tell our device
// to set the AC state to the value of DeviceTimer.ACState.On
//
// id is the ID of the device and state is of type #models.DeviceTimer
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) SetDeviceTimer(id string, state models.DeviceTimer) (string, error) {
	payload := fmt.Sprintf(
		`
			{
				"minutesFromNow": %d,
				"acState": {
					"on": %v,
					"mode": "%s",
					"fanLevel": "%s",
					"targetTemperature": %d,
					"temperatureUnit": "%s",
					"swing": "%s"
				}
			}
		`,
		state.TargetTimeSecondsFromNow/60,
		state.ACState.On,
		state.ACState.Mode,
		state.ACState.FanLevel,
		int64(state.ACState.TargetTemperature),
		state.ACState.TemperatureUnit,
		state.ACState.Swing,
	)

	resp, err := s.makePutRequest(
		"v1",
		fmt.Sprintf("pods/%s/timer", id),
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting timer: \n\t%v", err)
	}

	return resp, nil
}

// Toggle the device climate react state on or off.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) ToggleDeviceClimateReact(id string, enabled bool) (string, error) {
	payload := fmt.Sprintf(
		`
			{
				"enabled": %v
			}
		`,
		enabled,
	)

	resp, err := s.makePutRequest(
		"v2",
		fmt.Sprintf("pods/%s/smartmode", id),
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting climate react: \n\t%v", err)
	}

	return resp, nil
}

// Toggle a device schedule state on or off.
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) ToggleDeviceSchedule(deviceID string, scheduleID string, enabled bool) (string, error) {
	payload := fmt.Sprintf(
		`
			{
				"isEnabled": %v
			}
		`,
		enabled,
	)

	resp, err := s.makePutRequest(
		"v1",
		fmt.Sprintf("pods/%s/schedules/%s", deviceID, scheduleID),
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting climate react: \n\t%v", err)
	}

	return resp, nil
}

// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sensibo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/odinn1984/go-sensibo/models"
)

// SetDeviceTimerPayload is the payload for SetDeviceTimer API call
type SetDeviceTimerPayload struct {
	MinutesFromNow int                `json:"minutesFromNow"`
	ACState        models.ACStateData `json:"acState"`
}

// SetDeviceTimer sets the device timer.
//
// This function allows us to set the device time which will tell our device
// to set the AC state to the value of DeviceTimer.ACState.On
//
// id is the ID of the device and state is of type #models.DeviceTimer
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) SetDeviceTimer(ctx context.Context, id string, minutesFromNow int, state models.ACStateData) (string, error) {
	payload := SetDeviceTimerPayload{minutesFromNow, state}

	payloadStr, err := json.Marshal(payload)

	if err != nil {
		return "", fmt.Errorf("failed marshal on payload: \n\t%v", err)
	}

	resp, err := s.makePutRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/timer", id),
		bytes.NewBuffer(payloadStr),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting timer: \n\t%v", err)
	}

	return resp, nil
}

// ToggleDeviceClimateReactPayload is the payload for ToggleDeviceClimateReact API call
type ToggleDeviceClimateReactPayload struct {
	Enabled bool `json:"enabled"`
}

// ToggleDeviceClimateReact toggles the device climate react state on or off.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) ToggleDeviceClimateReact(ctx context.Context, id string, enabled bool) (string, error) {
	payload := ToggleDeviceClimateReactPayload{enabled}
	payloadStr, err := json.Marshal(payload)

	if err != nil {
		return "", fmt.Errorf("failed marshal on payload: \n\t%v", err)
	}

	resp, err := s.makePutRequest(
		ctx,
		"v2",
		fmt.Sprintf("pods/%s/smartmode", id),
		bytes.NewBuffer(payloadStr),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting climate react: \n\t%v", err)
	}

	return resp, nil
}

// ToggleDeviceSchedulePayload is the payload for the ToggleDeviceSchedule API call
type ToggleDeviceSchedulePayload struct {
	IsEnabled bool `json:"isEnabled"`
}

// ToggleDeviceSchedule toggles a device schedule state on or off.
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) ToggleDeviceSchedule(ctx context.Context, deviceID string, scheduleID string, enabled bool) (string, error) {
	payload := ToggleDeviceSchedulePayload{enabled}
	payloadStr, err := json.Marshal(payload)

	if err != nil {
		return "", fmt.Errorf("failed marshal on payload: \n\t%v", err)
	}

	resp, err := s.makePutRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/schedules/%s", deviceID, scheduleID),
		bytes.NewBuffer(payloadStr),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting climate react: \n\t%v", err)
	}

	return resp, nil
}

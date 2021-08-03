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

// SetDeviceACStatePayload is the payload for the SetDeviceACState API
type SetDeviceACStatePayload struct {
	ACState models.ACStateData `json:"acState"`
}

// SetDeviceACState sets the AC state of the device.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) SetDeviceACState(ctx context.Context, id string, state models.ACStateData) (string, error) {
	payload := SetDeviceACStatePayload{state}
	payloadStr, err := json.Marshal(payload)

	if err != nil {
		return "", fmt.Errorf("failed marshal on payload: \n\t%v", err)
	}

	resp, err := s.makePostRequest(
		ctx,
		"v2",
		fmt.Sprintf("pods/%s/acStates", id),
		bytes.NewBuffer(payloadStr),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting device ac state: \n\t%v", err)
	}

	return resp, nil
}

// CreateDeviceSchedulePayload is the payload for the CreateDeviceSchedule API
type CreateDeviceSchedulePayload struct {
	TargetTimeLocal string             `json:"targetTimeLocal"`
	TimeZone        string             `json:"timezone"`
	ACState         models.ACStateData `json:"acState"`
	RecurringDays   []string           `json:"recurOnDaysOfWeek"`
}

// CreateDeviceSchedule creates a new schedule.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) CreateDeviceSchedule(ctx context.Context, id string, schedule CreateDeviceSchedulePayload) (string, error) {
	payload := CreateDeviceSchedulePayload{
		schedule.TargetTimeLocal,
		schedule.TimeZone,
		schedule.ACState,
		schedule.RecurringDays,
	}
	payloadStr, err := json.Marshal(payload)

	if err != nil {
		return "", fmt.Errorf("failed marshal on payload: \n\t%v", err)
	}

	resp, err := s.makePostRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/schedules", id),
		bytes.NewBuffer(payloadStr),
	)

	if err != nil {
		return "", fmt.Errorf("failed creating a schedule: \n\t%v", err)
	}

	return resp, nil
}

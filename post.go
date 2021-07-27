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

// SetDeviceACState sets the AC state of the device.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) SetDeviceACState(ctx context.Context, id string, state models.ACStateData) (string, error) {
	payload := fmt.Sprintf(
		`
			{
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
		state.On,
		state.Mode,
		state.FanLevel,
		int64(state.TargetTemperature),
		state.TemperatureUnit,
		state.Swing,
	)

	resp, err := s.makePostRequest(
		ctx,
		"v2",
		fmt.Sprintf("pods/%s/acStates", id),
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting device ac state: \n\t%v", err)
	}

	return resp, nil
}

// CreateDeviceSchedule creates a new schedule.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) CreateDeviceSchedule(ctx context.Context, id string, schedule models.CreateDeviceSchedulePayload) (string, error) {
	recurringDaysJSONArr, _ := json.Marshal(schedule.RecurringDays)
	payload := fmt.Sprintf(
		`
			{
				"targetTimeLocal": "%s",
				"timezone": "%s",
				"acState": {
					"on": %v,
					"mode": "%s",
					"fanLevel": "%s",
					"targetTemperature": %d,
					"temperatureUnit": "%s",
					"swing": "%s"
				},
				"recurOnDaysOfWeek": %v
			}
		`,
		schedule.TargetTimeLocal,
		schedule.TimeZone,
		schedule.ACState.On,
		schedule.ACState.Mode,
		schedule.ACState.FanLevel,
		int64(schedule.ACState.TargetTemperature),
		schedule.ACState.TemperatureUnit,
		schedule.ACState.Swing,
		string(recurringDaysJSONArr),
	)

	resp, err := s.makePostRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/schedules", id),
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		return "", fmt.Errorf("failed creating a schedule: \n\t%v", err)
	}

	return resp, nil
}

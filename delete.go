// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sensibo

import (
	"context"
	"fmt"
)

// DeleteDeviceTimer deletes the timer on the device.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) DeleteDeviceTimer(ctx context.Context, id string) (string, error) {
	resp, err := s.makeDeleteRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/timer", id),
	)

	if err != nil {
		return "", fmt.Errorf("failed deleting timer: \n\t%v", err)
	}

	return resp, nil
}

// DeleteDeviceSchedule deletes a schedule.
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) DeleteDeviceSchedule(ctx context.Context, deviceID string, scheduleID string) (string, error) {
	resp, err := s.makeDeleteRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/schedules/%s", deviceID, scheduleID),
	)

	if err != nil {
		return "", fmt.Errorf("failed deleting schedule: \n\t%v", err)
	}

	return resp, nil
}

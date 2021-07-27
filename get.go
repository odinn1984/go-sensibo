// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sensibo

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/odinn1984/go-sensibo/models"
)

// GetAllDevices gets all of the devices you have access to.
//
// fields is a filter on which fields you will have values for
// in the response.
//
// e.g: To get all fields use "*" and to get "id" only use "id"
func (s *Sensibo) GetAllDevices(ctx context.Context, fields []string) ([]models.Device, error) {
	resp, err := s.makeGetRequest(
		ctx,
		"v2",
		"users/me/pods",
		map[string]string{"fields": strings.Join(fields, ",")},
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting all devices \n\t%v", err)
	}

	parsedResp := struct {
		Status string
		Result []models.Device
	}{}

	if err := json.Unmarshal([]byte(resp), &parsedResp); err != nil {
		return nil, fmt.Errorf("failed parsing result \n\t%v", err)
	}

	return parsedResp.Result, nil
}

// GetDevice gets a device by ID.
//
// id is the ID of the device
//
// fields is a filter on which fields you will have values for
// in the response.
//
// e.g: To get all fields use "*" and to get "id" only use "id"
func (s *Sensibo) GetDevice(ctx context.Context, id string, fields []string) (*models.Device, error) {
	resp, err := s.makeGetRequest(
		ctx,
		"v2",
		fmt.Sprintf("pods/%s", id),
		map[string]string{"fields": strings.Join(fields, ",")},
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting device \n\t%v", err)
	}

	parsedResp := struct {
		Status string
		Result models.Device
	}{}

	if err := json.Unmarshal([]byte(resp), &parsedResp); err != nil {
		return nil, fmt.Errorf("failed parsing result \n\t%v", err)
	}

	return &parsedResp.Result, nil
}

// GetDeviceACStates gets a device's AC stats by device ID.
//
// id is the ID of the device
// limit the amount of entries you get in the response
func (s *Sensibo) GetDeviceACStates(ctx context.Context, id string, limit uint) ([]models.ACState, error) {
	resp, err := s.makeGetRequest(
		ctx,
		"v2",
		fmt.Sprintf("pods/%s/acStates", id),
		map[string]string{"limit": fmt.Sprintf("%d", limit)},
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting AC State \n\t%v", err)
	}

	parsedResp := struct {
		Status string
		Result []models.ACState
	}{}

	if err := json.Unmarshal([]byte(resp), &parsedResp); err != nil {
		return nil, fmt.Errorf("failed parsing result \n\t%v", err)
	}

	return parsedResp.Result, nil
}

// GetDeviceHistoricalMeasurements gets historical measurements for a device.
//
// id is the ID of the device
// days is the number of days we want to get the data for
func (s *Sensibo) GetDeviceHistoricalMeasurements(ctx context.Context, id string, days uint) (*models.HistoricalMeasurements, error) {
	resp, err := s.makeGetRequest(
		ctx,
		"v2",
		fmt.Sprintf("pods/%s/historicalMeasurements", id),
		map[string]string{"days": fmt.Sprintf("%d", days)},
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting historical measurements \n\t%v", err)
	}

	parsedResp := struct {
		Status string
		Result models.HistoricalMeasurements
	}{}

	if err := json.Unmarshal([]byte(resp), &parsedResp); err != nil {
		return nil, fmt.Errorf("failed parsing result \n\t%v", err)
	}

	return &parsedResp.Result, nil
}

// GetDeviceClimateReactSettings gets climate react settings for a device.
//
// id is the ID of the device
func (s *Sensibo) GetDeviceClimateReactSettings(ctx context.Context, id string) (*models.ClimateReact, error) {
	resp, err := s.makeGetRequest(
		ctx,
		"v2",
		fmt.Sprintf("pods/%s/smartmode", id),
		map[string]string{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting climate react settings \n\t%v", err)
	}

	parsedResp := struct {
		Status string
		Result models.ClimateReact
	}{}

	if err := json.Unmarshal([]byte(resp), &parsedResp); err != nil {
		return nil, fmt.Errorf("failed parsing result \n\t%v", err)
	}

	return &parsedResp.Result, nil
}

// GetDeviceTimer gets the timer for a device.
//
// id is the ID of the device
func (s *Sensibo) GetDeviceTimer(ctx context.Context, id string) (*models.DeviceTimer, error) {
	resp, err := s.makeGetRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/timer", id),
		map[string]string{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting timer \n\t%v", err)
	}

	parsedResp := struct {
		Status string
		Result models.DeviceTimer
	}{}

	if err := json.Unmarshal([]byte(resp), &parsedResp); err != nil {
		return nil, fmt.Errorf("failed parsing result \n\t%v", err)
	}

	return &parsedResp.Result, nil
}

// GetDeviceSchedules gets all the schedules set on the device.
//
// id is the ID of the device
func (s *Sensibo) GetDeviceSchedules(ctx context.Context, id string) ([]models.DeviceSchedule, error) {
	resp, err := s.makeGetRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/schedules", id),
		map[string]string{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting schedules \n\t%v", err)
	}

	parsedResp := struct {
		Status string
		Result []models.DeviceSchedule
	}{}

	if err := json.Unmarshal([]byte(resp), &parsedResp); err != nil {
		return nil, fmt.Errorf("failed parsing result \n\t%v", err)
	}

	return parsedResp.Result, nil
}

// GetDeviceSchedule gets a schedule by ID on the device.
func (s *Sensibo) GetDeviceSchedule(ctx context.Context, deviceID string, scheduleID string) (*models.DeviceSchedule, error) {
	resp, err := s.makeGetRequest(
		ctx,
		"v1",
		fmt.Sprintf("pods/%s/schedules/%s", deviceID, scheduleID),
		map[string]string{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting schedule \n\t%v", err)
	}

	parsedResp := struct {
		Status string
		Result models.DeviceSchedule
	}{}

	if err := json.Unmarshal([]byte(resp), &parsedResp); err != nil {
		return nil, fmt.Errorf("failed parsing result \n\t%v", err)
	}

	return &parsedResp.Result, nil
}

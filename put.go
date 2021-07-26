package sensibo

import (
	"bytes"
	"fmt"

	"github.com/odinn1984/go-sensibo/models"
)

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

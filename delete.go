package sensibo

import (
	"fmt"
)

func (s *Sensibo) DeleteDeviceTimer(id string) (string, error) {
	resp, err := s.makeDeleteRequest(
		"v1",
		fmt.Sprintf("pods/%s/time", id),
	)

	if err != nil {
		return "", fmt.Errorf("failed deleting timer: \n\t%v", err)
	}

	return resp, nil
}

func (s *Sensibo) DeleteDeviceSchedule(deviceID string, scheduleID string) (string, error) {
	resp, err := s.makeDeleteRequest(
		"v1",
		fmt.Sprintf("pods/%s/schedules/%s", deviceID, scheduleID),
	)

	if err != nil {
		return "", fmt.Errorf("failed deleting schedule: \n\t%v", err)
	}

	return resp, nil
}

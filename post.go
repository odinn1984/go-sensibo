package sensibo

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/odinn1984/go-sensibo/models"
)

func (s *Sensibo) SetDeviceACState(id string, state models.ACStateData) (string, error) {
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
		"v2",
		fmt.Sprintf("pods/%s/acStates", id),
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		return "", fmt.Errorf("failed setting device ac state: \n\t%v", err)
	}

	return resp, nil
}

func (s *Sensibo) CreateDeviceSchedule(id string, schedule models.CreateDeviceSchedulePayload) (string, error) {
	recurringDaysJsonArr, _ := json.Marshal(schedule.RecurringDays)
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
		string(recurringDaysJsonArr),
	)

	resp, err := s.makePostRequest(
		"v1",
		fmt.Sprintf("pods/%s/schedules", id),
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		return "", fmt.Errorf("failed creating a schedule: \n\t%v", err)
	}

	return resp, nil
}

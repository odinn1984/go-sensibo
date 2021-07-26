package sensibo

import (
	"bytes"
	"fmt"
)

func (s *Sensibo) SetDeviceACStateProperty(id string, property string, value string) (string, error) {
	payload := fmt.Sprintf(
		`
			{
				"newValue": %s
			}
		`,
		value,
	)

	resp, err := s.makePatchRequest(
		"v1",
		fmt.Sprintf("pods/%s/acStates/%s", id, property),
		bytes.NewBuffer([]byte(payload)),
	)

	if err != nil {
		return "", fmt.Errorf("failed updating property: \n\t%v", err)
	}

	return resp, nil
}

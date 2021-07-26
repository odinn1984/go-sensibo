// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sensibo

import (
	"bytes"
	"fmt"
)

// Update a property in the AC state.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
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

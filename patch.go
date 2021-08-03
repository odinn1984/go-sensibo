// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sensibo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// SetDeviceACStatePropertyPayload is the payload for SetDeviceACStateProperty API call
type SetDeviceACStatePropertyPayload struct {
	NewValue string `json:"newValue"`
}

// SetDeviceACStateProperty updates a property in the AC state.
//
// id is the ID of the device
//
// It returns the direct response from Sensibo API as a string or error
// if an issue occurred
func (s *Sensibo) SetDeviceACStateProperty(ctx context.Context, id string, property string, value string) (string, error) {
	payload := SetDeviceACStatePropertyPayload{
		NewValue: value,
	}

	payloadStr, err := json.Marshal(payload)

	if err != nil {
		return "", fmt.Errorf("failed marshal on payload: \n\t%v", err)
	}

	resp, err := s.makePatchRequest(
		ctx,
		"v2",
		fmt.Sprintf("pods/%s/acStates/%s", id, property),
		bytes.NewBuffer(payloadStr),
	)

	if err != nil {
		return "", fmt.Errorf("failed updating property: \n\t%v", err)
	}

	return resp, nil
}

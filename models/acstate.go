// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// ACState holds extended information about the state of the AC.
type ACState struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	ACState struct {
		Timestamp SensiboTime `json:"timestamp"`
		ACStateData
	} `json:"acState"`
	ChangedProperties []string `json:"changedProperties"`
	Reason            string   `json:"reason"`
	FailureReason     string   `json:"failureReason"`
}

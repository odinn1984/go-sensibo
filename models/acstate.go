// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// ACState holds extended information about the state of the AC.
type ACState struct {
	ID      string
	Status  string
	ACState struct {
		Timestamp SensiboTime
		ACStateData
	}
	ChangedProperties []string
	Reason            string
	FailureReason     string
}

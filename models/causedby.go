// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package models

// Holds information about who performed the action.
type CausedBy struct {
	Username  string
	Email     string
	FirstName string
	LastName  string
}

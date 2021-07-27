// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*Package sensibo is a Go wrapper for the Sensibo API.

This implements the API calls that are listed in https://sensibo.github.io/
in a simple to use Go package that can be imported to any project.

You can run the following simple code to get you started:

	package main

	import "github.com/odinn1984/go-sensibo"

	func main() {
		client := sensibo.New("my-api-key")
		devices, err := client.GetAllDevices([]string{"*"})

		if err != nil {
			// Do some error handling
		}

		// Do something with devices
	}

The above example get all of the devices that are configured on your Sensibo account. */
package sensibo

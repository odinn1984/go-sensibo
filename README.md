# go-sensibo

[![Lint And Test](https://github.com/odinn1984/go-sensibo/actions/workflows/ci.yml/badge.svg)](https://github.com/odinn1984/go-sensibo/actions/workflows/ci.yml)
[![Known Vulnerabilities](https://snyk.io/test/github/odinn1984/go-sensibo/badge.svg)](https://snyk.io/test/github/odinn1984/go-sensibo)
[![Go Report Card](https://goreportcard.com/badge/github.com/odinn1984/go-sensibo)](https://goreportcard.com/report/github.com/odinn1984/go-sensibo)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.14-61CFDD.svg)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/odinn1984/go-sensibo)](https://pkg.go.dev/mod/github.com/odinn1984/go-sensibo)

This is a Go wrapper for the [Sensibo](https://sensibo.com) API.

This implements the API calls that are listed in [the documentation here](https://sensibo.github.io/) in a simple to use Go package that can be imported to any project.

## Install

To add go-sensibo to your project simply run:

```shell
go get github.com/odinn1984/go-sensibo
```

## How To Use?

To use `go-sensibo` simply add `require github.com/odinn1984/go-sinsibo v0.1.0` to your `go.mod` file or just add `github.com/odinn1984/go-sensibo` to an already existing `require` block.

Then you can run the following simple code to get you started:

```go
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
```

The above example get all of the devices that are configure on your Sensibo account.
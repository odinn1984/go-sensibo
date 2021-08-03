// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sensibo

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/odinn1984/go-sensibo/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSensibo_DeleteDeviceTimer(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name   string
		args   args
		DoMock func(req *http.Request) (*http.Response, error)
		want   string
		err    error
	}{
		{
			name: "returns api request response on success",
			args: args{
				ctx: context.Background(),
				id:  "1234",
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader("Success")),
				}, nil
			},
			want: "Success",
			err:  nil,
		},
		{
			name: "returns err on error",
			args: args{
				ctx: context.Background(),
				id:  "1234",
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader("Success")),
				}, fmt.Errorf("Error")
			},
			want: "",
			err:  fmt.Errorf("failed deleting timer: \n\tfailed making request \n\tCode: 200 \n\tMsg: Success \n\tErr: Error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(
				&mocks.HTTPClientMock{
					DoMock: tt.DoMock,
				},
				"api-key",
			)

			got, err := s.DeleteDeviceTimer(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestSensibo_DeleteDeviceSchedule(t *testing.T) {
	type args struct {
		ctx        context.Context
		deviceID   string
		scheduleID string
	}
	tests := []struct {
		name   string
		args   args
		DoMock func(req *http.Request) (*http.Response, error)
		want   string
		err    error
	}{
		{
			name: "returns api request response on success",
			args: args{
				ctx:        context.Background(),
				deviceID:   "1234",
				scheduleID: "1234",
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader("Success")),
				}, nil
			},
			want: "Success",
			err:  nil,
		},
		{
			name: "returns err on error",
			args: args{
				ctx:        context.Background(),
				deviceID:   "1234",
				scheduleID: "1234",
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader("Success")),
				}, fmt.Errorf("Error")
			},
			want: "",
			err:  fmt.Errorf("failed deleting schedule: \n\tfailed making request \n\tCode: 200 \n\tMsg: Success \n\tErr: Error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(
				&mocks.HTTPClientMock{
					DoMock: tt.DoMock,
				},
				"api-key",
			)

			got, err := s.DeleteDeviceSchedule(tt.args.ctx, tt.args.deviceID, tt.args.scheduleID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

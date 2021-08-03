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

func TestSensibo_SetDeviceACStateProperty(t *testing.T) {
	type args struct {
		ctx      context.Context
		id       string
		property string
		value    string
	}
	tests := []struct {
		name   string
		args   args
		DoMock func(req *http.Request) (*http.Response, error)
		want   string
		err    error
	}{
		{
			name: "successful execution",
			args: args{
				ctx:      context.Background(),
				id:       "1234",
				property: "prop",
				value:    "value",
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader("Success")),
					StatusCode: 200,
				}, nil
			},
			want: "Success",
			err:  nil,
		},
		{
			name: "returns an error or request failure: status code not 200",
			args: args{
				ctx:      context.Background(),
				id:       "1234",
				property: "prop",
				value:    "value",
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader("")),
					StatusCode: 301,
				}, nil
			},
			want: "",
			err:  fmt.Errorf("failed updating property: \n\tfailed making request \n\tCode: 301 \n\tMsg:  \n\tErr: <nil>"),
		},
		{
			name: "returns an error or request failure: returned err not nil",
			args: args{
				ctx:      context.Background(),
				id:       "1234",
				property: "prop",
				value:    "value",
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader("")),
					StatusCode: 301,
				}, fmt.Errorf("Error")
			},
			want: "",
			err:  fmt.Errorf("failed updating property: \n\tfailed making request \n\tCode: 301 \n\tMsg:  \n\tErr: Error"),
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

			got, err := s.SetDeviceACStateProperty(tt.args.ctx, tt.args.id, tt.args.property, tt.args.value)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Copyright 2021 To Levan Giguashvili. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sensibo

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/odinn1984/go-sensibo/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name       string
		apikey     string
		httpClient HTTPClient
		want       *Sensibo
	}{
		{
			name:       "sanity check",
			apikey:     "api-key",
			httpClient: http.DefaultClient,
			want: &Sensibo{
				APIKey:     "api-key",
				httpClient: http.DefaultClient,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(&http.Client{}, tt.apikey)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSensibo_getRequestURL(t *testing.T) {
	type args struct {
		version  string
		endpoint string
		params   map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get correct url with api key appended",
			args: args{
				version:  "v1",
				endpoint: "end/point",
				params:   map[string]string{},
			},
			want: "https://home.sensibo.com/api/v1/end/point?apiKey=api-key",
		},
		{
			name: "appends query params correctly",
			args: args{
				version:  "v1",
				endpoint: "end/point",
				params: map[string]string{
					"a": "b",
					"c": "d",
				},
			},
			want: "https://home.sensibo.com/api/v1/end/point?apiKey=api-key&a=b&c=d",
		},
		{
			name: "url escape query params",
			args: args{
				version:  "v1",
				endpoint: "end/point",
				params: map[string]string{
					"a": "b",
					"c": "d==#",
				},
			},
			want: "https://home.sensibo.com/api/v1/end/point?apiKey=api-key&a=b&c=d%3D%3D%23",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sensibo{
				APIKey: "api-key",
			}

			got := s.getRequestURL(tt.args.version, tt.args.endpoint, tt.args.params)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSensibo_makeRequest(t *testing.T) {
	type args struct {
		ctx    context.Context
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name   string
		args   args
		DoMock func(req *http.Request) (*http.Response, error)
		want   string
		err    error
	}{
		{
			name: "fails on bad method",
			args: args{
				ctx:    context.Background(),
				method: "bad method",
				url:    "https://sensibo.com",
				body:   nil,
			},
			DoMock: http.DefaultClient.Do,
			want:   "",
			err:    fmt.Errorf("unable to create new request: \n\tnet/http: invalid method \"bad method\""),
		},
		{
			name: "fails if http.DefaultClient.Do returns an error",
			args: args{
				ctx:    context.Background(),
				method: http.MethodGet,
				url:    "https://sensibo.com",
				body:   nil,
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
						Body:       ioutil.NopCloser(strings.NewReader("")),
						StatusCode: 200,
					},
					fmt.Errorf("Error")
			},
			want: "",
			err:  fmt.Errorf("failed making request \n\tCode: 200 \n\tMsg:  \n\tErr: Error"),
		},
		{
			name: "fails if response has non 200 status code",
			args: args{
				ctx:    context.Background(),
				method: http.MethodGet,
				url:    "https://sensibo.com",
				body:   nil,
			},
			DoMock: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader("")),
					StatusCode: 301,
				}, nil
			},
			want: "",
			err:  fmt.Errorf("failed making request \n\tCode: 301 \n\tMsg:  \n\tErr: <nil>"),
		},
		{
			name: "successful request",
			args: args{
				ctx:    context.Background(),
				method: http.MethodGet,
				url:    "https://sensibo.com",
				body:   nil,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(
				&mocks.HTTPClientMock{
					DoMock: tt.DoMock,
				},
				"api-key",
			)

			got, err := s.makeRequest(
				tt.args.ctx,
				tt.args.method,
				tt.args.url,
				tt.args.body,
			)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

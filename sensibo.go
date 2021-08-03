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
	"net/url"
	"sort"
)

// Sensibo holds all of the available functions to interact with the Sensibo API.
type Sensibo struct {
	APIKey     string
	httpClient HTTPClient
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New creates new Sensibo instance.
//
// apiKey is the API key that you got from https://home.sensibo.com/me/api
// To generate an API key just go to https://home.sensibo.com/me/api
// and click on "Add API Key" buttone, fill in the name and it will create the key
//
// httpClient is the client we want to use for http requests (e.g: http.DefaultClient)
//
// It returns a pointed to Sensibo with the key already stored in it
func New(httpClient HTTPClient, apikey string) *Sensibo {
	return &Sensibo{
		APIKey:     apikey,
		httpClient: httpClient,
	}
}

func (s *Sensibo) getRequestURL(
	version string,
	endpoint string,
	params map[string]string,
) string {
	baseURL := fmt.Sprintf("https://home.sensibo.com/api/%s/%s?apiKey=%s", version, endpoint, s.APIKey)
	queryKeys := []string{}
	queryParams := ""

	for k := range params {
		queryKeys = append(queryKeys, k)
	}

	sort.Strings(queryKeys)

	for _, k := range queryKeys {
		queryParams = fmt.Sprintf("%s&%s=%s", queryParams, k, url.QueryEscape(params[k]))
	}

	return fmt.Sprintf("%s%s", baseURL, queryParams)
}

func (s *Sensibo) makeRequest(ctx context.Context, method string, url string, body io.Reader) (string, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)

	if err != nil {
		return "", fmt.Errorf("unable to create new request: \n\t%v", err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("accept", "*/*")

	res, err := s.httpClient.Do(req)

	if err != nil || res.StatusCode != http.StatusOK {
		resBytes, ioErr := ioutil.ReadAll(res.Body)

		if ioErr != nil {
			return "", fmt.Errorf("io error occurred: \n\t%v", ioErr)
		}

		defer res.Body.Close()

		return "", fmt.Errorf(
			"failed making request \n\tCode: %v \n\tMsg: %v \n\tErr: %v",
			res.StatusCode,
			string(resBytes),
			err,
		)
	}

	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("failed to read response: \n\t%v", err)
	}

	return string(resBytes), nil
}

func (s *Sensibo) makeGetRequest(
	ctx context.Context,
	version string,
	endpoint string,
	params map[string]string,
) (string, error) {
	return s.makeRequest(ctx,
		http.MethodGet,
		s.getRequestURL(version, endpoint, params),
		nil,
	)
}

func (s *Sensibo) makePutRequest(
	ctx context.Context,
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	return s.makeRequest(
		ctx,
		http.MethodPut,
		s.getRequestURL(version, endpoint, map[string]string{}),
		body,
	)
}

func (s *Sensibo) makePatchRequest(
	ctx context.Context,
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	return s.makeRequest(
		ctx,
		http.MethodPatch,
		s.getRequestURL(version, endpoint, map[string]string{}),
		body,
	)
}

func (s *Sensibo) makePostRequest(
	ctx context.Context,
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	return s.makeRequest(
		ctx,
		http.MethodPost,
		s.getRequestURL(version, endpoint, map[string]string{}),
		body,
	)
}

func (s *Sensibo) makeDeleteRequest(
	ctx context.Context,
	version string,
	endpoint string,
) (string, error) {
	return s.makeRequest(
		ctx,
		http.MethodDelete,
		s.getRequestURL(version, endpoint, map[string]string{}),
		nil,
	)
}

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
)

// Sensibo holds all of the available functions to interact with the Sensibo API.
type Sensibo struct {
	APIKey string
}

// New creates new Sensibo instance.
//
// apiKey is the API key that you got from https://home.sensibo.com/me/api
// To generate an API key just go to https://home.sensibo.com/me/api
// and click on "Add API Key" buttone, fill in the name and it will create the key
//
// It returns a pointed to Sensibo with the key already stored in it
func New(apikey string) *Sensibo {
	return &Sensibo{
		APIKey: apikey,
	}
}

func (s *Sensibo) getRequestURL(
	ctx context.Context,
	version string,
	endpoint string,
	params map[string]string,
) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf("https://home.sensibo.com/api/%v/%v", version, endpoint),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("unable to create new request: \n\t%v", err)
	}

	query := req.URL.Query()

	query.Add("apiKey", s.APIKey)

	for k, v := range params {
		query.Add(k, v)
	}

	req.URL.RawQuery = query.Encode()

	return req.URL.String(), nil
}

func (s *Sensibo) makeRequest(ctx context.Context, method string, url string, body io.Reader) (string, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)

	if err != nil {
		return "", fmt.Errorf("unable to create new request: \n\t%v", err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("accept", "*/*")

	res, err := http.DefaultClient.Do(req)

	if err != nil || res.StatusCode != 200 {
		resBytes, err := ioutil.ReadAll(res.Body)
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
	url, err := s.getRequestURL(ctx, version, endpoint, params)

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest(ctx, "GET", url, nil)
}

func (s *Sensibo) makePutRequest(
	ctx context.Context,
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	url, err := s.getRequestURL(ctx, version, endpoint, map[string]string{})

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest(ctx, "PUT", url, body)
}

func (s *Sensibo) makePatchRequest(
	ctx context.Context,
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	url, err := s.getRequestURL(ctx, version, endpoint, map[string]string{})

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest(ctx, "PATCH", url, body)
}

func (s *Sensibo) makePostRequest(
	ctx context.Context,
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	url, err := s.getRequestURL(ctx, version, endpoint, map[string]string{})

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest(ctx, "POST", url, body)
}

func (s *Sensibo) makeDeleteRequest(
	ctx context.Context,
	version string,
	endpoint string,
) (string, error) {
	url, err := s.getRequestURL(ctx, version, endpoint, map[string]string{})

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest(ctx, "DELETE", url, nil)
}

package sensibo

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Sensibo struct {
	ApiKey string
}

func New(apikey string) *Sensibo {
	return &Sensibo{
		ApiKey: apikey,
	}
}

func (s *Sensibo) getRequestURL(version string, endpoint string, params map[string]string) (string, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://home.sensibo.com/api/%v/%v", version, endpoint),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("unable to create new request: \n\t%v", err)
	}

	query := req.URL.Query()

	query.Add("apiKey", s.ApiKey)

	for k, v := range params {
		query.Add(k, v)
	}

	req.URL.RawQuery = query.Encode()

	return req.URL.String(), nil
}

func (s *Sensibo) makeRequest(method string, url string, body io.Reader) (string, error) {
	req, err := http.NewRequest(method, url, body)

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
	version string,
	endpoint string,
	params map[string]string,
) (string, error) {
	url, err := s.getRequestURL(version, endpoint, params)

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest("GET", url, nil)
}

func (s *Sensibo) makePutRequest(
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	url, err := s.getRequestURL(version, endpoint, map[string]string{})

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest("PUT", url, body)
}

func (s *Sensibo) makePatchRequest(
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	url, err := s.getRequestURL(version, endpoint, map[string]string{})

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest("PATCH", url, body)
}

func (s *Sensibo) makePostRequest(
	version string,
	endpoint string,
	body io.Reader,
) (string, error) {
	url, err := s.getRequestURL(version, endpoint, map[string]string{})

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest("POST", url, body)
}

func (s *Sensibo) makeDeleteRequest(
	version string,
	endpoint string,
) (string, error) {
	url, err := s.getRequestURL(version, endpoint, map[string]string{})

	if err != nil {
		return "", fmt.Errorf("failed getting request url: \n\t%v", err)
	}

	return s.makeRequest("DELETE", url, nil)
}

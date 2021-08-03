package mocks

import "net/http"

type HTTPClientMock struct {
	DoMock func(*http.Request) (*http.Response, error)
}

func (h *HTTPClientMock) Do(req *http.Request) (*http.Response, error) {
	return h.DoMock(req)
}

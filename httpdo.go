package HttpDo

import (
	"io"
	"net/http"
)

type HttpDo struct {
	method       string
	requestBody  io.Reader
	url          string
	httpErr      error
	request      *http.Request
	client       *http.Client
	StatusCode   int
	ResponseBody io.ReadCloser
}

func New() *HttpDo {
	httpRequest := new(HttpDo)
	httpRequest.client = http.DefaultClient
	return httpRequest
}

func (h *HttpDo) Run() *HttpDo {
	if h.httpErr != nil {
		return h
	}
	req, err := http.NewRequest(h.method, h.url, h.requestBody)
	if err != nil {
		h.httpErr = err
		return h
	}
	h.request = req
	res, err := h.client.Do(h.request)
	if err != nil {
		h.httpErr = err
		return h
	}
	h.StatusCode = res.StatusCode
	h.ResponseBody = res.Body
	return h
}

func (h *HttpDo) Err() error {
	return h.httpErr
}

func (h *HttpDo) Url(url string) *HttpDo {
	if h.httpErr != nil {
		return h
	}
	h.url = url
	return h
}

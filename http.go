package http

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

type GetRequest struct {
	RequestUrl         string
	Header             map[string]string
	InsecureSkipVerify bool
	TimeOut            int64
}

type GetResponse struct {
	HttpCode int
	Header   http.Header
	Content  string
	Error    error
}

func (request *GetRequest) Get() *GetResponse {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: request.InsecureSkipVerify},
	}
	if request.TimeOut <= 0 {
		request.TimeOut = 2
	}
	if request.Header == nil {
		request.Header = make(map[string]string)
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * time.Duration(request.TimeOut),
	}

	response := new(GetResponse)
	response.HttpCode = http.StatusNotFound
	req, err := http.NewRequest("GET", request.RequestUrl, nil)
	if err != nil {
		response.Error = err
		return response
	}
	if len(request.Header) > 0 {
		for k, v := range request.Header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		response.Error = err
		return response
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	response.HttpCode = resp.StatusCode
	response.Header = resp.Header
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.Error = err
		return response
	}
	response.Content = string(body)
	return response
}

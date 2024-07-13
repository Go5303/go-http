package http

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type PostRequest struct {
	RequestUrl         string
	Header             map[string]string
	FormData           map[string]string
	BodyContent        string
	InsecureSkipVerify bool
	TimeOut            int64
}

type PostResponse struct {
	HttpCode int
	Header   http.Header
	Content  string
	Error    error
}

func (request *PostRequest) Post() *PostResponse {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: request.InsecureSkipVerify},
	}
	if request.TimeOut <= 0 {
		request.TimeOut = 2
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * time.Duration(request.TimeOut),
	}

	var req = new(http.Request)
	var err error
	if len(request.FormData) > 0 {
		paramData := url.Values{}
		for k, v := range request.FormData {
			paramData.Set(k, v)
		}
		req, err = http.NewRequest("POST", request.RequestUrl, strings.NewReader(paramData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	} else {
		bodyReader := new(bytes.Buffer)
		if len(request.BodyContent) > 0 {
			bodyReader = bytes.NewBuffer([]byte(request.BodyContent))
		}
		req, err = http.NewRequest("POST", request.RequestUrl, bodyReader)
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	response := new(PostResponse)
	response.HttpCode = http.StatusNotFound
	if err != nil {
		response.Error = err
		return response
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

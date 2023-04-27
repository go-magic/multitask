package example

import (
	"errors"
	"net/http"
)

type Http struct {
	StatusCode int
	Error      error
}

type HttpRequest struct {
	Url string
}

type HttpResponse struct {
	StatusCode int
}

func (h *Http) Check(request interface{}) (interface{}, error) {
	httpRequest, err := h.parseRequest(request)
	if err != nil {
		return nil, err
	}
	res, err := http.Get(httpRequest.Url)
	if err != nil {
		h.Error = err
		return nil, err
	}
	h.StatusCode = res.StatusCode
	defer res.Body.Close()
	return HttpResponse{StatusCode: res.StatusCode}, nil
}

func (h *Http) Parse(totalResponse *TotalResponse) func(interface{}) error {
	return func(response interface{}) error {
		httpResponse, err := h.parseResponse(response)
		if err != nil {
			return err
		}
		totalResponse.StatusCode = httpResponse.StatusCode
		return nil
	}
}

func (h *Http) parseRequest(request interface{}) (*HttpRequest, error) {
	request1, ok := request.(*HttpRequest)
	if ok {
		return request1, nil
	}
	request2, ok := request.(HttpRequest)
	if ok {
		return &request2, nil
	}
	return nil, errors.New("解析任务失败")
}

func (h *Http) parseResponse(response interface{}) (*HttpResponse, error) {
	response1, ok := response.(*HttpResponse)
	if ok {
		return response1, nil
	}
	response2, ok := response.(HttpResponse)
	if ok {
		return &response2, nil
	}
	return nil, errors.New("解析任务失败")
}

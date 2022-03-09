package example

import (
	"errors"
	"fmt"
	"github.com/go-magic/multitask/engine"
	"github.com/go-magic/multitask/task"
	"net/http"
	"testing"
)

type Http struct {
}

type HttpRequest struct {
	Url string
}

type TotalResponse struct {
	StatusCode int
	DnsTime    int
}

type HttpResponse struct {
	StatusCode int
}

func (h Http) Check(request interface{}) (interface{}, error) {
	httpRequest, err := h.parseRequest(request)
	if err != nil {
		return nil, err
	}
	res, err := http.Get(httpRequest.Url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return HttpResponse{StatusCode: res.StatusCode}, nil
}

func (h Http) Parse(totalResponse *TotalResponse) func(interface{}) error {
	return func(response interface{}) error {
		httpResponse, err := h.parseResponse(response)
		if err != nil {
			return err
		}
		totalResponse.StatusCode = httpResponse.StatusCode
		return nil
	}
}

func (h Http) parseRequest(request interface{}) (*HttpRequest, error) {
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

func (h Http) parseResponse(response interface{}) (*HttpResponse, error) {
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

type Dns struct {
}

type DnsRequest struct {
	Url string
}

type DnsResponse struct {
	DnsTime int
}

func (d Dns) Check(interface{}) (interface{}, error) {
	return DnsResponse{DnsTime: 1}, nil
}

func (d Dns) Parse(totalResponse *TotalResponse) func(interface{}) error {
	return func(response interface{}) error {
		httpResponse, err := d.parseResponse(response)
		if err != nil {
			return err
		}
		totalResponse.DnsTime = httpResponse.DnsTime
		return nil
	}
}

func (d Dns) parseResponse(response interface{}) (*DnsResponse, error) {
	response1, ok := response.(*DnsResponse)
	if ok {
		return response1, nil
	}
	response2, ok := response.(DnsResponse)
	if ok {
		return &response2, nil
	}
	return nil, errors.New("解析任务失败")
}

func TestName(t *testing.T) {
	engine.InitEngine(10)
	h := Http{}
	total := &TotalResponse{}
	httpRequest := HttpRequest{Url: "https://www.qq.com"}
	responseChan := make(chan task.Response, 2)
	engine.GetEngineInstance().AddRequest(task.Request{
		Handler:      h,
		Task:         httpRequest,
		Parser:       h.Parse(total),
		ResponseChan: responseChan,
	})
	dns := Dns{}
	dnsRequest := DnsRequest{Url: "https://www.qq.com"}
	engine.GetEngineInstance().AddRequest(task.Request{
		Handler:      dns,
		Task:         dnsRequest,
		Parser:       dns.Parse(total),
		ResponseChan: responseChan,
	})
	wait(responseChan)
	t.Log(total)
}

func wait(responseChan chan task.Response) {
	count := 0
	for {
		select {
		case response := <-responseChan:
			count++
			if err := response.Parser(response.Result); err != nil {
				fmt.Println(err)
			}
			if count == 2 {
				return
			}
		}
	}
}

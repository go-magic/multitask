package example

import (
	"errors"
	"fmt"
)

type Print struct {
}

type PrintRequest struct {
	Log string
}

func (p *Print) Check(request interface{}) (interface{}, error) {
	printRequest, err := p.parse(request)
	if err != nil {
		return nil, err
	}
	fmt.Println(printRequest.Log)
	return nil, nil
}

func (p *Print) parse(request interface{}) (*PrintRequest, error) {
	request1, ok := request.(*PrintRequest)
	if ok {
		return request1, nil
	}
	request2, ok := request.(PrintRequest)
	if ok {
		return &request2, nil
	}
	return nil, errors.New("解析任务失败")
}

func (p *Print) SaveLog(response interface{}) error {
	fmt.Println("save log success")
	return nil
}

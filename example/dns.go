package example

import "errors"

type Dns struct {
	DnsTime int
}

type DnsRequest struct {
	Url string
}

type DnsResponse struct {
	DnsTime int
}

type TotalResponse struct {
	StatusCode int
	DnsTime    int
}

func (d *Dns) Check(interface{}) (interface{}, error) {
	return &DnsResponse{DnsTime: 1}, nil
}

func (d *Dns) Parse(totalResponse *TotalResponse) func(interface{}) error {
	return func(response interface{}) error {
		httpResponse, err := d.parseResponse(response)
		if err != nil {
			return err
		}
		totalResponse.DnsTime = httpResponse.DnsTime
		return nil
	}
}

func (d *Dns) parseResponse(response interface{}) (*DnsResponse, error) {
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

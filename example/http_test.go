package example

import (
	"fmt"
	"github.com/go-magic/multitask/engine"
	"github.com/go-magic/multitask/task"
	"testing"
)

func TestName(t *testing.T) {
	engine.InitEngine(10, 20)
	h := &Http{}
	total := &TotalResponse{}
	responseChan := make(chan *task.Response, 2)
	tasks := make([]*task.Request, 0, 2)
	tasks = append(tasks, task.NewRequest(
		h,
		HttpRequest{Url: "https://www.baidu.com"},
		responseChan,
		h.Parse(total)))
	dns := &Dns{}
	tasks = append(tasks, task.NewRequest(
		dns,
		DnsRequest{Url: "https://www.qq.com"},
		responseChan,
		dns.Parse(total)))

	pt := &Print{}
	tasks = append(tasks, task.NewRequest(
		pt,
		PrintRequest{Log: "test"},
		responseChan,
		pt.SaveLog))
	engine.GetEngineInstance().AddRequests(tasks...)
	wait(len(tasks), responseChan)
	fmt.Println(h.StatusCode)
	t.Log(total)
}

func wait(wc int, responseChan chan *task.Response) {
	for {
		select {
		case response := <-responseChan:
			wc--
			if response.Parser != nil {
				if err := response.Parser(response.Result); err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println(response)
			if wc == 0 {
				return
			}
		}
	}
}

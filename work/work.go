package work

import (
	"fmt"
	"github.com/go-magic/multitask/task"
)

var (
	HandlerInvalid = fmt.Errorf("handler invalid")
)

type Work struct {
	requestChan chan *task.Request
}

func NewWork(requestChan chan *task.Request) *Work {
	return &Work{
		requestChan: requestChan,
	}
}

func (w Work) Start() {
	for {
		select {
		case request := <-w.requestChan:
			w.check(request)
		}
	}
}

func (w Work) check(request *task.Request) {
	response := task.NewResponse(request.Parser)
	defer func() {
		request.Send(response)
	}()
	if request.Handler == nil {
		response.SetError(HandlerInvalid)
		return
	}
	result, err := request.Handler.Check(request.Task)
	if err != nil {
		response.SetError(err)
	}
	response.SetResult(result)
}

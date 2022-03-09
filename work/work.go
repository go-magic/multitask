package work

import "github.com/go-magic/multitask/task"

type work struct {
	requestChan chan task.Request
}

func NewWork(requestChan chan task.Request) *work {
	return &work{
		requestChan: requestChan,
	}
}

func (w work) Start() {
	for {
		select {
		case request := <-w.requestChan:
			w.check(request)
		}
	}
}

func (w work) check(request task.Request) {
	response := task.Response{}
	response.Parser = request.Parser
	result, err := request.Handler.Check(request.Task)
	if err != nil {
		response.Error = err
	}
	response.Result = result
	request.ResponseChan <- response
}

package engine

import (
	"github.com/go-magic/multitask/task"
	"github.com/go-magic/multitask/work"
)

type Engine struct {
	requestChan chan task.Request
}

var (
	engine *Engine
)

func InitEngine(maxRoutine int) {
	engine = &Engine{}
	engine.requestChan = make(chan task.Request, 10)
	for i := 0; i < maxRoutine; i++ {
		w := work.NewWork(engine.requestChan)
		go w.Start()
	}
}

func GetEngineInstance() *Engine {
	return engine
}

func (e Engine) AddRequest(request task.Request) {
	go func() {
		e.requestChan <- request
	}()
}

func (e Engine) AddRequests(requests ...task.Request) {
	for _,request := range requests {
		e.AddRequest(request)
	}
}

package task

type Request struct {
	Handler      Handler
	Task         interface{}
	ResponseChan chan *Response
	Parser       Parser
}

func NewRequest(h Handler, task interface{},
	responseChan chan *Response, parser Parser) *Request {
	return &Request{
		Handler:      h,
		Task:         task,
		ResponseChan: responseChan,
		Parser:       parser,
	}
}

func (r Request) Send(response *Response) {
	r.ResponseChan <- response
}

type Response struct {
	Parser Parser
	Result interface{}
	Error  error
}

func NewResponse(parser Parser) *Response {
	return &Response{
		Parser: parser,
	}
}

func (r *Response) SetResult(result interface{}) {
	r.Result = result
}

func (r *Response) SetError(err error) {
	r.Error = err
}

package task

type Request struct {
	Handler      Handler
	Task         interface{}
	ResponseChan chan Response
}

type Response struct {
	Result interface{}
	Error  error
}

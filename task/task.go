package task

type Request struct {
	Handler      Handler
	Task         interface{}
	ResponseChan chan Response
	Parser       Parser
}

type Response struct {
	Parser Parser
	Result interface{}
	Error  error
}

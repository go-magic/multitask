package task

/*
Handler 任务执行实体
*/
type Handler interface {
	Check(interface{})(interface{},error)
}

/*
Parser 结果解析函数
*/
type Parser func(interface{})error

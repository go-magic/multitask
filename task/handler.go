package task

type Handler interface {
	Check(interface{})(interface{},error)
}

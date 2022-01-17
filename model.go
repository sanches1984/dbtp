package simple_proto

const delimiter = '\n'

const (
	OperationAdd    Operation = "add"
	OperationEdit   Operation = "edit"
	OperationDelete Operation = "delete"
	OperationGet    Operation = "get"
)

const (
	CodeOK            ResponseCode = 1
	CodeBadRequest    ResponseCode = 2
	CodeNotFound      ResponseCode = 3
	CodeConflict      ResponseCode = 4
	CodeInternalError ResponseCode = 5
	CodeTimeout       ResponseCode = 6
	CodeBadGateway    ResponseCode = 7
	CodeNotAccepted   ResponseCode = 8
)

type Operation string
type ResponseCode uint8
type HandleFunc func(w ResponseWriter, r *Request)
type TableProcessor map[Operation]HandleFunc

func NewTableProcessor() TableProcessor {
	return map[Operation]HandleFunc{
		OperationAdd: func(w ResponseWriter, r *Request) {
			w.WriteCode(CodeNotAccepted)
		},
		OperationEdit: func(w ResponseWriter, r *Request) {
			w.WriteCode(CodeNotAccepted)
		},
		OperationDelete: func(w ResponseWriter, r *Request) {
			w.WriteCode(CodeNotAccepted)
		},
		OperationGet: func(w ResponseWriter, r *Request) {
			w.WriteCode(CodeNotAccepted)
		},
	}
}

func (t TableProcessor) WithOperation(operation Operation, fn HandleFunc) TableProcessor {
	t[operation] = fn
	return t
}

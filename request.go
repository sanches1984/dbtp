package dbtp

import (
	"encoding/json"
)

type Request struct {
	Addr      string    `json:"-"`
	Table     string    `json:"table"`
	Operation Operation `json:"operation"`
	ObjectID  uint64    `json:"object_id"`
	Data      []byte    `json:"data"`
}

func (r Request) message() string {
	data, _ := json.Marshal(r)
	return string(data) + string(delimiter)
}

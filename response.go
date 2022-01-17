package dbtp

import "encoding/json"

type ResponseWriter interface {
	WriteCode(code ResponseCode)
	WriteData(id uint64, data []byte)
	getResponse() *Response
}

type responseWriter struct {
	response *Response
}

func newResponseWriter() ResponseWriter {
	return &responseWriter{
		response: &Response{},
	}
}

func (w *responseWriter) WriteCode(code ResponseCode) {
	w.response.Code = code
}

func (w *responseWriter) WriteData(id uint64, data []byte) {
	w.response.ObjectID = id
	w.response.Data = data
}

func (w responseWriter) getResponse() *Response {
	return w.response
}

type Response struct {
	Code     ResponseCode `json:"code"`
	ObjectID uint64       `json:"object_id"`
	Data     []byte       `json:"data"`
}

func (r Response) message() string {
	data, _ := json.Marshal(r)
	return string(data) + string(delimiter)
}

func newResponse(code ResponseCode) []byte {
	return []byte(Response{Code: code}.message())
}

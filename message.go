package dbtp

import "encoding/json"

func parseRequestMessage(msg string) (Request, error) {
	data := []byte(msg[:len(msg)-1])

	var r Request
	err := json.Unmarshal(data, &r)
	return r, err
}

func parseResponseMessage(msg string) (Response, error) {
	data := []byte(msg[:len(msg)-1])

	var r Response
	err := json.Unmarshal(data, &r)
	return r, err
}

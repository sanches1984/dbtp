package dbtp

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

var DefaultClient = client{}

type Client interface {
	Do(req *Request) (*Response, error)
}

type client struct{}

func (c *client) Do(req *Request) (*Response, error) {
	conn, err := net.Dial("tcp", req.Addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	_, err = fmt.Fprint(conn, req.message())
	if err != nil {
		return nil, err
	}

	message, err := bufio.NewReader(conn).ReadString(delimiter)
	if err != nil {
		return nil, err
	}

	resp, err := parseResponseMessage(message)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func NewRequest(url string, data []byte) (*Request, error) {
	urlParams := strings.Split(url, "/")
	if len(urlParams) != 3 {
		return nil, errors.New("bad url path")
	}

	return &Request{
		Addr:      urlParams[0],
		Table:     urlParams[1],
		Operation: Operation(urlParams[2]),
		Data:      data,
	}, nil
}

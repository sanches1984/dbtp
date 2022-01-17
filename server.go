package simple_proto

import (
	"bufio"
	"io"
	"log"
	"net"
)

type Server interface {
	Handle(table string, proc TableProcessor)
	Listen(addr string) error
}

type proto struct {
	tables map[string]TableProcessor
}

func NewServer() Server {
	return &proto{
		tables: make(map[string]TableProcessor),
	}
}

func (p *proto) Handle(table string, proc TableProcessor) {
	p.tables[table] = proc
}

func (p *proto) Listen(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		message, err := bufio.NewReader(conn).ReadString(delimiter)
		if err != nil {
			if err == io.EOF {
				continue
			}
			panic(err)
		}

		log.Println("got message")
		request, err := parseRequestMessage(message)
		if err != nil {
			conn.Write(newResponse(CodeInternalError))
			continue
		}
		request.Addr = addr

		processor, ok := p.tables[request.Table]
		if !ok {
			conn.Write(newResponse(CodeBadGateway))
			continue
		}

		operationFn, ok := processor[request.Operation]
		if !ok {
			conn.Write(newResponse(CodeBadGateway))
			continue
		}

		resp := newResponseWriter()
		operationFn(resp, &request)

		conn.Write([]byte(resp.getResponse().message()))
	}
}

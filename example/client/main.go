package main

import (
	dbtp "github.com/sanches1984/simple-proto"
	"log"
)

func main() {
	request, err := dbtp.NewRequest("127.0.0.1:8086/table/add", []byte("some data"))
	if err != nil {
		panic(err)
	}

	log.Println("do request")
	response, err := dbtp.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	log.Printf("response: %d %s", response.Code, response.Data)
}

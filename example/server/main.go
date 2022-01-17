package main

import (
	"github.com/sanches1984/dbtp"
	"log"
)

func main() {
	proc := dbtp.NewTableProcessor().
		WithOperation(dbtp.OperationAdd, func(w dbtp.ResponseWriter, r *dbtp.Request) {
			// some processing
			log.Println("handle add")
			w.WriteCode(dbtp.CodeOK)
		}).
		WithOperation(dbtp.OperationGet, func(w dbtp.ResponseWriter, r *dbtp.Request) {
			// some processing
			log.Println("handle get")
			w.WriteCode(dbtp.CodeOK)
			w.WriteData(r.ObjectID, r.Data)
		})

	srv := dbtp.NewServer()
	srv.Handle("/table_name", proc)

	log.Println("start listener on 8086")
	if err := srv.Listen("127.0.0.1:8086"); err != nil {
		panic(err)
	}
}

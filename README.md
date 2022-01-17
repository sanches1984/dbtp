# Database Transport Protocol (DBTP)
Custom protocol over TCP for db operations.

## Install package

    go get -v github.com/sanches1984/dbtp

## Server
Example:
```go
tblProc := dbtp.NewTableProcessor().
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

server := dbtp.NewServer()
server.Handle("/table_name", tblProc)

server.Listen(":8086")
```

## Client
Supported operations:
- add
- edit
- delete
- get

Example:
```go
request, _ := dbtp.NewRequest(":8086/table_name/edit", []byte("some data"))
request.ObjectID = 1

response, _ := dbtp.DefaultClient.Do(request)
```
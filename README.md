# grpc-demo
You can:
- Find a student by *id*
- Create a new student with *id* and *name*
- Get a list of students filtered by *keyword*

### How to run

1. Generate gRPC code
``` sh
$ protoc -I student/ student/student.proto --go_out=plugins=grpc:student
```
2. Run gRPC server
``` sh
$ go run server/server.go
```
3. Run the client
- Find a student
``` sh
$ go run client/client.go -func find -id 102140023
```
- Create a student
``` sh
$ go run client/client.go -func create -id 102140024 -name "Foo Bar"
```
- Get a list of students (*keyword* option is optional)
``` sh
$ go run client/client.go -func get -keyword "Foo"
```
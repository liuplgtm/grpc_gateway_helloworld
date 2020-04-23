after cloning the project, run this in side the project folder.
```
$ go mod init nice
$ protoc -I/usr/local/include -I. \
  -I$GOPATH/src \  
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  hello/hello.proto  
$ protoc -I/usr/local/include -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. \
    hello/hello.proto
```
Note: protoc generates the necessary stub for gRPC and gRPC_gateway.

terminal 1:
```
$ go run server/main.go
```

terminal 2:
```
$ go run gateway/main.go
```
Note: gateway wraps around the endpoint localhost:9090 provided by terminal 1.

terminal 3:
```
$ curl -X POST -i http://localhost:8080/v1/example/echo/abcde -d '{"num": "100"}'
```
HTTP/1.1 200 OK
Content-Type: application/json
Grpc-Metadata-Content-Type: application/grpc
Date: Thu, 23 Apr 2020 17:27:28 GMT
Content-Length: 26

{"id":"abcde","num":"100"}%

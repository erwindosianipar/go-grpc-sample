# go-grpc-sample

> Sample basic implementation gRPC in Golang

## Serve

> Run the gRPC server and client

``` bash
make run
```

## API

> Fetch Github account with username given

``` bash
GET http://localhost:50052/github/{username}
Content-Type: application/json
```

## Protobuf

> Generate gRPC protobuf file

``` bash
make protobuf
```

> Generate protobuf gateway file

## Gateway

``` bash
make gateway
```

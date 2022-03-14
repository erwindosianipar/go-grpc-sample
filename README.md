# go-grpc-sample

> Sample basic implementation gRPC in Golang

## Serve

> Run server and client then hit API to make request

``` bash
make server -B
```

``` bash
make client -B
```

## API

``` bash
curl http://localhost:8080/erwindosianipar
```

## Generate

``` bash
protoc proto/github.proto \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative
```

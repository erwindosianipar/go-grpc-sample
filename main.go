package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/erwindosianipar/go-grpc-sample/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

const (
	HOST_SERVER = ":50051"
	HOST_CLIENT = ":50052"
)

type githubServiceServer struct {
	proto.UnimplementedGithubServiceServer
}

func (s *githubServiceServer) Github(ctx context.Context, in *proto.GithubInput) (*proto.GithubOutput, error) {
	resp, err := http.Get("https://api.github.com/users/" + in.GetUsername())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var output proto.GithubOutput
	err = json.Unmarshal(content, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func main() {
	go func() {
		mux := runtime.NewServeMux()
		err := proto.RegisterGithubServiceHandlerServer(context.Background(), mux, &githubServiceServer{})
		if err != nil {
			log.Fatalf("failed to register github service handler: %v", err)
		}

		log.Fatal(http.ListenAndServe(HOST_CLIENT, mux))
	}()

	listener, err := net.Listen("tcp", HOST_SERVER)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterGithubServiceServer(server, &githubServiceServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

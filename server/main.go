package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/erwindosianipar/go-grpc-sample/proto"

	"google.golang.org/grpc"
)

const (
	HOST = ":50051"
)

type GithubServiceServer struct {
	proto.UnimplementedGithubServiceServer
}

func (s *GithubServiceServer) Github(ctx context.Context, in *proto.GithubInput) (*proto.GithubOutput, error) {
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
	json.Unmarshal(content, &output)

	return &output, nil
}

func main() {
	listener, err := net.Listen("tcp", HOST)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGithubServiceServer(s, &GithubServiceServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"log"

	"github.com/erwindosianipar/go-grpc-sample/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	pb := proto.NewGithubServiceClient(conn)

	g := gin.New()
	g.GET("/:username", func(c *gin.Context) {
		res, err := pb.Github(context.Background(), &proto.GithubInput{Username: c.Param("username")})
		if err != nil {
			log.Fatalf("could not fetch: %v", err)
		}

		c.JSON(200, res)
	})

	g.Run(":8080")
}

package main

import (
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
	"github.com/luynt/note/grpc/pb"
	"context"
)

func main() {
	conn, err := grpc.Dial("localhost:7878", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := "sdhajshdjhsjdhja shdjsahdjshd"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}



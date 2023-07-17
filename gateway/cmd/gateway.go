package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/leoomi/remnants/auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:13370", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthenticationClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Verify(ctx, &pb.VerifyRequest{
		Token: "banaba",
	})
	if err != nil {
		log.Fatalf("could not verify: %v", err)
	}
	log.Printf("Greeting: %s", r.String())
}

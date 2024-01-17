package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"ticketing/ent/proto/entpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Open a connection to the server.
	conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	// Create a User service Client on the connection.
	client := entpb.NewConcertServiceClient(conn)

	// Ask the server to create a random User.
	ctx := context.Background()
	created, err := client.Create(ctx, &entpb.CreateConcertRequest{
		Concert: &entpb.Concert{
			Title:    fmt.Sprintf("World Tour %d", rand.Intn(100)),
			Artist:   "Justin Beiber",
			ImageUrl: "hellos",
		},
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("user created with id: %d", created.Id)

	// On a separate RPC invocation, retrieve the user we saved previously.
	get, err := client.Get(ctx, &entpb.GetConcertRequest{
		Id: created.Id,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved user with id=%d: %v", get.Id, get)
}

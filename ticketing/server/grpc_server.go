package main

import (
	"context"
	"log"
	"net"
	"os"

	"ticketing/ent"
	"ticketing/ent/proto/entpb"

	"entgo.io/ent/dialect"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("failed loading env: %v", err)
	}
	DATABASE_URL := os.Getenv("POSTGRES_ADDRESS")
	// Create an ent.Client for local postgres. Disable SSL usage (default is enabled).
	client, err := ent.Open(dialect.Postgres, DATABASE_URL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize the generated User service.
	concertSvc := entpb.NewConcertService(client)
	concertSessionSvc := entpb.NewConcertSessionService(client)
	sectionSvc := entpb.NewSectionService(client)
	ticketSvc := entpb.NewTicketService(client)

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

	// Register the services with the server.
	entpb.RegisterConcertServiceServer(server, concertSvc)
	entpb.RegisterConcertSessionServiceServer(server, concertSessionSvc)
	entpb.RegisterSectionServiceServer(server, sectionSvc)
	entpb.RegisterTicketServiceServer(server, ticketSvc)

	// Open port 5003 for listening to traffic.
	lis, err := net.Listen("tcp", ":5003")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}

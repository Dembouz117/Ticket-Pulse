package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"ticketing/ent"
	"ticketing/ent/migrate"

	"entgo.io/ent/dialect"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import for side effect: registers postgres driver.
)

func Example() {
	err := godotenv.Load("../../.env")
	DATABASE_URL := os.Getenv("POSTGRES_ADDRESS")
	// Create an ent.Client for local postgres. Disable SSL usage (default is enabled).
	client, err := ent.Open(dialect.Postgres, DATABASE_URL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// if err := client.Schema.Create(ctx); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
	task1, err := client.Concert.Create().SetArtist("Justin Beiber").SetTitle("Tour").SetImageUrl("example.com/image.jpg").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a user: %v", err)
	}
	fmt.Println(task1)
	task2, err2 := client.Concert.Create().SetArtist("Ariana Grande").SetTitle("Tour").SetImageUrl("example23.com/image.jpg").Save(ctx)
	if err2 != nil {
		log.Fatalf("failed creating a user: %v", err2)
	}
	fmt.Println(task1)
	fmt.Println(task2)

	// Output:
	// User()
}

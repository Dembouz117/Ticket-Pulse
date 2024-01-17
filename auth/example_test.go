package main

import (
	"auth/ent"
	"context"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import for side effect: registers postgres driver.
)

func Example() {
	err := godotenv.Load("../.env")
	DATABASE_URL := os.Getenv("POSTGRES_ADDRESS")
	
	// Create an ent.Client for local postgres. Disable SSL usage (default is enabled).
	client, err := ent.Open(dialect.Postgres, DATABASE_URL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	task1, err := client.User.Create().SetEmail("changxianxiang@gmail.com").SetName("Chang Xian Xiang").SetPhone("99999999").SetPassword("P@ssw0rd").SetRole("admin").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a user: %v", err)
	}
	fmt.Println(task1)
	// Output:
	// User()
}

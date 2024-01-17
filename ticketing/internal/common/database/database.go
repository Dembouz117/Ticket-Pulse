package database

import (
	"context"
	"log"
	"ticketing/ent"

	"entgo.io/ent/dialect"
)

func ConnectDatabase(postgresAddress string, ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open(dialect.Postgres, postgresAddress)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}
	return client, nil
}
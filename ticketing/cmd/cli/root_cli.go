package cli

import (
	"context"
	"fmt"
	"log"
	"os"
	"ticketing/ent"
	"ticketing/internal/common/cache"
	"ticketing/internal/common/database"
	"ticketing/internal/common/database/seed"
	"ticketing/internal/common/repository"

	"github.com/spf13/cobra"
)

var (
	rootCmd     = &cobra.Command{Use: "myapp"}
	ctx         = context.Background()
	client      *ent.Client
	redisClient *cache.RedisCache
)

var clearDataCmd = &cobra.Command{
	Use:   "cleardata",
	Short: "Clears all concert, session, section, and ticket records",
	Run: func(cmd *cobra.Command, args []string) {
		clearData(client, redisClient)
		os.Exit(1)
	},
}

var seedDataCmd = &cobra.Command{
	Use:   "seeddata",
	Short: "Seed data for the application",
	Run: func(cmd *cobra.Command, args []string) {
		seedData(client, redisClient)
		os.Exit(1)
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migratedata",
	Short: "Apply migration on the database",
	Run: func(cmd *cobra.Command, args []string) {
		migrateDatabase(client)
		os.Exit(1)
	},
}

func Execute(postgresAddress string, redisAddress string) {

	dbClient, err := database.ConnectDatabase(postgresAddress, ctx)
	if err != nil {
		// Handle the database connection error
		log.Printf("Failed to connect to the database: %s", err)
		return
	}
	defer dbClient.Close()

	client = dbClient

	redisClientConnect, err := cache.ConnectRedisCache(redisAddress)
	if err != nil {
		log.Printf("Failed to connect to the cache: %s", err)
		return
	}
	redisClient = cache.NewRedisCache(redisClientConnect)

	// Add the subcommands to the root command
	rootCmd.AddCommand(clearDataCmd)
	rootCmd.AddCommand(seedDataCmd)
	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func seedData(client *ent.Client, redisClient *cache.RedisCache) {
	// If we don't have any concert yet, seed the database.
	if !client.Concert.Query().ExistX(ctx) {
		concertRepo := repository.NewConcertRepository(client, redisClient)
		sessionRepo := repository.NewSessionRepository(client)
		sectionRepo := repository.NewSectionRepository(client, redisClient)
		ticketRepo := repository.NewTicketRepository(client, redisClient)

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}

		fmt.Printf("Current working directory: %s\n", currentDir)

		filePath := "seed.json"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Printf("The JSON file '%s' does not exist.", filePath)
			return
		}
		if err := seed.SeedDatabase(ctx, concertRepo, sessionRepo, sectionRepo, ticketRepo, filePath); err != nil {
			log.Fatalf("failed seeding the database: %v", err)
		}
	}
}

func clearData(client *ent.Client, redisClient *cache.RedisCache) {
	// Create repository instances
	concertRepo := repository.NewConcertRepository(client, redisClient)
	sessionRepo := repository.NewSessionRepository(client)
	sectionRepo := repository.NewSectionRepository(client, redisClient)
	ticketRepo := repository.NewTicketRepository(client, redisClient)

	// Delete all tickets
	if err := ticketRepo.DeleteAllTickets(context.Background()); err != nil {
		log.Fatalf("Failed to delete tickets: %v", err)
	}
	fmt.Println("All ticket records deleted.")

	// Delete all sections
	if err := sectionRepo.DeleteAllSections(context.Background()); err != nil {
		log.Fatalf("Failed to delete sections: %v", err)
	}
	fmt.Println("All section records deleted.")

	// Delete all sessions
	if err := sessionRepo.DeleteAllSessions(context.Background()); err != nil {
		log.Fatalf("Failed to delete sessions: %v", err)
	}
	fmt.Println("All session records deleted.")

	// Delete all concerts
	if err := concertRepo.DeleteAllConcerts(context.Background()); err != nil {
		log.Fatalf("Failed to delete concerts: %v", err)
	}
	fmt.Println("All concert records deleted.")
}

func migrateDatabase(client *ent.Client) {
	fmt.Println("Migrating database")
	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	fmt.Println("Migration has been applied")
}

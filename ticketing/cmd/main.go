//	@title			TicketPulse Ticketing API
//	@version		1.0
//	@description	This is the TicketPulse API for ticketing microservice. Concerts, sessions, sections and tickets are all here.

//	@BasePath	/api/v1

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"ticketing/cmd/cli"
	docs "ticketing/docs"
	"ticketing/ent"
	"ticketing/ent/migrate"
	"ticketing/ent/proto/entpb"
	adminRoutes "ticketing/internal/admin/routes"
	"ticketing/internal/common/cache"
	"ticketing/internal/common/database"
	"ticketing/internal/common/middleware"
	userRoutes "ticketing/internal/user/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

var (
	redisAddress    string
	postgresAddress string
	secretKey       string
	ctx             = context.Background()
)

func automaticMigration(client *ent.Client) {
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func setUpRouter(client *ent.Client, redisClient *cache.RedisCache) *gin.Engine {
	// Initialize the Gin router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://sg1.biddlr.com", "https://sg1.biddlr.com"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/ticketing/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/ticketing/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ticketing microservice is running",
		})
	})

	userPublicRouterGroup := router.Group("/ticketing/api/v1")
	{
		userRoutes.SetupPublicRoutes(userPublicRouterGroup, client, redisClient)
	}

	router.Use(middleware.AuthMiddleware(secretKey))
	userProtectedRouterGroup := router.Group("/ticketing/api/v1")
	{
		userRoutes.SetupProtectedRoutes(userProtectedRouterGroup, client, redisClient)
	}

	adminRouterGroup := router.Group("/ticketing/api/v1/admin", middleware.AdminMiddleware(secretKey))
	{
		adminRoutes.SetupAdminRoutes(adminRouterGroup, client, redisClient)
	}
	return router
}

func setUpgRPC(dbClient *ent.Client) {
	// Initialize the generated User service.
	concertSvc := entpb.NewConcertService(dbClient)
	concertSessionSvc := entpb.NewConcertSessionService(dbClient)
	sectionSvc := entpb.NewSectionService(dbClient)
	ticketSvc := entpb.NewTicketService(dbClient)

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

func main() {

	godotenv.Load("../../.env")
	redisAddress = os.Getenv("REDIS_ADDRESS")
	postgresAddress = os.Getenv("POSTGRES_ADDRESS")
	secretKey = os.Getenv("SECRET_KEY")

	cli.Execute(postgresAddress, redisAddress)

	dbClient, err := database.ConnectDatabase(postgresAddress, ctx)
	if err != nil {
		// Handle the database connection error
		log.Printf("Failed to connect to the database: %s", err)
		return
	}
	defer dbClient.Close()

	automaticMigration(dbClient)

	docs.SwaggerInfo.BasePath = "/api/v1"

	go setUpgRPC(dbClient)

	redisClientConnect, err := cache.ConnectRedisCache(redisAddress)
	if err != nil {
		log.Printf("Failed to connect to the database: %s", err)
		return
	}
	redisClient := cache.NewRedisCache(redisClientConnect)

	// Initialize the Gin router
	router := setUpRouter(dbClient, redisClient)
	fmt.Println("Ticketing microservice is running on :8081")
	router.Run(":8081")
}

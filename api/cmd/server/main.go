package main

import (
	"log"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neilchetna/mind-mapper/internal/rest"
	"github.com/neilchetna/mind-mapper/pkg"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func establishClerkClient() {
	clerkSecret := os.Getenv("CLERK_SECRET_KEY")
	clerk.SetKey(clerkSecret)
}

func connectToQueue() *asynq.Client {
	queueHost := os.Getenv("REDIS_DB_ADDR")
	q := asynq.NewClient(asynq.RedisClientOpt{Addr: queueHost})
	return q
}

func startHTTPServer(db *gorm.DB, q *asynq.Client) error {
	e := echo.New()

	e.Use(middleware.Logger())

	rest.BuildRoutes(e, db, q)

	const defaultAddress = ":8080"
	return e.Start(defaultAddress)
}

func main() {
	// Load env variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(pkg.ErrLoadingENV.Error())
	}

	// Connect to the database
	db, err := pkg.ConnectToDatabase()
	if err != nil {
		log.Fatalf("%w: %v", pkg.ErrConnectionDB.Error(), err)
	}
	pool, _ := db.DB()
	defer pool.Close()

	// Connect to queue
	qdb := connectToQueue()
	defer qdb.Close()

	// Init Clerk Client
	establishClerkClient()

	// Build and start HTTP server
	if err := startHTTPServer(db, qdb); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/neilchetna/mind-mapper/internal/worker"
	"github.com/neilchetna/mind-mapper/pkg"
)

func main() {
	// Load env variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(pkg.ErrLoadingENV.Error())
	}

	// Connect to the database
	db, err := pkg.ConnectToDatabase()
	if err != nil {
		log.Fatalf("%w: %v", pkg.ErrConnectionDB, err)
	}
	pool, _ := db.DB()
	defer pool.Close()

	queueHost := os.Getenv("REDIS_DB_ADDR")

	// Connect to the queue
	q := asynq.RedisClientOpt{Addr: queueHost}
	srv := asynq.NewServer(q, asynq.Config{Concurrency: 10})

	// Register job handler
	mux := asynq.NewServeMux()
	w := worker.SuggestNodesWorkerBuilder(db)
	mux.HandleFunc(pkg.Tasks.SuggestNodes, w.SuggestNodesHandler)

	log.Print("PING: worker")
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

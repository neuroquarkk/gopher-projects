package main

import (
	"context"
	"log"
	"mtt/internal/platform/config"
	"mtt/internal/platform/database"
	"mtt/internal/platform/middleware"
	"mtt/internal/tasks"
	"net/http"
)

func main() {
	ctx := context.Background()

	config := config.LoadConfig()
	pool := database.ConnDB(ctx, config.DB_URL)
	defer func() {
		log.Println("Closing database")
		pool.Close()
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from the server"))
	})

	taskHandler := tasks.NewHandler(pool)
	taskHandler.RegisterRoutes(mux)

	globalHandler := middleware.LoggingMiddleware(mux)

	server := &http.Server{
		Addr:    ":" + config.PORT,
		Handler: globalHandler,
	}
	log.Printf("Starting server on PORT %s\n", config.PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error: failed to start the server %v\n", err)
	}
}

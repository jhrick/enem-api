package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jhrick/enem/api/internal/core/infra/config"
	"github.com/jhrick/enem/api/internal/core/infra/db/postgres"
	"github.com/jhrick/enem/api/internal/core/infra/http/handler"
	"github.com/jhrick/enem/api/internal/core/infra/utils/env"
)

func main() {
  env.Load()
  config := config.NewConfig()

  ctx := context.Background()

  database := postgres.NewPostgres(ctx, *config.DB)
  if err := database.Connect(); err != nil {
    panic(err)
  }

  handler := handler.NewHandler()

  go func() {
    if err := http.ListenAndServe(config.Port, handler.Router); err != nil {
      if !errors.Is(err, http.ErrServerClosed) {
        panic(err)
      }
    }
  }()
  
  log.Println("Server Running!")

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, os.Interrupt)
  <-quit

  log.Println("Server Closed.")
}

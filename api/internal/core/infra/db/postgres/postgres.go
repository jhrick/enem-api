package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jhrick/enem/api/internal/core/infra/config"
)

type Postgres struct {
  Ctx    context.Context
  Pool   *pgxpool.Pool
  Config config.DB
}

func NewPostgres(ctx context.Context, config config.DB) *Postgres {
  return &Postgres{
    Ctx: ctx,
    Config: config,
  }
}

func (pg *Postgres) Connect() error {
  pool, err := pgxpool.New(pg.Ctx, fmt.Sprintf(
    "user=%s password=%s host=%s port=%s dbname=%s",
    pg.Config.User,
    pg.Config.Password,
    pg.Config.Host,
    pg.Config.Port,
    pg.Config.DB,
  ))
  if err != nil {
    return fmt.Errorf("database connection error: %v", err)
  }

  log.Println("Database connection successful!")
  pg.Pool = pool

  return nil
}

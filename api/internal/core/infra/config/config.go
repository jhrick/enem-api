package config

import "github.com/jhrick/enem/api/internal/core/infra/utils/env"

type Config struct {
  Port string
  DB   *DB
}

type DB struct {
  Host     string
  Port     string
  User     string
  DB       string
  Password string
}

func NewConfig() *Config {
  port := ":" + env.GetEnvOrDefault("APP_PORT", "8080")

  return &Config{
    Port: port,
    DB: &DB{
      Host: env.GetEnvOrDefault("DATABASE_HOST", "localhost"),
      Port: env.GetEnvOrDefault("DATABASE_PORT", "5432"),
      User: env.GetEnvOrError("DATABASE_USER"),
      DB: env.GetEnvOrError("DATABASE_NAME"),
      Password: env.GetEnvOrError("DATABASE_PASSWORD"),
    },
  }
}

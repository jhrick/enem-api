package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
  if err := godotenv.Load("../.env"); err != nil {
    panic(err)
  }
}

func GetEnvOrError(key string) string {
  env := os.Getenv(key)

  if len(env) == 0 {
    log.Fatalf("not found %s environment variable\n", key)
  }

  return env
  
}

func GetEnvOrDefault(key, defaultValue string) string {
  env := os.Getenv(key)

  if len(env) == 0 {
    env = defaultValue
  }

  return env
}

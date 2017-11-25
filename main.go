package main

import (
  "dict/api"
  env "github.com/joho/godotenv"
  "log"
)

func main() {
  err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

  api.Run();
}

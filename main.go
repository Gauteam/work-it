package main

import (
	"log"

	"github.com/Gauteam/work-it/internal/pkg/env"
)

func main() {
	log.Println("hello hello")

	log.Printf("%s: %s", "DB_HOST", *env.DbHost)
	log.Printf("%s: %s", "DB_NAME", *env.DbName)
	log.Printf("%s: %s", "DB_PORT", *env.DbPort)
	log.Printf("%s: %s", "DB_USER", *env.DbUser)
}

package main

import (
	"log"

	"github.com/iamrajiv/neon-db-cli/cmd"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rootCmd := cmd.GetRootCommand()
	rootCmd.Execute()
}

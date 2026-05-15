package main

import (
	"log"

	"github.com/jiehui555/rpitool/cmd"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cmd.Execute()
}

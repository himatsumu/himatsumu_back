package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}
	fmt.Println(os.Getenv("GO_URL"))
}

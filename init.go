package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {

	err := godotenv.Overload(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	fmt.Println(os.Getenv("GO_URL"))
}

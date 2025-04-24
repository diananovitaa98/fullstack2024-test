package main

import (
	"fmt"
	"fullstacktest/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app.Run()
}

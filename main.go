package main

import (
	"dvdrental/helper"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helper.LogErrorAndPanic(err)

	fmt.Println("This is simple API for Dvdrental Database")
}

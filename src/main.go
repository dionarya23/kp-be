package main

import (
	"fmt"

	"github.com/dionarya23/kredit-plus/src/drivers/db"
	"github.com/dionarya23/kredit-plus/src/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbConnection, err := db.CreateConnection()
	if err != nil {
		fmt.Println("Error creating database connection:", err)
		return
	}

	defer func() {
		if err := dbConnection.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	h := http.New(&http.Http{
		DB: dbConnection,
	})

	h.Launch()
}

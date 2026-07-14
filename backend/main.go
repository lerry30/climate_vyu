package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load .env locally; on Vercel the variables are already
// injected into the environment, so a missing file here is fine.

func main() {
	// It requires to load the godotenv to access .env file
	// contents
	err := godotenv.Load()
	if err != nil {
		//log.Fatal("Error loading .env file")
		log.Println("No .env file found, using system environment variables")
	}

	key := os.Getenv("OPEN_WEATHER_API_KEY")
	ow := NewOpenWeather(key)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // local fallback
	}

	server := NewAPIServer(":" + port)
	server.AddExternalAPI("open-weather", ow)
	server.Run()
}

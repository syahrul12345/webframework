package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"webframework/controller"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		log.Println(e)
		log.Println("Checking machine environment variables instead of in the .env file... this might be a production build")
		prod := os.Getenv("is_production")
		if prod != "" {
			log.Printf("Production status : %s\n", prod)
		} else {
			log.Println("Prudction status: False. Error reading environment variables")
		}

	}
	// Create the Webserver
	mux := http.NewServeMux()
	// Define rest endpoints
	mux.HandleFunc("/", controller.Serve)
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	// Allow the default testing port for CORS
	prod, _ := strconv.ParseBool(os.Getenv("is_production"))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: !prod,
	})
	handler := c.Handler(mux)
	// Get port from env file
	port := os.Getenv("port")
	log.Printf("Webserver is on http://127.0.0.1:%s\n", port)
	http.ListenAndServe(":"+port, handler)
}

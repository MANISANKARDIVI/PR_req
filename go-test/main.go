package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

// helloHandler handles HTTP requests to the root URL.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Log a message when this handler is called
	log.Info("helloHandler called")
	fmt.Fprint(w, "Hello, World!") 
}

// infoHandler demonstrates an info-level log
func infoHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("infoHandler called")
	fmt.Fprint(w, "This is an info message!")
}

// errorHandler demonstrates an error-level log
func errorHandler(w http.ResponseWriter, r *http.Request) {
	log.Error("errorHandler called")
	http.Error(w, "This is an error message!", http.StatusInternalServerError)
}

// main function sets up the HTTP server.
func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// Setting up handlers
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/error", errorHandler)

	fmt.Println("Server is running on http://localhost:8080")
	// Start the server and log any error if it fails to start #added this line
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to indicate the type of data being returned
	w.Header().Set("Content-Type", "text/plain")

	// Write the response data to the http.ResponseWriter
	fmt.Fprintf(w, "Hello world from %s in %s deployed in %s", os.Getenv("PODNAME"), os.Getenv("PODNAMESPACE"), os.Getenv("NODENAME"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

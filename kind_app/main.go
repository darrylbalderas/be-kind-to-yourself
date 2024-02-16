package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello world from %s in %s deployed in %s", os.Getenv("PODNAME"), os.Getenv("PODNAMESPACE"), os.Getenv("NODENAME"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

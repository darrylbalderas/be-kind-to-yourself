package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("Hello world from %s", os.Getenv("PODNAME"))
}

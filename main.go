package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Define command-line flags
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	// Create and start the server
	server := NewServer(*port)
	fmt.Printf("Starting TCP proxy server on port %d...\n", *port)
	
	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

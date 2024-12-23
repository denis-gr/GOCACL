// Entrypoint for the server
package main

import (
	"flag"
	"gocalc/internal/server"
)

func main() {
	ipPort := flag.String("ipPort", ":8080", "IP:Port to listen on")
	flag.Parse()
	server.StartServer(*ipPort)
}

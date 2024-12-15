package main

import (
	"github.com/HaikalRFadhilahh/go-todos-list/internal/server"
)

func main() {
	// Start Server
	server.InitApiServer().RunServer()
}

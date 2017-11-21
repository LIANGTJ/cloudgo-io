package main

import (
	"os"
	"service"
	flag "github.com/spf13/pflag"
)

const (
	// default port 
	PORT string = "8080"
	
)

func main() {
	port := os.Getenv("PORT")
	portPtr := flag.StringP("port","p",PORT, "Port for http listening")
	flag.Parse()
	if len(*portPtr) != 0 {
		port = *portPtr
	}
	server := service.NewServer()
	server.Run(":" + port)
}
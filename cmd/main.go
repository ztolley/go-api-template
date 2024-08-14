package main

import "github.com/ztolley/goapi/cmd/api"

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}

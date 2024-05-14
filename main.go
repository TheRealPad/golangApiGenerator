package main

import (
	"httpServer/src/core"
)

func main() {
	api := core.ApiService{Api: &core.Api{Port: 8080}}
	api.Listen()
}

package main

import (
	"fmt"
	"httpServer/src/core"
	"httpServer/src/initialisation"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Need one argument - the json file containing the configuration")
		return
	}
	api := core.ApiService{Api: &core.Api{Json: initialisation.JsonHandler{File: args[1]}}}
	api.Listen()
}

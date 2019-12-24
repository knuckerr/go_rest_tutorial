package main

import (
	"github.com/knuckerr/go_rest/api/controllers"
)

func main() {
	var server = controllers.Server{}
	server.Run()
}

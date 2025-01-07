package main

import (
	"github.com/bakigoal/go-fiber-example/cmd"
	"log"
)

func main() {
	app := cmd.NewApp()
	log.Fatal(app.Server.Listen(":3000"))
}

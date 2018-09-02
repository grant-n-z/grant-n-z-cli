package main

import (
	"os"

	"github.com/urfave/cli"
)

var (
	app cli.App
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Usage = "make an explosive entrance"
	app.Run(os.Args)
}

func setUri() {

}

func createUser() {

}

func getToken() {

}

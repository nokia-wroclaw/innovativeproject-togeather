package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func info(app *cli.App) {
	app.Name = "TogEATher"
}

func commands(app *cli.App) {
	app.Commands = []cli.Command {
		{
			Name: "server",
			Usage: "Run the app server",
			Action: func(c *cli.Context) {
				runApp()
			},
		},
	}
}

func main() {
	app := cli.NewApp()
	info(app)
	commands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runApp() {
	// do all the connections here, pass config
	redis, err := redisConnect()
	if err != nil {
		log.Fatalf("redis connection error: %s", err.Error())
	}
	defer redis.Close()

	runServer()
}


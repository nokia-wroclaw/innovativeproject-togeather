package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/urfave/cli"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/lobby"
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/restaurant"
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/user"
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

	db, err := dbConnect()
	if err != nil {
		log.Fatalf("db connection error: %s", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	defer db.Close()

	restaurantStore := restaurant.NewStore(db)
	restaurantService := restaurant.NewService(restaurantStore)

	userStore := user.NewStore(db)
	userService := user.NewService(userStore)

	lobbyStore := lobby.NewStore(db)
	lobbyService := lobby.NewService(lobbyStore, userService)

	runServer(
		restaurantService,
		lobbyService,
		userService,
	)
}


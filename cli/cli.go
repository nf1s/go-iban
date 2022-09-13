package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "CLI to populate values to Postgres db"
	app.Usage = "CLI to populate values to Postgres db"
	app.Author = "P.F.C"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "populate",
			Aliases: []string{"p"},
			Usage:   "parse csv and adds data to postgres db",
			Action: func(c *cli.Context) {
				countries := readCSVFile("../data/ibans.csv")
				populateIbans(countries)
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name:  "XorHunt",
	Usage: "Open Source Recon Tools From Template .yaml",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "lang",
			Value:    "english",
			Usage:    "language for the greeting",
			Required: true,
		},
	},
	Action: func(*cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

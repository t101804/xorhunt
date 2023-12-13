package main

import (
	"log"
	"os"

	"github.com/t101804/xorhunt/options"
	"github.com/t101804/xorhunt/pkg/runner"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "XorHunt"
	app.Usage = "Massive Open Source Custom Recon & Scanner Tools"
	opts := options.GetOpts(app)
	app.Action = func(ctx *cli.Context) error {
		opts.ValOpts()
		hunt, err := runner.NewRunner(opts)
		if err != nil {
			return err
		}
		err = hunt.Start()
		if err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/t101804/xorhunt/options"
	"github.com/t101804/xorhunt/pkg/templateparser"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "XorHunt"
	app.Usage = "Massive Open Source Custom Recon & Scanner Tools"
	opts := options.GetOpts(app)
	app.Action = func(ctx *cli.Context) error {
		opts.ValOpts()
		tmpl := templateparser.ReadTemplate(opts)
		fmt.Println(tmpl.Path)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

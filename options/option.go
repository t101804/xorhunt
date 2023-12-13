package options

import (
	"io"

	"github.com/urfave/cli/v2"
)

type App *cli.App

type GlobalOptions struct {
	TemplateName      string // Use a template file .yaml
	IterNum           int    // Set an iterator number for templates to be value in the path
	IterFile          string // set an iterator file for templates to be value in the path
	Verbose           bool   // Verbose logging
	ChangeToIp        bool   // Using a address ip
	ExcludeIps        bool   // Exclue using Ips or no in iterator file
	Threads           int    // Threads controls the number of threads to use for active enumerations
	Timeout           int    // Timeout is the seconds to wait for sources to respond
	Config            string // Config contains the location of the config file
	Proxy             string // HTTP proxy
	Output            io.Writer
	RateLimit         int  // Global maximum number of HTTP requests to send per second
	DisableAutoUpdate bool // DisableUpdateCheck disable update checking
}

func GetOpts(app *cli.App) *GlobalOptions {
	opt := &GlobalOptions{}

	// TODO : Separate scan and recon mode ( use subcommands : https://cli.urfave.org/v2/examples/subcommands/ )
	flg := []cli.Flag{
		&cli.IntFlag{Destination: &opt.IterNum, Name: "num", Aliases: []string{"n"}, Usage: "Set an iterator number for templates to be value in the path"},
		&cli.StringFlag{Destination: &opt.IterFile, Name: "list", Aliases: []string{"l"}, Usage: "set an iterator file for templates to be value in the path e.g for ip list or domain lists file"},
		&cli.StringFlag{Destination: &opt.TemplateName, Name: "template", Aliases: []string{"tmpl"}, Usage: "Choose a Template file e.g (templates/autorecon.yaml)", Required: true},
		&cli.BoolFlag{Destination: &opt.Verbose, Name: "verbose", Usage: "Verbose logging", Value: false},
		&cli.BoolFlag{Destination: &opt.ChangeToIp, Name: "changetoip", Aliases: []string{"cip"}, Usage: "Change your lists into ip [ e.g: google.com will be 1.1.1.1 ] in iterator file", Value: false},
		&cli.BoolFlag{Destination: &opt.ExcludeIps, Name: "excludeip", Aliases: []string{"xip"}, Usage: "Excluding using ip in iterator file", Value: false},
		&cli.IntFlag{Destination: &opt.Threads, Name: "threads", Aliases: []string{"t"}, Usage: "Adjust threads", Value: 10},
		&cli.BoolFlag{Destination: &opt.DisableAutoUpdate, Name: "disableupdate", Usage: "Disable auto update"},
	}

	// flgrecon := []cli.Flag{}
	// cmd := []*cli.Command{
	// 	{
	// 		Name:    "recon",
	// 		Aliases: []string{"r"},
	// 		Usage:   "run XorHunt with Recon mode using given template",
	// 		Flags:   flgrecon,
	// 		Action: func(ctx *cli.Context) error {
	// 			fmt.Println("recon mode")

	// 			return nil
	// 		},
	// 	},
	// 	{
	// 		Name:    "scan",
	// 		Aliases: []string{"s"},
	// 		Usage:   "run XorHunt with scanning vulnerability mode using given template",
	// 	},
	// }
	app.Flags = flg
	// app.Commands = cmd
	return opt
}

// Validate a options
func (Opts *GlobalOptions) ValOpts() {
	if !Opts.DisableAutoUpdate {
		u := Updater()
		u()
	}
}

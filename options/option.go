package options

import (
	"github.com/urfave/cli/v2"
)

type App *cli.App

type GlobalOptions struct {
	TemplateName string // Use a template file .yaml
	IterNum      int    // Set an iterator number for templates to be value in the path
	IterFile     string // set an iterator file for templates to be value in the path
	Verbose      bool   // Verbose logging
	Threads      int    // Threads controls the number of threads to use for active enumerations
	Timeout      int    // Timeout is the seconds to wait for sources to respond
	Config       string // Config contains the location of the config file
	Proxy        string // HTTP proxy

	RateLimit         int  // Global maximum number of HTTP requests to send per second
	DisableAutoUpdate bool // DisableUpdateCheck disable update checking
}

func GetOpts(app *cli.App) *GlobalOptions {
	opt := &GlobalOptions{}
	flg := []cli.Flag{

		&cli.StringFlag{Destination: &opt.TemplateName, Name: "template", Usage: "Choose a Template file e.g (templates/autorecon.yaml)", Required: true},
		&cli.IntFlag{Destination: &opt.IterNum, Name: "inum", Usage: "Set an iterator number for templates to be value in the path"},
		&cli.StringFlag{Destination: &opt.IterFile, Name: "ilist", Usage: "set an iterator file for templates to be value in the path e.g for ip list or domain lists file"},
		&cli.BoolFlag{Destination: &opt.Verbose, Name: "verbose", Usage: "Verbose logging", Value: false},
		&cli.IntFlag{Destination: &opt.Threads, Name: "threads", Usage: "Adjust threads", Value: 10},
		&cli.BoolFlag{Destination: &opt.DisableAutoUpdate, Name: "disableupdate", Usage: "Disable auto update"},
	}
	app.Flags = flg
	return opt
}

// Validate a options
func (Opts *GlobalOptions) ValOpts() {
	if !Opts.DisableAutoUpdate {
		u := Updater()
		u()
	}
}

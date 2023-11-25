package options

import "github.com/t101804/xorhunt/pkg/logger"

const banner = `
  ═╗ ╦┌─┐┬─┐╦ ╦┬ ┬┌┐┌┌┬┐
  ╔╩╦╝│ │├┬┘╠═╣│ ││││ │  We ♥ Open Source
  ╩ ╚═└─┘┴└─╩ ╩└─┘┘└┘ ┴ 
`
const ToolName = `xorhunt`
const version = `v2.0.1-dev`
const ads = `Premium and fast reverse ip repcyber.com`

func printBanner() {
	logger.Print().Msgf("%s  %s\n\n", banner, version)
	logger.Ads().Msgf("%s\n\n", ads)
}

func Updater() func() {
	return func() {
		printBanner()
		// implement auto update if there any new update in the github repo
	}
}

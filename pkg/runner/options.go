package runner

import (
	"github.com/t101804/xorhunt/pkg/logger"
	"github.com/t101804/xorhunt/pkg/templateparser"
)

func (options *Runner) loadProvidersFrom(templateFileName string) {
	tmpl, err := templateparser.ReadTemplate(templateFileName)
	if err != nil {
		logger.Error().Msgf("Could not read providers from %s: %s\n", templateFileName, err)
	}
	p := tmpl.TmplStruct.ValPath(options.options.IterNum, options.options.IterFile)
	options.tmpl = tmpl
	options.tmpl.Path = p
}

package templateparser

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/t101804/xorhunt/options"
	"github.com/t101804/xorhunt/pkg/logger"
)

type Tmpl struct {
	TmplStruct *TemplateStruct
	Path       *Path
}

func (tmpl *TemplateStruct) ValPath(opt *options.GlobalOptions) *Path {
	p := &Path{}
	if strings.Contains(tmpl.Config.Path, "{iternum}") {
		if opt.IterNum == 0 {
			logger.Fatal().Msg("you using the template that using iter number but you not specify the flag iter number")
		}
		p.IterNum = opt.IterNum
	}
	if strings.Contains(tmpl.Config.Path, "{iterfile}") {
		if opt.IterFile == "" {
			logger.Fatal().Msg("you using the template that using iter file but you not specify the flag iter lists")
		}
		p.IterFile = opt.IterFile
	}
	return p
}

func ReadTemplate(opt *options.GlobalOptions) *Tmpl {
	viper.SetConfigFile(opt.TemplateName)
	if err := viper.ReadInConfig(); err != nil {

		logger.Fatal().Msgf("error '%s' %s", opt.TemplateName, err)

	}
	tmpls := &TemplateStruct{
		Title:       viper.GetString("title"),
		Author:      viper.GetString("author"),
		LastUpdated: viper.GetString("last_updated"),
		Config: struct {
			Site         string      `yaml:"site"`
			Path         string      `yaml:"path"`
			CustomHeader []string    `yaml:"custom_header"`
			Regex        interface{} `yaml:"regex"`
			JSONParse    interface{} `yaml:"json_parse"`
		}{
			Site:         viper.GetString("config.site"),
			Path:         viper.GetString("config.path"),
			CustomHeader: viper.GetStringSlice("config.custom_header"),
			Regex:        viper.Get("config.regex"),
			JSONParse:    viper.Get("config.json_parse"),
		},
	}

	return &Tmpl{
		TmplStruct: tmpls,
		Path:       tmpls.ValPath(opt),
	}
}

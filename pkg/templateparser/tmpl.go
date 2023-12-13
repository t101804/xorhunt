package templateparser

import (
	"fmt"
	"strings"

	fileutil "github.com/projectdiscovery/utils/file"
	"github.com/spf13/viper"

	"github.com/t101804/xorhunt/pkg/logger"
)

type Tmpl struct {
	TmplStruct *TemplateStruct
	Path       *Path
}

func (tmpl *TemplateStruct) ValPath(iterNum int, iterFile string) *Path {
	p := &Path{}
	if strings.Contains(tmpl.Config.Path, "{iternum}") {
		if iterNum == 0 {
			logger.Fatal().Msg("you using the template that using iter number but you not specify the flag iter number")
		}
		p.IterNum = iterNum
	}
	if strings.Contains(tmpl.Config.Path, "{iterfile}") {
		if iterFile == "" {
			logger.Fatal().Msg("you using the template that using iter file but you not specify the flag iter lists [ --list urlisttobeiterate , e.g: --list listIPs.txt ]")
		}
		if !fileutil.FileExists(iterFile) {
			logger.Fatal().Msg("the file that you provide is not exist [ make sure u put the correct file path ]")
		}
		p.IterFile = iterFile
	}
	return p
}

func ReadTemplate(tmplFile string) (*Tmpl, error) {
	viper.SetConfigFile(tmplFile)
	if err := viper.ReadInConfig(); err != nil {

		return nil, fmt.Errorf("error '%s' %s", tmplFile, err)

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
	}, nil
}

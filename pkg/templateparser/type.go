package templateparser

type TemplateStruct struct {
	Title       string `yaml:"title"`
	Author      string `yaml:"author"`
	LastUpdated string `yaml:"last_updated"`
	Config      struct {
		Site string `yaml:"site"`
		Path string `yaml:"path"`

		CustomHeader []string    `yaml:"custom_header"`
		Regex        interface{} `yaml:"regex"`
		JSONParse    interface{} `yaml:"json_parse"`
	} `yaml:"config"`
}

type Path struct {
	IterNum  int
	IterFile string
}

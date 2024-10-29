package colors

type MTRSystemMapColors struct {
	ColorNames          map[string]Color     `yaml:"color_names"`
	LinesColors         map[string]ColorItem `yaml:"lines_colors"`
	MiscellaneousColors map[string]ColorItem `yaml:"miscellaneous_colors"`
}

type RGB [3]uint8

type Color struct {
	Name    string `yaml:"name"`
	Pantone string `yaml:"pantone,omitempty"`
	RGB     `yaml:"rgb"`
}

type ColorItem struct {
	FullName string `yaml:"full_name"`
	ColorRef `yaml:"color"`
}

type ColorRef struct {
	Ref string `yaml:"$ref"`
	*Color
}

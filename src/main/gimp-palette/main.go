package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

var dataDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	curDir := filepath.Dir(file)

	dataDir = filepath.Join(curDir, "../../../data")
}

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

func main() {
	dataFile := filepath.Join(dataDir, "colors/mtr-system-map.yaml")

	yamlData, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	var data MTRSystemMapColors
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Printf("%+v\n", data)
}

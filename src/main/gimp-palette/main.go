package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gbarail/mtr-colors/colors"
	"gopkg.in/yaml.v3"
)

var dataDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	curDir := filepath.Dir(file)

	dataDir = filepath.Join(curDir, "../../../data")
}

func main() {
	dataFile := filepath.Join(dataDir, "colors/mtr-system-map.yaml")

	yamlData, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	var data colors.MTRSystemMapColors
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Printf("%+v\n", data)
}

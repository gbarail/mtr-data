package main

import (
	"path"
	"runtime"

	"github.com/gbarail/mtr-colors/types"
)

var (
	_, file, _, _ = runtime.Caller(0)
	dataDir       = path.Join(path.Dir(file), "..", "data", "colors")
)

var (
	MTRLogoYAMLFile      = path.Join(dataDir, "mtr-logo.yaml")
	MTRSystemMapYAMLFile = path.Join(dataDir, "mtr-system-map.yaml")
)

const GIMPPaletteExtension = ".gpl"

func ReadMTRLogoYAMLFile() (types.MTRLogoData, error) {
	data, err := ReadAndUnmarshalYAML[types.MTRLogoData](MTRLogoYAMLFile)
	if err != nil {
		return nil, err
	}
	return *data, nil
}

func ReadMTRSystemMapYAMLFile() (*types.MTRSystemMapData, error) {
	data, err := ReadAndUnmarshalYAML[types.MTRSystemMapData](MTRSystemMapYAMLFile)
	if err != nil {
		return nil, err
	}
	return data, nil
}

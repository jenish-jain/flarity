package files

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/jenish-jain/logger"
)

func Read(path string) []byte {
	logger.Debug("Reading file from path %s", path)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func Write(path string, data interface{}) {
	logger.Debug("Writing file to path", "path", path)
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

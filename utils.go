package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func getProjectRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	return basePath
}

func ReadInputFile(filename string) string {
	rootDir := getProjectRootDir()
	data, err := os.ReadFile(rootDir + "/input/" + filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

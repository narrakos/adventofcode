package utils

import (
	"path/filepath"
	"runtime"
)

func GetProjectRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	return basePath
}

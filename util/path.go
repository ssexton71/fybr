package util

import (
	"errors"
	"os"
	"strings"
)

type Path struct {
	Path string
}

func NewPath(path string) *Path { return &Path{Path: path} }

func (path *Path) IsHttp() bool {
	return strings.HasPrefix(path.Path, "http://") ||
		strings.HasPrefix(path.Path, "https://")
}

func (path *Path) ReadFile() ([]byte, error) {
	return os.ReadFile(path.Path)
}

func (path *Path) ReadHttp() ([]byte, error) {
	return nil, errors.ErrUnsupported
}

func (path *Path) ReadData() ([]byte, error) {
	if path.IsHttp() {
		return path.ReadHttp()
	} else {
		return path.ReadFile()
	}
}

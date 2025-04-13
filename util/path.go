package util

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)

type Path struct {
	Path     string
	Progress func(int)
}

func (path *Path) IsHttp() bool {
	return strings.HasPrefix(path.Path, "http://") ||
		strings.HasPrefix(path.Path, "https://")
}

func SlurpAll(rdr io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(rdr)
	return buf.Bytes(), err
}

func Slurp(rdr io.Reader, progress func(int)) ([]byte, error) {
	if progress == nil {
		return SlurpAll(rdr)
	}
	const bufsz = 64 * 1024
	out := []byte{}
	buf := make([]byte, bufsz)
	for {
		n, err := rdr.Read(buf)
		if n == bufsz {
			out = slices.Concat(out, buf)
		} else if n > 0 {
			out = slices.Concat(out, buf[:n])
		}
		progress(len(out))
		if err == io.EOF {
			break
		}
		if err != nil {
			return out, err
		}
	}
	return out, nil
}

func (path *Path) ReadFile() ([]byte, error) {
	if path.Progress == nil {
		return os.ReadFile(path.Path)
	}
	rdr, err := os.Open(path.Path)
	if err != nil {
		return nil, err
	}
	return Slurp(rdr, path.Progress)
}

func (path *Path) ReadHttp() ([]byte, error) {
	client := http.Client{}

	resp, err := client.Get(path.Path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return Slurp(resp.Body, path.Progress)
}

func (path *Path) ReadData() ([]byte, error) {
	if path.IsHttp() {
		return path.ReadHttp()
	} else {
		return path.ReadFile()
	}
}

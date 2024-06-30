package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type circuitKey int

const (
	WASM circuitKey = iota
	ZKEY
)

func ReadFileByPath(basePath string, fileName string) ([]byte, error) {
	path := filepath.Join(basePath, fileName)
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s' by path '%s': %w", fileName, path, err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s' by path '%s': %w", fileName, path, err)
	}

	return data, nil
}

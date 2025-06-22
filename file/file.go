package file

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// FileSuiteBasic
type FileSuiteBasic struct{}

func New() *FileSuiteBasic {
	return new(FileSuiteBasic)
}

// Mkdir create a folder
func (fs *FileSuiteBasic) Mkdir(folderPath string) (string, error) {
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return "", err
	}
	return folderPath, nil
}

// WriteFile Write data to a file
func (fs *FileSuiteBasic) WriteFile(path string, data []byte) error {
	const fileMode = 0644
	return os.WriteFile(path, data, fileMode)
}

// ReadFile Read data from a file
func (fs *FileSuiteBasic) ReadFile(f string) ([]byte, error) {
	return os.ReadFile(f)
}

// IsExist Check if the path exists
func (fs *FileSuiteBasic) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Joinpath join the root path and all sub
func (fs *FileSuiteBasic) JoinPath(root string, elems ...string) string {
	return filepath.Join(append([]string{root}, elems...)...)
}

// GetFileData get data direct
func (fs *FileSuiteBasic) ReadFileAt(root string, elems ...string) ([]byte, error) {
	path := fs.JoinPath(root, elems...)
	if !fs.Exists(path) {
		return nil, fmt.Errorf("file not found: %s", path)
	}
	return os.ReadFile(path)
}

// RootPath Get project root path
//
//	GOTMPDIR required
//
//	subPath: join sub path
func (fs *FileSuiteBasic) RootPath(subPath ...string) (path string, err error) {
	// go build path
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	runPath, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	if err != nil {
		return "", err
	}

	// go run path
	buildPath, err := filepath.EvalSymlinks(os.Getenv("GOTMPDIR"))
	if err != nil {
		return "", err
	}

	// go run cache (go1.24)
	cachePath, err := filepath.EvalSymlinks(os.Getenv("GOCACHE"))
	if err != nil {
		return "", err
	}

	if strings.Contains(runPath, buildPath) || strings.Contains(runPath, cachePath) {
		var absPath string
		_, filename, _, ok := runtime.Caller(1)
		if ok {
			absPath = filepath.Dir(filepath.Dir(filename))
			fullPaths := []string{absPath}
			fullPaths = append(fullPaths, subPath...)
			fullPath := filepath.Join(fullPaths...)
			return fullPath, nil
		}
	}
	fullPaths := []string{runPath}
	fullPaths = append(fullPaths, subPath...)
	fullPath := filepath.Join(fullPaths...)
	return fullPath, nil
}

package minitools

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// FileSuiteBasic
type FileSuiteBasic struct{}

// Mkdir create a folder
func (fs *FileSuiteBasic) Mkdir(folderPath string) (string, error) {
	_, err := os.Stat(folderPath)
	if err == nil {
		return folderPath, nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			return "", err
		}
		return folderPath, nil
	}
	return folderPath, nil
}

// WriteFile Write data to a file
func (fs *FileSuiteBasic) WriteFile(f string, data []byte) error {
	return os.WriteFile(f, data, 0666)
}

// ReadFile Read data from a file
func (fs *FileSuiteBasic) ReadFile(f string) ([]byte, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CheckType check var type
func (fs *FileSuiteBasic) CheckType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

// IsExist Check if the path exists
func (fs *FileSuiteBasic) IsExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
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

	if strings.Contains(runPath, buildPath) {
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

// Joinpath join the root path and all sub
func (fs *FileSuiteBasic) Joinpath(root string, sub ...string) (string, error) {
	mainRoot, err := fs.RootPath(root)
	if err != nil {
		return "", err
	}
	fullpath := append([]string{mainRoot}, sub...)
	return filepath.Join(fullpath...), nil
}

// GetFileData get data direct
func (fs *FileSuiteBasic) GetFileData(root string, sub ...string) ([]byte, error) {
	fullPath, err := fs.Joinpath(root, sub...)
	if err != nil {
		return nil, err
	}
	err = fs.IsExist(fullPath)
	if err != nil {
		return nil, err
	}
	data, err := fs.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

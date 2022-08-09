package minitools

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// FileSuiteBasic 文件操作基类
type FileSuiteBasic struct{}

// Create 创建文件夹
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

// Write 保存文件
func (fs *FileSuiteBasic) WriteFile(f string, data []byte) error {
	return os.WriteFile(f, data, 0666)
}

// Read 读取文件
func (fs *FileSuiteBasic) ReadFile(f string) ([]byte, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CheckType 检查文件类型
func (fs *FileSuiteBasic) CheckType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

// CheckExist 检查文件是否存在
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

// RootPath 获取当前项目根目录 main()
//
//	GOTMPDIR 是必须的
//
//	subPath: 拼接子目录
func (fs *FileSuiteBasic) RootPath(subPath ...string) (path string, err error) {
	// go build 可执行文件路径
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	runPath, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	if err != nil {
		return "", err
	}
	// go run 调试路径
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

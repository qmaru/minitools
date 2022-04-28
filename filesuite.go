package minitools

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

// FileSuiteBasic 文件操作基类
type FileSuiteBasic struct{}

// Create 创建文件夹
func (fs *FileSuiteBasic) Create(folderPath string) (string, error) {
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

// Write 写入文件
func (fs *FileSuiteBasic) Write(f string, data []byte) error {
	return ioutil.WriteFile(f, data, 0666)
}

// Read 读取文件
func (fs *FileSuiteBasic) Read(f string) ([]byte, error) {
	data, err := ioutil.ReadFile(f)
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
func (fs *FileSuiteBasic) CheckExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}

// LocalPath 获取当前路径
//	debug: true / false
//	true: 获取 go run 的路径
//	false: 获取可执行文件的路径
func (fs *FileSuiteBasic) LocalPath(debug bool) (path string, err error) {
	if debug {
		path, err = os.Getwd()
		if err != nil {
			return "", err
		}
	} else {
		path, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

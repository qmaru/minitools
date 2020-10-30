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
func (fs *FileSuiteBasic) Create(folderPath string) string {
	_, err := os.Stat(folderPath)
	if err == nil {
		return folderPath
	}
	if os.IsNotExist(err) {
		_ = os.Mkdir(folderPath, os.ModePerm)
		return folderPath
	}
	return folderPath
}

// Write 写入文件
func (fs *FileSuiteBasic) Write(f string, data []byte) {
	err := ioutil.WriteFile(f, data, 0666)
	if err != nil {
		panic(err)
	}
}

// Read 读取文件
func (fs *FileSuiteBasic) Read(f string) []byte {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return data
}

// CheckType 检查文件类型
func (fs *FileSuiteBasic) CheckType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

// CheckExist 检查文件是否存在
func (fs *FileSuiteBasic) CheckExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// LocalPath 获取当前路径
//	debug: true / false
//	true: 获取 go run 的路径
//	false: 获取可执行文件的路径
func (fs *FileSuiteBasic) LocalPath(debug bool) (path string) {
	var err error
	if debug {
		path, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		path, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
	}
	return path
}

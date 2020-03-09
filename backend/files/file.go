package files

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

// 调用os.MkdirAll递归创建文件夹
func CreateFile(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

//  判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
 * 重置配置文件路径
 * 该方法目前仅用于测试 ，有些多余
*/
func GetAbsPath(confPath string) string {
	getwd, err := os.Getwd() // os.Getwd()获取当前程序执行路径,可以理解为exe所在的路径
	if err != nil {
		color.Red(fmt.Sprintf("Getwd err %v", err))
	}

	end := filepath.Base(getwd) // filepath.Base() returns the last element of path.
	if end != "backend" { // 项目运行目录
		return filepath.Join(filepath.Dir(getwd), confPath)
	}

	return confPath
}
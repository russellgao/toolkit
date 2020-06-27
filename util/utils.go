package util

import (
	"fmt"
	"io/ioutil"
)

// 遍历文件夹获取所有文件
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

// 读取文件
func ReadFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	return string(content), err
}

// 写入文件
func WriteToFile(filename, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	return err
}

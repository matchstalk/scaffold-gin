package tools

import (
	"io/ioutil"
	"os"
)

func GetPathFiles(path string) (fileNames []string, err error) {
	var files []os.FileInfo
	files, err = ioutil.ReadDir(path)
	if err != nil {
		return
	}
	fileNames = make([]string, 0)
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return
}

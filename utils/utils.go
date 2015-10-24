package utils

import (
	"io/ioutil"
	"path"
	"runtime"
)

func GetVersion() string {
	_, filename, _, _ := runtime.Caller(1)
	filePath := path.Join(filename, "../VERSION")
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(f)
}

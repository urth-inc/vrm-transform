package fileUtil

import (
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) (buf []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

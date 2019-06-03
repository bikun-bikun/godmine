package file

import (
	"io/ioutil"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Save(path string, data []byte) error {
	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}

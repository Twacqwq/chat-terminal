package utils

import (
	"os"
)

func Mkdir(path string) error {
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil || os.IsNotExist(err) {
			return err
		}
	}

	return nil
}

package utils

import (
	"io/ioutil"
)

func readFile(f string) (string, error) {
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

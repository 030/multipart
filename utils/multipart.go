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

func multipart(f ...string) error {
	for _, v := range f {
		_, err := ioutil.ReadFile(v)
		if err != nil {
			return err
		}
	}
	return nil
}

package utils

import (
	"io/ioutil"
	"strings"
)

func readFile(f string) (string, error) {
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func writeField(s string) string {
	parts := strings.Split(s, "=")
	return parts[0] + " " + parts[1]
}

func multipartUpload(f ...string) error {
	for _, v := range f {
		if strings.Contains(v, "=@") {
			parts := strings.Split(v, "=@")
			_, err := ioutil.ReadFile(parts[1])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
package utils

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func StringToSlice(s string) []string {
	a := strings.Split(s, ",")

	return a
}

func split(s string, el string) (string, string) {
	parts := strings.Split(s, el)
	k := parts[0]
	v := parts[1]

	return k, v
}

func readFile(f string) ([]byte, error) {
	b, err := ioutil.ReadFile(filepath.Clean(f))
	if err != nil {
		return nil, err
	}

	return b, nil
}

func addFileToWriter(w *multipart.Writer, b []byte, fn, f string) error {
	part, err := w.CreateFormFile(fn, f)
	if err != nil {
		return err
	}

	if _, err = part.Write(b); err != nil {
		return err
	}

	return nil
}

func metadataAndFile(w *multipart.Writer, s string) error {
	k, v := split(s, "=@")
	b, err := readFile(v)
	if err != nil {
		return err
	}

	if err = addFileToWriter(w, b, k, v); err != nil {
		return err
	}

	return nil
}

func addKeyValueToWriter(w *multipart.Writer, k, v string) error {
	if err := w.WriteField(k, v); err != nil {
		return err
	}

	return nil
}

func metadataAndExtension(w *multipart.Writer, s string) error {
	k, v := split(s, "=")
	if err := addKeyValueToWriter(w, k, v); err != nil {
		return err
	}

	return nil
}

func Body(w *multipart.Writer, s []string) error {
	for _, val := range s {
		switch {
		case strings.Contains(val, "=@"):
			err := metadataAndFile(w, val)
			if err != nil {
				return err
			}
		case strings.Contains(val, "="):
			err := metadataAndExtension(w, val)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("the string should at least contain a '=', but was: '%v'", val)
		}
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var body = new(bytes.Buffer)
var writer = multipart.NewWriter(body)

// Upload struct
type Upload struct {
	URL      string
	Username string
	Password string
}

// MultipartUpload splits the string into a slice, created a multipart
// and that is posted to an URL
func (u Upload) MultipartUpload(s string) error {
	args := strings.Split(s, ",")
	err := multipartBody(args...)
	if err != nil {
		return err
	}
	err2 := u.upload()
	if err2 != nil {
		return err2
	}
	return nil
}

func addFileToWriter(b []byte, fn, f string) error {
	part, err := writer.CreateFormFile(fn, f)
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := part.Write(b)
	if err2 != nil {
		return err2
	}
	return nil
}

func addKeyValueToWriter(k, v string) error {
	err := writer.WriteField(k, v)
	if err != nil {
		return err
	}
	return nil
}

func readFile(f string) ([]byte, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func writeField(s string) string {
	parts := strings.Split(s, "=")
	return parts[0] + " " + parts[1]
}

func multipartBody(f ...string) error {
	log.Debug("The input string: ", f)
	for _, v := range f {
		log.Debug("The elements that reside in the input string: ", v)

		if strings.Contains(v, "=@") {
			parts := strings.Split(v, "=@")
			b, err := ioutil.ReadFile(parts[1])
			if err != nil {
				return err
			}
			addFileToWriter(b, parts[0], parts[1])
		} else {
			parts := strings.Split(v, "=")
			err := addKeyValueToWriter(parts[0], parts[1])
			if err != nil {
				return err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return err
	}

	return nil
}

func (u Upload) upload() error {
	req, err := http.NewRequest("POST", u.URL, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(u.Username, u.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) || (err != nil) {
		return fmt.Errorf("HTTPStatusCode: '%d'; ResponseMessage: '%s'; ErrorMessage: '%v'", resp.StatusCode, string(b), err)
	}
	return nil
}

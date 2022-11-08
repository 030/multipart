package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

var body = new(bytes.Buffer)
var writer = multipart.NewWriter(body)

// Upload struct
type Upload struct {
	URL      string
	Username string
	Password string
}

// MultipartUpload splits the string into a slice, creates a multipart
// and that is posted to an URL
func (u Upload) MultipartUpload(s string) error {
	a := stringToSlice(s)
	if err := multipartBody(a); err != nil {
		return err
	}

	if err := u.upload(); err != nil {
		return err
	}

	return nil
}

func stringToSlice(s string) []string {
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

func addFileToWriter(b []byte, fn, f string) error {
	part, err := writer.CreateFormFile(fn, f)
	if err != nil {
		return err
	}

	if _, err = part.Write(b); err != nil {
		return err
	}

	return nil
}

func metadataAndFile(s string) error {
	k, v := split(s, "=@")
	b, err := readFile(v)
	if err != nil {
		return err
	}

	if err = addFileToWriter(b, k, v); err != nil {
		return err
	}

	return nil
}

func addKeyValueToWriter(k, v string) error {
	if err := writer.WriteField(k, v); err != nil {
		return err
	}

	return nil
}

func metadataAndExtension(s string) error {
	k, v := split(s, "=")
	err := addKeyValueToWriter(k, v)
	if err != nil {
		return err
	}
	return nil
}

func multipartBody(s []string) error {
	for _, val := range s {
		if strings.Contains(val, "=@") {
			err := metadataAndFile(val)
			if err != nil {
				return err
			}
		} else if strings.Contains(val, "=") {
			err := metadataAndExtension(val)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("the string should at least contain a '=', but was: '%v'", val)
		}
	}

	err := writer.Close()
	if err != nil {
		return err
	}

	return nil
}

func (u Upload) uploadRequest() (*http.Request, error) {
	req, err := http.NewRequest("POST", u.URL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(u.Username, u.Password)
	return req, nil
}

func uploadResponse(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u Upload) upload() error {
	req, err := u.uploadRequest()
	if err != nil {
		return err
	}

	resp, err := uploadResponse(req)
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

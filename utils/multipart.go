package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var body = new(bytes.Buffer)
var writer = multipart.NewWriter(body)

type upload struct {
	url      string
	username string
	password string
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

// create the body of the multipart by writing files and key-values
func multipartBody(f ...string) (*bytes.Buffer, error) {
	for _, v := range f {
		if strings.Contains(v, "=@") {
			parts := strings.Split(v, "=@")
			b, err := ioutil.ReadFile(parts[1])
			if err != nil {
				return nil, err
			}
			addFileToWriter(b, parts[0], parts[1])
		} else {
			parts := strings.Split(v, "=")
			err := addKeyValueToWriter(parts[0], parts[1])
			if err != nil {
				return nil, err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	if body.String() == "" {
		return nil, errors.New("Body should not be empty")
	}

	return body, nil
}

func (u upload) multipartUpload() error {
	//"http://localhost:9999/service/rest/v1/components?repository=maven-releases"
	req, err := http.NewRequest("POST", u.url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(u.username, u.password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		log.WithFields(log.Fields{
			"Error":          err,
			"Message":        string(b),
			"HTTPStatusCode": resp.StatusCode,
		}).Fatal("Multipart upload failed")
	}
	return nil
}

package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

var body = new(bytes.Buffer)
var writer = multipart.NewWriter(body)

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

func multipartUpload(f ...string) error {
	for _, v := range f {
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

	// To
	// Be
	// Refactored
	err := writer.Close()
	if err != nil {
		log.Fatal(err)
	}
	// log.Fatal(body.String())
	req, err := http.NewRequest("POST", "http://localhost:9999/service/rest/v1/components?repository=maven-releases", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	req.SetBasicAuth("admin", "admin123")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	message, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(string(message))
		log.Fatal(resp.StatusCode)
	}

	return nil
}

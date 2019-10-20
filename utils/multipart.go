package utils

import (
	"bytes"
	"fmt"
	"mime/multipart"
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

// // MultipartUpload splits the string into a slice, creates a multipart
// // and that is posted to an URL
// func (u Upload) MultipartUpload(s string) error {
// 	args := strings.Split(s, ",")
// 	err := multipartBody(args...)
// 	if err != nil {
// 		return err
// 	}
// 	err = u.upload()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func stringToSlice(s string) []string {
	args := strings.Split(s, ",")
	return args
}

func multipartBody(s []string) (string, string, error) {
	var metadata, file string
	var parts []string
	for _, v := range s {
		if strings.Contains(v, "=@") {
			parts = strings.Split(v, "=@")
			metadata = parts[0]
			file = parts[1]
		} else if strings.Contains(v, "=") {
			parts = strings.Split(v, "=")
			metadata = parts[0]
			file = parts[1]
		} else {
			return "", "", fmt.Errorf("The string should at least contain a '=', but was: '%v'", parts)
		}
	}
	return metadata, file, nil
}

// func addFileToWriter(b []byte, fn, f string) error {
// 	part, err := writer.CreateFormFile(fn, f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = part.Write(b)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func addKeyValueToWriter(k, v string) error {
// 	err := writer.WriteField(k, v)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func readFile(f string) ([]byte, error) {
// 	b, err := ioutil.ReadFile(f)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return b, nil
// }

// // func writeField(s string) (string, error) {
// // 	parts := strings.Split(s, "=")
// // 	if len(parts) == 1 {
// // 		return "", fmt.Errorf("The string should at least contain a '=', but was: '%v'", parts)
// // 	}
// // 	return parts[0] + " " + parts[1], nil
// // }

// func multipartBody(f ...string) error {
// 	fmt.Println("CP2")
// 	log.Debug("The input string: ", f)
// 	for _, v := range f {
// 		log.Debug("The elements that reside in the input string: ", v)

// 		if strings.Contains(v, "=@") {
// 			fmt.Println("CP3")
// 			parts := strings.Split(v, "=@")
// 			fmt.Println("CP3a")
// 			b, err := ioutil.ReadFile(parts[1])
// 			fmt.Println("CP3b")
// 			if err != nil {
// 				return err
// 			}
// 			fmt.Println("CP3c")
// 			addFileToWriter(b, parts[0], parts[1])
// 			fmt.Println("CP3d")
// 		} else {
// 			fmt.Println("CP4")
// 			fmt.Println(v)
// 			parts := strings.Split(v, "=")
// 			fmt.Println(len(parts))
// 			if len(parts) == 1 {
// 				return fmt.Errorf("The string should at least contain a '=', but was: '%v'", parts)
// 			}
// 			fmt.Println("CP4b")

// 			err := addKeyValueToWriter(parts[0], parts[1])
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	err := writer.Close()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (u Upload) upload() error {
// 	req, err := http.NewRequest("POST", u.URL, body)
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Add("Content-Type", writer.FormDataContentType())
// 	req.SetBasicAuth(u.Username, u.Password)

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	b, err := ioutil.ReadAll(resp.Body)
// 	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) || (err != nil) {
// 		return fmt.Errorf("HTTPStatusCode: '%d'; ResponseMessage: '%s'; ErrorMessage: '%v'", resp.StatusCode, string(b), err)
// 	}
// 	return nil
// }

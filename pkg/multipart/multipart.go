package multipart

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/030/multipart/internal/app/multipart/utils"
)

// Upload struct.
type Upload struct {
	URL      string
	Username string
	Password string
}

var (
	body   = new(bytes.Buffer)
	writer = multipart.NewWriter(body)
)

func (u Upload) request() (*http.Request, error) {
	req, err := http.NewRequest("POST", u.URL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(u.Username, u.Password)

	return req, nil
}

func response(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// MultipartUpload splits the string into a slice, creates a multipart
// and that is posted to an URL.
func (u Upload) Upload(s string) error {
	a := utils.StringToSlice(s)
	if err := utils.Body(writer, a); err != nil {
		return fmt.Errorf("%w", err)
	}

	req, err := u.request()
	if err != nil {
		return err
	}

	resp, err := response(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) || (err != nil) {
		return fmt.Errorf("HTTPStatusCode: '%d'; ResponseMessage: '%s'; ErrorMessage: '%w'", resp.StatusCode, string(b), err)
	}

	return nil
}

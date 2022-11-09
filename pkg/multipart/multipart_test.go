package multipart

import (
	"os"
	"testing"

	"github.com/030/mij"
	"github.com/030/multipart/internal/app/multipart/multiparttest"
)

const expPostErr string = "Post \"\": unsupported protocol scheme \"\""

func TestMain(m *testing.M) {
	containers := []mij.DockerImage{multiparttest.Image(10000)}
	if err := multiparttest.Setup(containers); err != nil {
		panic(err)
	}

	code := m.Run()
	if err := multiparttest.Shutdown(containers); err != nil {
		panic(err)
	}

	os.Exit(code)
}

var u = Upload{URL: "http://localhost:10000/service/rest/v1/components?repository=maven-releases", Username: "admin", Password: "admin123"}

func TestUpload(t *testing.T) {
	err := u.Upload("maven2.asset1=@../../test/testdata/file2.pom,maven2.asset1.extension=pom")
	if err != nil {
		t.Errorf("Error should be 'nil', but was '%v'", err)
	}
}

func TestUploadFail(t *testing.T) {
	err := u.Upload("hello")
	want := "the string should at least contain a '=', but was: 'hello'"
	if err != nil {
		if err.Error() != want {
			t.Errorf("An error was expected. Got: '%v', want: '%s'", err, want)
		}
	}

	u = Upload{URL: "", Username: "", Password: ""}
	err = u.Upload("maven2.asset1.extension=pom")
	want = expPostErr
	if err != nil {
		if err.Error() != want {
			t.Errorf("An error was expected. Got: '%v', want: '%s'", err, want)
		}
	}
}

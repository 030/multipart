package utils

import (
	"os"
	"reflect"
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

func TestMultipartUpload(t *testing.T) {
	u := Upload{URL: "http://localhost:10000/service/rest/v1/components?repository=maven-releases", Username: "admin", Password: "admin123"}

	// happy
	err := u.MultipartUpload("maven2.asset1=@../../../../test/testdata/file2.pom,maven2.asset1.extension=pom")
	if err != nil {
		t.Errorf("Error should be 'nil', but was '%v'", err)
	}

	// unhappy
	err = u.MultipartUpload("hello")
	want := "the string should at least contain a '=', but was: 'hello'"
	if err != nil {
		if err.Error() != want {
			t.Errorf("An error was expected. Got: '%v', want: '%s'", err, want)
		}
	}

	// unhappy - upload should fail if endpoint is incorrect
	u = Upload{URL: "", Username: "", Password: ""}
	err = u.MultipartUpload("maven2.asset1.extension=pom")
	want = expPostErr
	if err != nil {
		if err.Error() != want {
			t.Errorf("An error was expected. Got: '%v', want: '%s'", err, want)
		}
	}
}

func TestUpload(t *testing.T) {
	u := Upload{URL: "", Username: "", Password: ""}
	err := u.upload()
	want := expPostErr
	if err.Error() != want {
		t.Errorf("An error was expected. Received '%v'; want '%s'", err, want)
	}
}

func TestStringToSlice(t *testing.T) {
	testMap := map[string][]string{
		"":                {""},
		"hello,world":     {"hello", "world"},
		"hello,world,ola": {"hello", "world", "ola"},
	}

	for k, v := range testMap {
		got := stringToSlice(k)
		want := v
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Slice not identical. Expected %s, but was %s.", want, got)
		}
	}
}

func TestReadFile(t *testing.T) {
	_, err := readFile("abc")
	want := "open abc: no such file or directory"
	if err.Error() != want {
		t.Errorf("An error was expected. Got '%v', want '%s'", err, want)
	}
}

func TestMultipartBody(t *testing.T) {
	in := []string{"this is invalid input"}

	err := multipartBody(in)
	want := "the string should at least contain a '=', but was: 'this is invalid input'"

	if err.Error() != want {
		t.Errorf("An error was expected. Received '%v'; want '%s'", err, want)
	}
}

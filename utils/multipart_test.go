package utils

import (
	"reflect"
	"testing"
)

const expPostErr string = "Post \"\": unsupported protocol scheme \"\""

func TestMultipartUpload(t *testing.T) {
	u := Upload{URL: "http://localhost:9999/service/rest/v1/components?repository=maven-releases", Username: "admin", Password: "admin123"}

	// happy
	err := u.MultipartUpload("maven2.asset1=@test-files-multipart/file2.pom,maven2.asset1.extension=pom")
	if err != nil {
		t.Errorf("Error should be 'nil', but was '%v'", err)
	}

	// unhappy
	err = u.MultipartUpload("hello")
	want := "The string should at least contain a '=', but was: 'hello'"
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
		"":                []string{""},
		"hello,world":     []string{"hello", "world"},
		"hello,world,ola": []string{"hello", "world", "ola"},
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
	want := "The string should at least contain a '=', but was: 'this is invalid input'"

	if err.Error() != want {
		t.Errorf("An error was expected. Received '%v'; want '%s'", err, want)
	}
}

package utils

import (
	"reflect"
	"testing"
)

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

func TestMultipartBody(t *testing.T) {
	// happy 1 - metadata and file
	in := []string{"maven2.asset1=@test-files-multipart/file1.pom"}

	got1, got2, _ := multipartBody(in)
	want1 := "maven2.asset1"
	want2 := "test-files-multipart/file1.pom"

	if got1 != want1 || got2 != want2 {
		t.Errorf("Got: '%s' && '%s', want '%s' && '%s'", got1, got2, want1, want2)
	}

	// happy 2 - metadata and extension
	in = []string{"maven2.asset1.extension=pom"}

	got1, got2, _ = multipartBody(in)
	want1 = "maven2.asset1.extension"
	want2 = "pom"

	if got1 != want1 || got2 != want2 {
		t.Errorf("Got: '%s' && '%s', want '%s' && '%s'", got1, got2, want1, want2)
	}

	// unhappy
	in = []string{"this is invalid input"}

	_, _, err := multipartBody(in)
	want := "The string should at least contain a '=', but was: '[]'"

	if err.Error() != want {
		t.Errorf("An error was expected. Received '%v'; want '%s'", err, want)
	}
}

// func TestReadFile(t *testing.T) {
// 	got, err := readFile("test-files/helloworld.txt")
// 	if err != nil {
// 		t.Errorf("Cannot read file: %v; ", err)
// 	}
// 	want := "helloworld"
// 	if string(got) != want {
// 		t.Errorf("readFile() = %s; want %s", got, want)
// 	}

// 	_, err = readFile("folder-does-not-exist/helloworld.txt")
// 	want = "open folder-does-not-exist/helloworld.txt: no such file or directory"
// 	if err.Error() != want {
// 		t.Errorf("An error was expected. Received '%v'; want '%s'", err, want)
// 	}
// }

// func TestMultipartBody(t *testing.T) {
// 	err := multipartBody("maven2.asset1=@test-files-multipart/file1.pom",
// 		"maven2.asset1.extension=pom",
// 		"maven2.asset2=@test-files-multipart/file1.jar",
// 		"maven2.asset2.extension=jar",
// 		"maven2.asset3=@test-files-multipart/file1-sources.jar",
// 		"maven2.asset3.extension=sources.jar")
// 	if err != nil {
// 		t.Errorf("Unexpected error; got '%v'", err)
// 	}

// 	err = multipartBody("something.some=@test-files-multipart/does-not-exist.txt")
// 	if err == nil {
// 		t.Errorf("Expected error, got '%v'", err)
// 	}

// 	err = multipartBody("hello,world")
// 	want := "The string should at least contain a '=', but was: '[hello,world]'"
// 	if err.Error() != want {
// 		t.Errorf("Expected error, got '%v', want: '%v'", err, want)
// 	}
// }

// // func TestWriteField(t *testing.T) {
// // 	// happy
// // 	got, _ := writeField("maven2.asset3.extension=jar")
// // 	want := "maven2.asset3.extension jar"
// // 	if got != want {
// // 		t.Errorf("Want '%v', got '%v'", want, got)
// // 	}

// // 	// unhappy
// // 	_, err := writeField("haha")
// // 	want = "The string should at least contain a '=', but was: '[haha]'"
// // 	if err.Error() != want {
// // 		t.Errorf("Expected error, got '%v', want: '%v'", err, want)
// // 	}
// // }

// func TestUpload(t *testing.T) {
// 	u := Upload{URL: "", Username: "", Password: ""}
// 	err := u.upload()
// 	want := "Post : unsupported protocol scheme \"\""
// 	if err.Error() != want {
// 		t.Errorf("An error was expected. Received '%v'; want '%s'", err, want)
// 	}

// 	u = Upload{URL: "http://releasesoftwaremoreoften.com", Username: "admin", Password: "incorrect password"}
// 	err = u.upload()
// 	want = `HTTPStatusCode: '405'; ResponseMessage: '<html>
// <head><title>405 Not Allowed</title></head>
// <body bgcolor="white">
// <center><h1>405 Not Allowed</h1></center>
// </body>
// </html>'; ErrorMessage: '<nil>'`
// 	if err.Error() != want {
// 		t.Errorf("An error was expected. Got: '%v'; want: '%s'", err, want)
// 	}
// }

// // func TestMultipartUpload(t *testing.T) {
// // 	u := Upload{URL: "", Username: "", Password: ""}
// // 	u.MultipartUpload("hello,world")
// // }

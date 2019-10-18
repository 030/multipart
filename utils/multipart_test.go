package utils

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	got, err := readFile("test-files/helloworld.txt")
	if err != nil {
		t.Errorf("Cannot read file: %v; ", err)
	}
	want := "helloworld"
	if string(got) != want {
		t.Errorf("readFile() = %s; want %s", got, want)
	}

	_, err2 := readFile("folder-does-not-exist/helloworld.txt")
	want2 := "open folder-does-not-exist/helloworld.txt: no such file or directory"
	if err2.Error() != want2 {
		t.Errorf("An error was expected. Received '%v'; want '%s'", err2, want2)
	}
}

func TestMultipart(t *testing.T) {
	err := multipartBody("maven2.asset1=@test-files-multipart/file1.pom",
		"maven2.asset1.extension=pom",
		"maven2.asset2=@test-files-multipart/file1.jar",
		"maven2.asset2.extension=jar",
		"maven2.asset3=@test-files-multipart/file1-sources.jar",
		"maven2.asset3.extension=sources.jar")
	if err != nil {
		t.Errorf("Unexpected error; got '%v'", err)
	}

	err = multipartBody("something.some=@test-files-multipart/does-not-exist.txt")
	if err == nil {
		t.Errorf("Expected error, got '%v'", err)
	}

	err = multipartBody("hello,world")
	want := "The string should at least contain a '=', but was: '[hello,world]'"
	if err.Error() != want {
		t.Errorf("Expected error, got '%v', want: '%v'", err, want)
	}
}

func TestWriteField(t *testing.T) {
	got := writeField("maven2.asset3.extension=jar")
	want := "maven2.asset3.extension jar"
	if got != want {
		t.Errorf("Want '%v', got '%v'", want, got)
	}
}

func TestUpload(t *testing.T) {
	u := Upload{URL: "", Username: "", Password: ""}
	err := u.upload()
	want := "Post : unsupported protocol scheme \"\""
	if err.Error() != want {
		t.Errorf("An error was expected. Received '%v'; want '%s'", err, want)
	}

	u = Upload{URL: "http://releasesoftwaremoreoften.com", Username: "admin", Password: "incorrect password"}
	err = u.upload()
	want = `HTTPStatusCode: '405'; ResponseMessage: '<html>
<head><title>405 Not Allowed</title></head>
<body bgcolor="white">
<center><h1>405 Not Allowed</h1></center>
</body>
</html>'; ErrorMessage: '<nil>'`
	if err.Error() != want {
		t.Errorf("An error was expected. Got: '%v'; want: '%s'", err, want)
	}
}

// func TestMultipartUpload(t *testing.T) {
// 	u := Upload{URL: "", Username: "", Password: ""}
// 	u.MultipartUpload("hello,world")
// }

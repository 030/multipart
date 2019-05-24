package utils

import "testing"

func TestReadFile(t *testing.T) {
	got, err := readFile("test-files/helloworld.txt")
	if err != nil {
		t.Errorf("Cannot read file: %v; ", err)
	}
	want := "helloworld"
	if got != want {
		t.Errorf("readFile() = %s; want %s", got, want)
	}

	_, err2 := readFile("folder-does-not-exist/helloworld.txt")
	want2 := "open folder-does-not-exist/helloworld.txt: no such file or directory"
	if err2.Error() != want2 {
		t.Errorf("An error was expected. Received '%v'; want '%s'", err2, want2)
	}
}

func TestMultipart(t *testing.T) {
	err := multipart("maven2.asset1=@test-files-multipart/1.pom",
		"maven2.asset1.extension=pom",
		"maven2.asset2=@test-files-multipart/2.jar",
		"maven2.asset2.extension=jar",
		"maven2.asset3=@test-files-multipart/3-sources.jar",
		"maven2.asset3.extension=jar")
	if err != nil {
		t.Errorf("Unexpected error; got '%v'", err)
	}

	err2 := multipart("something.some=@test-files-multipart/does-not-exist.txt")
	if err2 == nil {
		t.Errorf("Expected error, got '%v'", err2)
	}
}

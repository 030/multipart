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
	err := multipart("test-files-multipart/1.txt", "test-files-multipart/2.txt", "test-files-multipart/3.txt")
	if err != nil {
		t.Errorf("Unexpected error; got '%v'", err)
	}

	err2 := multipart("test-files-multipart/does-not-exist.txt")
	if err2 == nil {
		t.Errorf("Expected error, got '%v'", err2)
	}
}

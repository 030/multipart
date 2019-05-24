package utils

import "testing"

func TestReadFile(t *testing.T) {
	got := readFile()
	want := "helloworld"
	if got != want {
		t.Errorf("readFile() = %s; want %s", got, want)
	}
}

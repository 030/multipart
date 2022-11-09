package utils

import (
	"reflect"
	"testing"
)

func TestStringToSlice(t *testing.T) {
	testMap := map[string][]string{
		"":                {""},
		"hello,world":     {"hello", "world"},
		"hello,world,ola": {"hello", "world", "ola"},
	}

	for k, v := range testMap {
		got := StringToSlice(k)
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

	err := Body(nil, in)
	want := "the string should at least contain a '=', but was: 'this is invalid input'"

	if err.Error() != want {
		t.Errorf("An error was expected. Received '%v'; want '%s'", err, want)
	}
}

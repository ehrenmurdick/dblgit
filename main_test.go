package main

import "errors"
import "os/exec"
import "os"
import "testing"

func Test_getPairsFilePath_whenPairsFileIsInHome_thenReturnsPairsFilePath(t *testing.T) {
	getenv = func(key string) string {
		return "/home"
	}

	stat = func(path string) (os.FileInfo, error) {
		if path == "/home/.pairs" {
			return nil, nil
		}
		return nil, errors.New("file missin!")
	}

	expected := "/home/.pairs"
	actual, err := getPairsFilePath()
	if err != nil {
		t.Fatalf("unexpected error! %s", err.Error())
	}
	if expected != actual {
		t.Fatalf("expected path to be \"%s\", was \"%s\"", expected, actual)
	}
}

func Test_getPairsFilePath_whenPairsFileIsInCurrentDir_thenReturnsPairsFilePath(t *testing.T) {
	getenv = func(key string) string {
		return "/home"
	}

	stat = func(path string) (os.FileInfo, error) {
		if path == ".pairs" {
			return nil, nil
		}
		return nil, errors.New("file missin!")
	}

	expected := ".pairs"
	actual, err := getPairsFilePath()
	if err != nil {
		t.Fatalf("unexpected error! %s", err.Error())
	}
	if expected != actual {
		t.Fatalf("expected path to be \"%s\", was \"%s\"", expected, actual)
	}
}

func Test_getPairsFilePath_whenPairsIsMissing_thenReturnsError(t *testing.T) {
	getenv = func(key string) string {
		return "/home"
	}

	stat = func(path string) (os.FileInfo, error) {
		return nil, errors.New("file missin!")
	}

	expected := "No .pairs file found! Put one in your home directory or current working directory!"
	_, actual := getPairsFilePath()
	if expected != actual.Error() {
		t.Fatalf("expected path to be \"%s\", was \"%s\"", expected, actual.Error())
	}
}

type mockCmd struct {
}

func Test_readGitConfig_whenKeyIsPresent_thenReturnsValue(t *testing.T) {
	command = func(name string, args ...string) *exec.Cmd {
		return &mockCmd{}
	}
}

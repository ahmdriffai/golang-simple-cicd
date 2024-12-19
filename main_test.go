package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Hello dunia tipu-tipu sekali, berhasil cicd"
	if expected != "Hello dunia tipu-tipu sekali, berhasil cicd" {
		t.Errorf("Expected %s, got %s", expected, "Hello dunia tipu-tipu sekali, berhasil cicd")
	}
}

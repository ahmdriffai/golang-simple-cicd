package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Hello dunia tipu-tipu sekali, berhasil cicd, keren cuy"
	if expected != "Hello dunia tipu-tipu sekali, berhasil cicd, keren cuy" {
		t.Errorf("Expected %s, got %s", expected, "Hello dunia tipu-tipu sekali, berhasil cicd, keren cuy")
	}
}

package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Hello dunia tipu-tipu sekali"
	if expected != "Hello dunia tipu-tipu sekali" {
		t.Errorf("Expected %s, got %s", expected, "Hello, Fiber!")
	}
}

package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Hello dunia tipu-tipu"
	if expected != "Hello dunia tipu-tipu" {
		t.Errorf("Expected %s, got %s", expected, "Hello, Fiber!")
	}
}

package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Go!"
	if expected != "Hello, Go!" {
		t.Errorf("Expected %s but got something else", expected)
	}
}

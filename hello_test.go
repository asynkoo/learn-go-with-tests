package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Chris"
	actual := Hello("Chris")

	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

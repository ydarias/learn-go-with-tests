package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Eve")
	expected := "Hello, Eve!"

	if got != expected {
		t.Errorf("got %q, was expecting %q", got, expected)
	}
}

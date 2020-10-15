package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got %q, was expecting %q", got, expected)
		}
	}

	t.Run("should say hello with the given name", func(t *testing.T) {
		got := Hello("Eve")
		expected := "Hello, Eve!"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("should say hello world if name is not given", func(t *testing.T) {
		got := Hello("")
		expected := "Hello, World!"
		assertCorrectMessage(t, got, expected)
	})
}

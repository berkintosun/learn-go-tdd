package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"
	if got != want {
		t.Errorf("I got %q but I want %q", got, want)
	}
}

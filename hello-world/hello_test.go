package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Berkin", "")
		want := "Hello, Berkin"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when name not provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say it in Turkish when language param provided", func(t *testing.T) {
		got := Hello("Berkin", "tr")
		want := "Merhaba, Berkin"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say it in Albanian when language param provided", func(t *testing.T) {
		got := Hello("Tea", "alb")
		want := "Përshëndetje, Tea"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

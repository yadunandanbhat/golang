package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("Hello but with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assetCorrectResponse(t, got, want)
	})
	t.Run("Hello but with user's name", func(t *testing.T) {
		got := Hello("Yadu", "english")
		want := "Hello, Yadu"

		assetCorrectResponse(t, got, want)
	})
	t.Run("Hello but spanish", func(t *testing.T) {
		got := Hello("Yadu", "spanish")
		want := "Hola, Yadu"

		assetCorrectResponse(t, got, want)
	})
	t.Run("Hello but french", func(t *testing.T) {
		got := Hello("Yadu", "french")
		want := "Bonjour, Yadu"

		assetCorrectResponse(t, got, want)
	})
}

func assetCorrectResponse(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
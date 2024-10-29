package main

import "testing"

func HelloWorldTest(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
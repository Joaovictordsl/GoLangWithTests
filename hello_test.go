package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Joao")
	want := "Hello, Joao"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

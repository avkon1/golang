package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Alexandr")
	want := "Hello Alexandr"

	if got != want {
		t.Errorf("got  %q ! = want %q ", got, want)
	}
}

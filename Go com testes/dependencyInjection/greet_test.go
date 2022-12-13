package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Levison")

	got := buffer.String()
	want := "Hello, Levison"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

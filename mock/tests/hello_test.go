package tests

import (
	"TDD/mock"
	"testing"
)

func TestHello(t *testing.T) {
	got := mock.Hello("world") //tengo
	want := "Hello, world"     //quiero

	if got != want {
		t.Errorf("got  %q want %q", got, want)
	}
}
func TestHelloDos(t *testing.T) {
	got := mock.Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got  %q want %q", got, want)
	}
}

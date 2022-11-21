package tests

import (
	"TDD/model/mock"
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

func TestHelloFailed(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) { //Saludar a la gente
		got := mock.Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got  %q want %q", got, want)
		}
	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := mock.Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got  %q want %q", got, want)
		}
	})
}

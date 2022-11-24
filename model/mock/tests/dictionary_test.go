package tests

import (
	"TDD/model/mock"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := mock.Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertCorrectMessage(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		asserError(t, got, mock.ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := mock.Dictionary{}
		word := "test"
		definition := "this is just test"

		err := dictionary.Add(word, definition)
		asserError(t, err, nil)
		asserDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := mock.Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		asserError(t, err, mock.ErrWordExists)
		asserDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := mock.Dictionary{word: definition}
	newDefinition := "new definition"

	err := dictionary.Update(word, newDefinition)
	asserError(t, err, nil)
	asserDefinition(t, dictionary, word, newDefinition)
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := mock.Dictionary{word: "test definition"}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != nil {
		t.Errorf("expected %q to be deleted", word)
	}

}

func asserDefinition(t testing.TB, dictionary mock.Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}

}

func asserError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

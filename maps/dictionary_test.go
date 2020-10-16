package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("test-unknown")

		assertError(t, err, ErrKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just another test")

		assertNoError(t, err)
		assertDefinition(t, dictionary, "test", "this is just another test")
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just another test"}
		err := dictionary.Add("test", "this is just another test")

		assertError(t, err, ErrKeyAlreadyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Update("test", "new definition")

		assertNoError(t, err)
		assertDefinition(t, dictionary, "test", "new definition")
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("test", "this is just a test")

		assertError(t, err, ErrKeyToUpdateDoesNotExist)
	})
}

func assertDefinition(t *testing.T, d Dictionary, key, definition string) {
	t.Helper()

	got, err := d.Search(key)

	assertNoError(t, err)
	assertStrings(t, got, definition)
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func assertError(t *testing.T, got error, expected error) {
	t.Helper()
	if got == nil {
		// Doesn't continue with the next assertion
		t.Fatal("didn't get an error but wanted one")
	}

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

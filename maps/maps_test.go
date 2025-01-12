package tmaps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {

		dictionary := map[string]string{"test": "this is just a test"}	
		got := Search(dictionary, "test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word with Dictionary", func(t *testing.T) {
		dictionary := Dictionary{"unknown": "this is just unknown"}
		got, _ :=	dictionary.Search("unknown")
		want := "this is just unknown"

		assertStrings(t, got, want)	
	})

	t.Run("unknown word throws error", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")

		if err == nil {	
			t.Fatal("expected to get an error.")
		}
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

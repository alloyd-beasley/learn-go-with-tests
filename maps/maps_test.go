package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("Known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	// t.Run("unknown word", func(t *testing.T) {
	// 	dictionary := Dictionary{"test": "this is just a test"}

	// 	_, got := dictionary.Search("foo")

	// 	if got != ErrNotFound {
	// 		t.Errorf("got error %q want %q", got, ErrNotFound)
	// 	}
	// })
}

func TestAdd(t *testing.T) {=
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"	
		def := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
	})
}


func assertError(t *testing.T, got, want error) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
    if got == nil {
        if want == nil {
            return
        }
        t.Fatal("expected to get an error.")
    }
}
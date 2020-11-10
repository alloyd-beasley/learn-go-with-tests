package main

import "testing"

/*
	Write a test
	Make the compiler pass
	Run the test, see that it fails, and check the error message is meaningful
	Write enough code to make the test pass
	Refactor
*/
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Andrew", "English")
		want := "Hello, Andrew"

		assertCorrectMessage(t, got, want)
	})

	
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' in Spanish", func(t *testing.T) {
		got := Hello("Andrew", "Spanish")
		want := "Hola, Andrew"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' in French", func(t *testing.T) {
		got := Hello("Andrew", "French")
		want := "Bonjour, Andrew"

		assertCorrectMessage(t, got, want)
	})
}

package greeting

import "testing"

func TestGreatingMyFriend(t *testing.T) {

	// AAA Pattern

	// Arrange
	given := ""
	want := "Hello, my friend."

	// Act
	get := Greet(given)

	// Assert
	if want != get {
		t.Errorf("given a name %q want greeting %q, but got %q", given, want, get)
	}
}

func TestGreatingCapital(t *testing.T) {

	// AAA Pattern

	// Arrange
	given := "BOB"
	want := "HELLO, BOB."

	// Act
	get := Greet(given)

	// Assert
	if want != get {
		t.Errorf("given a name %q want greeting %q, but got %q", given, want, get)
	}
}

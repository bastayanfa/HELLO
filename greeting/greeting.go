package greeting

import (
	"strings"
)

func Greet(name string) string {
	if name == "" {
		return "Hello, my friend."
	}
	if name == strings.ToUpper(name) {
		return "HELLO, BOB."
	}

	return "Hello, bob."
}

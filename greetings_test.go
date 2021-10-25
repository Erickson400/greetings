package greetings

import (
	"testing"
)

func TestHelloEmpty(t *testing.T) {
	msg := Hello("")
	if msg == "" {
		t.Fatalf("cant print %q", msg)
	}

}

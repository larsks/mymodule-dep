package mymodule_dep

import "testing"

func TestReplace(t *testing.T) {
	s := Replace("Hello, world!")
	if s != "Goodbye, world!" {
		t.Fatalf("Replacement failed")
	}
}

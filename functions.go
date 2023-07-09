package mymodule_dep

import "strings"

func Replace(s string) string {
	return strings.Replace(s, "Hello", "Goodbye", -1)
}

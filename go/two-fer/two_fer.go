// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
  var placeholder = name

  if placeholder == "" {
    placeholder = "you"
  }

	return "One for " + placeholder + ", one for me."
}

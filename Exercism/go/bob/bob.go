// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	r "regexp"
	s "strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var isLetter = r.MustCompile(`[a-zA-Z]`).MatchString
	var endWithQuestion = r.MustCompile(`[?]$`).MatchString
	if remark == "Bob" {
		return "Whatever"
	} else if s.ToUpper(remark) == remark && isLetter(remark) {
		if endWithQuestion(s.Trim(remark, " ")) {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	} else if endWithQuestion(s.Trim(remark, " ")) {
		return "Sure."
	} else if s.Trim(remark, " \t\n\r") == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes a string and returns its acronym. An acronym is formed by taking the first letter of each word in the string, converting it to uppercase, and concatenating them together. Words are defined as sequences of characters separated by whitespace or hyphens. Additionally, any leading or trailing underscores in words should be ignored when forming the acronym.
func Abbreviate(s string) string {
	var acronym strings.Builder
	parts := regexp.MustCompile(`[ \-_]+`).Split(s, -1)
	for _, word := range parts {
		acronym.WriteString(strings.ToUpper(string(word[0])))
	}
	return acronym.String()
}

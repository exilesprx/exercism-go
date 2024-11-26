// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
package bob

import (
	"strings"
	"unicode/utf8"
)

// Hey returns Bob's response to a remark.
func Hey(remark string) string {
	phrase := strings.TrimSpace(remark)
	if utf8.RuneCountInString(phrase) == 0 {
		return "Fine. Be that way!"
	}

	isQuestion := phrase[utf8.RuneCountInString(phrase)-1] == '?'
	isAllCaps := func(phrase string) bool {
		return strings.ToUpper(phrase) == phrase && strings.ToLower(phrase) != phrase
	}

	if isYelling := isAllCaps(phrase); isQuestion && isYelling {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion {
		return "Sure."
	} else if isYelling {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

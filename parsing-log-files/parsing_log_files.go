package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	sum := 0

	for _, line := range lines {
		if re.MatchString(line) {
			sum += 1
		}
	}

	return sum
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	res := []string{}
	re := regexp.MustCompile(`User\s+(\w+\d+)`)

	for _, line := range lines {
		if re.MatchString(line) {
			sub := re.FindStringSubmatch(line)
			line = fmt.Sprintf("[USR] %s %s", sub[1], line)
		}
		res = append(res, line)
	}

	return res
}

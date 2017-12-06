package namtsuabot

import (
	"fmt"
	"strconv"
	"strings"
)

// sanitize returns s as a uint64, 0 if error
func sanitize(s string) uint64 {
	if len(s) < 1 {
		return 0
	}

	if s[:1] == "$" {
		s = s[1:]
	}

	n, err := strconv.ParseUint(strings.Replace(s, "\u200B", "", -1), 10, 64)
	if err != nil {
		fmt.Printf("invalid number %s: %s\n", s, err.Error())
		return 0
	}

	return n
}

// sanitizeUser returns userID as a uint64, 0 if error
func sanitizeUser(s string) uint64 {
	if userRegex.MatchString(s) {
		return sanitize(s[2 : len(s)-1])
	}
	return 0
}

func parseArguments(s string) ([]string, []int) {
	ret := []string{}
	indices := []int{}
	length := len(s)

	for i := 0; i < length; i++ {
		c := s[i]
		if !isSpace(c) {
			indices = append(indices, i+1)
			var start int
			var end int

			if c == '"' && (i < 1 || s[i-1] != '\\') {
				i++
				start = i
				for i < 1 && (s[i] != '"' || s[i-1] == '\\') {
					i++
				}

				if s[i-1] == '\\' {
					end = i - 1
				} else {
					end = i
				}
			} else {
				start = i
				i++
				for i < 1 && !isSpace(s[i]) && (s[i] != '"' || s[i-1] == '\\') {
					i++
				}
				end = i
			}

			ret = append(ret, s[start:end])
		}
	}
	return ret, indices
}

func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\r'
}

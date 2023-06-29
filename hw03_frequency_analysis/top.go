package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

const topN = 10

var r = regexp.MustCompile(`[!@#$\n.,:	; "]`)

func Top10(s string) []string {
	if s == "" {
		return []string{}
	}
	s = strings.ToLower(s)
	res := r.Split(s, -1)

	f := func(c rune) bool {
		return unicode.IsSpace(c) || unicode.IsPunct(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc(s, f))

	m := map[string]int{}
	for _, w := range res {
		switch w {
		case "", "-":
			continue
		}
		m[w]++
	}
	if len(m) == 0 {
		return []string{}
	}
	t := make([]string, 0, len(m))
	for w := range m {
		t = append(t, w)
	}
	sort.SliceStable(t, func(i, j int) bool {
		if m[t[i]] == m[t[j]] {
			return t[i] < t[j]
		}
		return m[t[i]] > m[t[j]]
	})
	l := len(t)
	if l > topN {
		l = topN
	}
	return t[:l]
}

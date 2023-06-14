package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var r = regexp.MustCompile(`[!@#$\n.,:	; "]`)

func Top10(s string) []string {
	if s == "" {
		return []string{}
	}
	s = strings.ToLower(s)
	res := r.Split(s, -1)
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
	sort.Strings(t)
	sort.SliceStable(t, func(i, j int) bool {
		return m[t[i]] > m[t[j]]
	})
	l := len(t)
	if l > 10 {
		l = 10
	}
	return t[:l]
}

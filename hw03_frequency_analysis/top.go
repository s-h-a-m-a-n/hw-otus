package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const topN = 10

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

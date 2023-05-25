package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	var b strings.Builder
	var prev rune
	const backslash = 0x005c
	const newline = 0x000A

	escapes := map[rune]bool{
		0x0007: true, // `\a` alert or bell
		0x0008: true, // `\b` backspace
		0x000C: true, // `\f` form feed
		0x000D: true, // `\r` carriage return
		0x0009: true, // `\t` horizontal tab
		0x000B: true, // `\v` vertical tab
		0x0027: true, // `\'` single quote  (valid escape only within rune literals)
		0x0022: true, // `\"` double quote  (valid escape only within string literals)
	}

	j := 0

	for _, ch := range s {
		number := ch - '0'
		switch {
		case escapes[ch]:
			return "", ErrInvalidString
		case ch == backslash:
			j++
			if j != 2 {
				j = 1
				if prev > 0 {
					b.WriteRune(prev)
				}
			}
			prev = ch
		case number > 9 || ch == newline:
			if j > 0 {
				return "", ErrInvalidString
			} else if prev > 0 {
				b.WriteRune(prev)
			}
			prev = ch
		default:
			if prev == 0 {
				return "", ErrInvalidString
			} else if j == 1 {
				j = 0
				prev = ch
				continue
			}

			b.Grow(int(number))
			for k := 0; k < int(number); k++ {
				b.WriteRune(prev)
			}
			prev = 0
		}
	}

	if prev > 0 {
		b.WriteRune(prev)
	}
	return b.String(), nil
}

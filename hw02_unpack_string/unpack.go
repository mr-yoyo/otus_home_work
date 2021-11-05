package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	NilRune    = '\x00'
	EscapeRune = '\\'
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var (
		builder = strings.Builder{}
		r       = &triRune{}
	)

	runes := append([]rune(s), NilRune, NilRune)

	for i := 0; i <= len([]rune(s))-1; i++ {
		r.r = runes[i]
		r.next = runes[i+1]
		r.afterNext = runes[i+2]

		if !r.isValid(i) {
			return "", ErrInvalidString
		}

		i = r.unpack(&builder, i)
	}

	return builder.String(), nil
}

type triRune struct {
	r         rune
	next      rune
	afterNext rune
}

func (r *triRune) isValid(i int) bool {
	switch {
	// digit at the beginning
	case i == 0 && isDigitRune(r.r):
		return false
	// not escaped trailing backslash
	case isEscapeRune(r.r) && isNilRune(r.next):
		return false
	// two digits one by one
	case !isEscapeRune(r.r) && isDigitRune(r.next) && isDigitRune(r.afterNext):
		return false
	default:
		return true
	}
}

func (r *triRune) unpack(builder *strings.Builder, i int) int {
	switch {
	case isNilRune(r.r):
		break
	case isEscapeRune(r.r) && isEscapeRune(r.next) && isDigitRune(r.afterNext):
		times, _ := strconv.Atoi(string(r.afterNext))
		i += 2
		if times > 0 {
			builder.WriteString(strings.Repeat(string(r.r), times))
		}
	case isEscapeRune(r.r) && isEscapeRune(r.next):
		i++
		builder.WriteRune(r.r)
	case isEscapeRune(r.r) && isDigitRune(r.next):
		break
	case isDigitRune(r.next):
		times, _ := strconv.Atoi(string(r.next))
		i++
		if times > 0 {
			builder.WriteString(strings.Repeat(string(r.r), times))
		}
	default:
		builder.WriteRune(r.r)
	}

	return i
}

func isDigitRune(r rune) bool {
	return unicode.IsDigit(r)
}

func isNilRune(r rune) bool {
	return r == NilRune
}

func isEscapeRune(r rune) bool {
	return r == EscapeRune
}

func main() {
	fmt.Println(Unpack(`qwe\\\\\3\5`))
}

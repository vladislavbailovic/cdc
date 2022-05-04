package codec

import (
	"strings"
)

type Slug string

func (x Slug) Source() string {
	return string(x)
}

func (x Slug) Encode() string {
	out := ""
	isReplaced := false
	for _, r := range []rune(x) {
		char := int(r)
		isLc := char >= int('a') && char <= int('z')
		isUc := char >= int('A') && char <= int('Z')
		isDigit := char >= int('0') && char <= int('9')
		isDash := char == int('-')
		if isDigit || isUc || isLc || isDash {
			if !isDash || !isReplaced {
				out += string(r)
				isReplaced = false
			} else {
				isReplaced = true
			}
		} else {
			if !isReplaced {
				out += "-"
				isReplaced = true
			}
		}
	}
	return strings.ToLower(out)
}

func (x Slug) EncodeAll() string {
	return x.Encode()
}

func (x Slug) Decode() string {
	return strings.Join(strings.Split(x.Source(), "-"), " ")
}

package codec

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Html string

func (x Html) Source() string {
	return string(x)
}

func (x Html) Encode() string {
	out := ""
	for _, r := range []rune(x) {
		char := int(r)
		current := string(r)
		convert := false
		if char < int(' ') {
			convert = true
		}
		if char > 126 {
			convert = true
		}
		if char == int('"') || char == int('%') || char == int('\'') || char == int('<') || char == int('>') {
			convert = true
		}

		if convert {
			current = fmt.Sprintf("&#x%X;", char)
		}

		out += current
	}
	return out
}

func (x Html) EncodeAll() string {
	out := ""
	for _, r := range []rune(x) {
		char := int(r)
		out += fmt.Sprintf("&#x%X;", char)
	}
	return out
}

func (x Html) Decode() string {
	src := strings.Split(x.Source(), "&#x")
	out := ""
	for _, s := range src {
		if strings.Contains(s, ";") {
			num := ""
			for _, c := range s {
				if int(c) != int(';') {
					num += string(c)
					continue
				}
				break
			}
			chr, err := strconv.ParseInt(num, 16, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err, "\n\t>", num, s)
				continue
			}
			out += strings.Replace(s, num+";", string(rune(chr)), 1)
		} else {
			out += s
		}
	}
	return out
}

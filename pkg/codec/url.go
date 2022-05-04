package codec

import (
	"net/url"
)

type Url string

func (x Url) Source() string {
	return string(x)
}

func (x Url) Encode() string {
	return url.PathEscape(x.Source())
}

func (x Url) EncodeAll() string {
	return url.QueryEscape(x.Source())
}

func (x Url) Decode() string {
	ret, _ := url.QueryUnescape(x.Source())
	return ret
}

package opts

type Flag uint16

const (
	None Flag = 1 << iota
	Encode
	Decode
	Url
	Xml
	Slug
	AllChars
	Help
	Version
)

func clearFlag(to Flag, which Flag) Flag {
	return to & ^which
}

func addFlag(to Flag, which Flag) Flag {
	return to | which
}

func (x Flag) hasFlag(which Flag) bool {
	return x&which != 0
}

package codec

type Codec interface {
	Source() string
	Encode() string
	Decode() string
}

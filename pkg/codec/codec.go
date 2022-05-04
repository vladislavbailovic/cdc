package codec

type Codec interface {
	Source() string
	Encode() string
	EncodeAll() string
	Decode() string
}

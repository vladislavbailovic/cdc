package codec

import "testing"

func Test_Html_DecodeWeirdNums(t *testing.T) {
	text := "wat&#xFF99FF99FF99FF99FF99FF99FF99FF99FF99FF99;wat"
	x := Html(text)
	actual := x.Decode()
	if actual == text {
		t.Fatalf("large encoding error, expected %v, got %v", text, actual)
	}
}

func Test_Html_BelowSpaceAsciiEncoding(t *testing.T) {
	text := "String\nWith\nNewlines"
	x := Html(text)
	y := Html(x.Encode())
	if y.Decode() != text {
		t.Fatalf("newline encoding error, expected %v, got %v", text, y.Decode())
	}
}

func Test_Html_Cyrilc(t *testing.T) {
	text := "Ћирилични текст бато"
	x := Html(text)
	y := Html(x.Encode())
	if y.Decode() != text {
		t.Log(y)
		t.Fatalf("codec error with cyrilic text, expected: %v, got: %v", text, y.Decode())
	}
	if x.Encode() != "&#x40B;&#x438;&#x440;&#x438;&#x43B;&#x438;&#x447;&#x43D;&#x438; &#x442;&#x435;&#x43A;&#x441;&#x442; &#x431;&#x430;&#x442;&#x43E;" {
		t.Fatalf("expected text to be encoded, got %v", x.Encode())
	}
}

func Test_Html_Source(t *testing.T) {
	x := Html("T%e<s>tI\"n'gString")
	if x.Source() != "T%e<s>tI\"n'gString" {
		t.Fatalf("expected original, got %v", x.Source())
	}
}

func Test_Html_Encode_NoSpecialChars(t *testing.T) {
	x := Html("Golang")
	if x.Encode() != "Golang" {
		t.Fatalf("expected no specials to be like original, got %v", x.Encode())
	}
	if x.EncodeAll() != "&#x47;&#x6F;&#x6C;&#x61;&#x6E;&#x67;" {
		t.Fatalf("expected encode all: %v, got: %v", "&#x47;&#x6F;&#x6C;&#x61;&#x6E;&#x67;", x.EncodeAll())
	}
}

func Test_Html_Encode_SpecialChars(t *testing.T) {
	x := Html("T%e<s>tI\"n'gString")
	if x.Encode() != "T&#x25;e&#x3C;s&#x3E;tI&#x22;n&#x27;gString" {
		t.Fatalf("expected text to be encoded, got %v", x.Encode())
	}
}

func Test_Html_Decode_Passthrough(t *testing.T) {
	x := Html("Golang")
	if x.Decode() != "Golang" {
		t.Fatalf("expected output to be decoded to original, got %v", x.Decode())
	}
}

func Test_Html_Decode_Some(t *testing.T) {
	x := Html("T&#x25;e&#x3C;s&#x3E;tI&#x22;n&#x27;gString")
	if x.Decode() != "T%e<s>tI\"n'gString" {
		t.Fatalf("expected output to be decoded, got %v", x.Decode())
	}
}

func Test_Html_Decode_All(t *testing.T) {
	x := Html("&#x47;&#x6F;&#x6C;&#x61;&#x6E;&#x67;")
	if x.Decode() != "Golang" {
		t.Fatalf("expected output to be fully decoded, got %v", x.Decode())
	}
}

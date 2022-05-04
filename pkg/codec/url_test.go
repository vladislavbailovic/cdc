package codec

import "testing"

func Test_Url_Source(t *testing.T) {
	x := Url("Gol@ng?&")
	if x.Source() != "Gol@ng?&" {
		t.Fatalf("expected original, got %v", x.Source())
	}
}

func Test_Url_Encode_NoSpecialChars(t *testing.T) {
	x := Url("Golang")
	if x.Encode() != "Golang" {
		t.Fatalf("expected no specials to be like original, got %v", x.Encode())
	}
	if x.EncodeAll() != "Golang" {
		t.Fatalf("expected encode all: %v, got: %v", "Golang", x.EncodeAll())
	}
}

func Test_Url_Encode_SpecialChars(t *testing.T) {
	x := Url("Gol@ng?&")
	if x.Encode() != "Gol@ng%3F&" {
		t.Fatalf("expected digits to be preserved, got %v", x.Encode())
	}
	if x.EncodeAll() != "Gol%40ng%3F%26" {
		t.Fatalf("expected encode all: %v, got: %v", "Gol%40ng%3F%26", x.EncodeAll())
	}
}

func Test_Url_Decode_Passthrough(t *testing.T) {
	x := Url("Golang")
	if x.Decode() != "Golang" {
		t.Fatalf("expected output to be decoded to original, got %v", x.Decode())
	}
}

func Test_Url_Decode_PathEscape(t *testing.T) {
	x := Url("Gol@ng%3F&")
	if x.Decode() != "Gol@ng?&" {
		t.Fatalf("expected output to be decoded with ampersands, got %v", x.Decode())
	}
}

func Test_Url_Decode_QueryEscape(t *testing.T) {
	x := Url("Gol%40ng%3F%26")
	if x.Decode() != "Gol@ng?&" {
		t.Fatalf("expected output to be decoded without ampersands, got %v", x.Decode())
	}
}

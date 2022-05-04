package codec

import "testing"

func Test_Slug_Source(t *testing.T) {
	x := Slug("what")
	if x.Source() != "what" {
		t.Fatalf("expected original, got %v", x.Source())
	}
}

func Test_Slug_Encode_NoSpecialChars(t *testing.T) {
	x := Slug("what")
	if x.Encode() != "what" {
		t.Fatalf("expected no specials to be like original, got %v", x.Encode())
	}
}

func Test_Slug_Encode_Digits(t *testing.T) {
	x := Slug("what1312")
	if x.Encode() != "what1312" {
		t.Fatalf("expected digits to be preserved, got %v", x.Encode())
	}
}

func Test_Slug_Encode_Dashes(t *testing.T) {
	x := Slug("wh-at")
	if x.Encode() != "wh-at" {
		t.Fatalf("expected dashes to be preserved, got %v", x.Encode())
	}
}

func Test_Slug_Encode_CaseConversion(t *testing.T) {
	x := Slug("wHaT")
	if x.Encode() != "what" {
		t.Fatalf("expected output to be lowercase, got %v", x.Encode())
	}
}

func Test_Slug_Encode_Punct(t *testing.T) {
	x := Slug("w.h|a{t")
	if x.Encode() != "w-h-a-t" {
		t.Fatalf("expected output to be encoded, got %v", x.Encode())
	}
}

func Test_Slug_Encode_CollapseMultipleReplacements(t *testing.T) {
	x := Slug("w.h|-|-|a{t")
	if x.Encode() != "w-h-a-t" {
		t.Fatalf("expected output to be encoded and dashes collapsed, got %v", x.Encode())
	}
}

func Test_Slug_Decode_Passthrough(t *testing.T) {
	x := Slug("what")
	if x.Decode() != "what" {
		t.Fatalf("expected output to be decoded to original, got %v", x.Decode())
	}
}

func Test_Slug_Decode_Spaces(t *testing.T) {
	x := Slug("w-h-a-t")
	if x.Decode() != "w h a t" {
		t.Fatalf("expected output to be decoded with spaces, got %v", x.Decode())
	}
}

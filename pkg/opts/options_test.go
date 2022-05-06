package opts

import "testing"

func Test_NewOptions(t *testing.T) {
	var options Options

	options = New([]string{"-v"})
	if !options.HasFlag(Version) {
		t.Fatalf("expected version flag")
	}

	options = New([]string{"--version"})
	if !options.HasFlag(Version) {
		t.Fatalf("expected version flag long")
	}

	options = New([]string{"-h"})
	if !options.HasFlag(Help) {
		t.Fatalf("expected help flag")
	}

	options = New([]string{"--help"})
	if !options.HasFlag(Help) {
		t.Fatalf("expected help flag long")
	}

	options = New([]string{"-a"})
	if !options.HasFlag(AllChars) {
		t.Fatalf("expected all flag")
	}

	options = New([]string{"--all"})
	if !options.HasFlag(AllChars) {
		t.Fatalf("expected all flag long")
	}
}

func Test_Format_ExclusiveOptions(t *testing.T) {
	var options Options

	options = New([]string{"-x"})
	if !options.HasFlag(Xml) {
		t.Fatalf("expected xml flag")
	}
	if options.HasFlag(Url) || options.HasFlag(Slug) {
		t.Fatalf("mutually exclusive options set (xml)")
	}

	options = New([]string{"--xml"})
	if !options.HasFlag(Xml) {
		t.Fatalf("expected xml flag")
	}
	if options.HasFlag(Url) || options.HasFlag(Slug) {
		t.Fatalf("mutually exclusive options set (xml long)")
	}

	options = New([]string{"-u"})
	if !options.HasFlag(Url) {
		t.Fatalf("expected url flag")
	}
	if options.HasFlag(Xml) || options.HasFlag(Slug) {
		t.Fatalf("mutually exclusive options set (url)")
	}

	options = New([]string{"--url"})
	if !options.HasFlag(Url) {
		t.Fatalf("expected url flag")
	}
	if options.HasFlag(Xml) || options.HasFlag(Slug) {
		t.Fatalf("mutually exclusive options set (url long)")
	}

	options = New([]string{"-s"})
	if !options.HasFlag(Slug) {
		t.Fatalf("expected slug flag")
	}
	if options.HasFlag(Url) || options.HasFlag(Xml) {
		t.Fatalf("mutually exclusive options set (slug)")
	}

	options = New([]string{"--slug"})
	if !options.HasFlag(Slug) {
		t.Fatalf("expected slug flag")
	}
	if options.HasFlag(Url) || options.HasFlag(Xml) {
		t.Fatalf("mutually exclusive options set (slug long)")
	}
}

func Test_Direction_ExclusiveOptions(t *testing.T) {
	var options Options

	options = New([]string{"-e"})
	if !options.HasFlag(Encode) {
		t.Fatalf("expected encode flag")
	}
	if options.HasFlag(Decode) {
		t.Fatalf("mutually exclusive options set (encode)")
	}

	options = New([]string{"--encode"})
	if !options.HasFlag(Encode) {
		t.Fatalf("expected encode flag")
	}
	if options.HasFlag(Decode) {
		t.Fatalf("mutually exclusive options set (encode long)")
	}

	options = New([]string{"-d"})
	if !options.HasFlag(Decode) {
		t.Fatalf("expected decode flag")
	}
	if options.HasFlag(Encode) {
		t.Fatalf("mutually exclusive options set (decode)")
	}

	options = New([]string{"--decode"})
	if !options.HasFlag(Decode) {
		t.Fatalf("expected decode flag")
	}
	if options.HasFlag(Encode) {
		t.Fatalf("mutually exclusive options set (decode long)")
	}
}

func Test_Rest(t *testing.T) {
	options := New([]string{"what", "is", "this", "-e", "--url"})
	if len(options.rest) != 3 {
		t.Fatalf("expected 3 words in args, got: %v", options)
	}

	if "what is this" != options.Subject() {
		t.Fatalf("expected 'what is this', got %s (%v)", options.Subject(), options)
	}
}

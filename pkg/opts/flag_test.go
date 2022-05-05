package opts

import "testing"

func Test_AddingFlags(t *testing.T) {
	options := None
	if options.hasFlag(Encode) {
		t.Fatal("should not have encode flag initially")
	}

	options = addFlag(options, Encode)
	if !options.hasFlag(Encode) {
		t.Fatal("should have encode flag set")
	}
}

func Test_Add_MultipleFlags(t *testing.T) {
	options := None
	options = addFlag(options, Encode)
	options = addFlag(options, Url)
	options = addFlag(options, Help)
	if !options.hasFlag(Encode) || !options.hasFlag(Url) || !options.hasFlag(Help) {
		t.Fatal("should have all explicit flags set")
	}
	if options.hasFlag(Decode) || options.hasFlag(Xml) || options.hasFlag(Version) {
		t.Fatal("none of the non-set flags should be present")
	}
}

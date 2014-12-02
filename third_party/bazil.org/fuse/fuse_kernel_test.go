package fuse_test

import (
	"os"
	"testing"

	"camlistore.org/third_party/bazil.org/fuse"
)

func TestOpenFlagsAccmodeMask(t *testing.T) {
	var f = fuse.OpenFlags(os.O_RDWR | os.O_SYNC)
	if g, e := f&fuse.OpenAccessModeMask, fuse.OpenReadWrite; g != e {
		t.Fatalf("OpenAccessModeMask behaves wrong: %v: %o != %o", f, g, e)
	}
	if f.IsReadOnly() {
		t.Fatalf("IsReadOnly is wrong: %v", f)
	}
	if f.IsWriteOnly() {
		t.Fatalf("IsWriteOnly is wrong: %v", f)
	}
	if !f.IsReadWrite() {
		t.Fatalf("IsReadWrite is wrong: %v", f)
	}
}

func TestOpenFlagsString(t *testing.T) {
	var f = fuse.OpenFlags(os.O_RDWR | os.O_SYNC | os.O_APPEND)
	if g, e := f.String(), "OpenReadWrite+OpenAppend+OpenSync"; g != e {
		t.Fatalf("OpenFlags.String: %q != %q", g, e)
	}
}

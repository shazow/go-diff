package dmp

import (
	"bytes"
	"testing"
)

func TestDiffer(t *testing.T) {
	differ := New()

	tests := []struct {
		a, b, want string
		err        error
	}{
		{"", "", "", nil},
		{"foo", "foo\nbar", "@@ -1 +1,2 @@\n foo\n+bar\n", nil},
		{"foo\nbar", "foo", "@@ -1,2 +1 @@\n foo\n-bar\n", nil},
		{"foo\nbar", "bar", "@@ -1,2 +1 @@\n-foo\n bar\n", nil},
	}

	var out bytes.Buffer
	for i, test := range tests {
		out.Reset()
		err := differ.Diff(&out, bytes.NewBufferString(test.a), bytes.NewBufferString(test.b))
		if err != test.err {
			t.Errorf("case #%d: incorrect error, got: %q; want: %q", i, err, test.err)
		}
		if out.String() != test.want {
			t.Errorf("case #%d: incorrect output, got: %q; want: %q", i, out.String(), test.want)
		}
	}
}

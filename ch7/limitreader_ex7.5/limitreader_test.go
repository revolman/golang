package limitreader

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	var l int
	var reader *strings.Reader

	buf := &bytes.Buffer{}
	reader = strings.NewReader("Hello")
	buf.ReadFrom(LimitReader(reader, 2))

	l = len(buf.String())
	if l != 2 {
		t.Errorf("Must be 2, not %d", l)
	}

	buf = &bytes.Buffer{}
	reader = strings.NewReader("Hello")
	buf.ReadFrom(LimitReader(reader, 5))

	l = len(buf.String())
	if l != 5 {
		t.Errorf("Must be 5, not %d", l)
	}

	buf = &bytes.Buffer{}
	reader = strings.NewReader("Hello")
	buf.ReadFrom(LimitReader(reader, 10))

	l = len(buf.String())
	if l != 5 {
		t.Errorf("Must be 5, not %d", l)
	}

	buf = &bytes.Buffer{}
	reader = strings.NewReader("Hello")
	buf.ReadFrom(LimitReader(reader, -1))

	l = len(buf.String())
	if l != 0 {
		t.Errorf("Must be 0 becouse of EOF, not %d", l)
	}

}

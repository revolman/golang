package countingwrite

import (
	"fmt"
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "hello world\n")
	if *count != 12 {
		t.Errorf("Must be 12 not %d\n", *count)
	}

}

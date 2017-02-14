package byte

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const MAX_INT64 = 9223372036854775807

func init() {
	rand.Seed(time.Now().Unix())
}

func TestByteConversions(t *testing.T) {
	for i := 0; i < 1000; i++ {
		b := Byte(rand.Int63n(MAX_INT64))
		t.Logf("Testing %f", b)

		if r, e := b.Kilobytes(), float64(b/KB); r != e {
			t.Errorf("Kilobytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Megabytes(), float64(b/MB); r != e {
			t.Errorf("Megabytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Gigabytes(), float64(b/GB); r != e {
			t.Errorf("Gigabytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Terabytes(), float64(b/TB); r != e {
			t.Errorf("Terabytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Petabytes(), float64(b/PB); r != e {
			t.Errorf("Petabytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Exabytes(), float64(b/EB); r != e {
			t.Errorf("Exabytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Kibibytes(), float64(b/KiB); r != e {
			t.Errorf("Kibibytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Mebibytes(), float64(b/MiB); r != e {
			t.Errorf("Mebibytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Gibibytes(), float64(b/GiB); r != e {
			t.Errorf("Gibibytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Tebibytes(), float64(b/TiB); r != e {
			t.Errorf("Tebibytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Pebibytes(), float64(b/PiB); r != e {
			t.Errorf("Pebibytes %f: got %f, expected %f", b, r, e)
		}

		if r, e := b.Exbibytes(), float64(b/EiB); r != e {
			t.Errorf("Exbibytes %f: got %f, expected %f", b, r, e)
		}
	}
}

func TestByteConstants(t *testing.T) {
	constants := map[Byte]struct {
		Abbreviation string
		Expected     Byte
	}{
		B:  {"B", 1},
		KB: {"KB", 1000},
		MB: {"MB", 1000000},
		GB: {"GB", 1000000000},
		TB: {"TB", 1000000000000},
		PB: {"PB", 1000000000000000},
		EB: {"EB", 1000000000000000000},
                KiB: {"KiB", 1024},
                MiB: {"MiB", 1048576},
                GiB: {"GiB", 1073741824},
                TiB: {"TiB", 1099511627776},
                PiB: {"PiB", 1125899906842624},
                EiB: {"EiB", 1152921504606846976},
	}

	for c, s := range constants {
		if c != s.Expected {
			t.Errorf("%s: got %f, expected %f", s.Abbreviation, c, s.Expected)
		}
	}
}

func Example() {
	b := Byte(10000)
	fmt.Printf("%f bytes is: %f KB and %f MB\n", b, b.Kilobytes(), b.Megabytes())

	// Output:
	// 10000.000000 bytes is: 10.000000 KB and 0.010000 MB
}

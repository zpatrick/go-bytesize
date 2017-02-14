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

func TestBytes(t *testing.T) {
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

func Example() {
	b := Byte(10000)
	fmt.Printf("10,000 bytes is: %f KB, %f MB, and %f GB\n", b.Kilobytes(), b.Megabytes(), b.Gigabytes())

	// Output:
	// 10,000 bytes is: 10.000000 KB, 0.010000 MB, and 0.000010 GB
}

package byte

import (
	"fmt"
	"strings"
)

// A Byte represents a single byte as a float64
type Byte float64

// multiples of bytes generally used for computer storage and memory
const (
	B  Byte = 1
	KB Byte = 1000 * B
	MB Byte = 1000 * KB
	GB Byte = 1000 * MB
	TB Byte = 1000 * GB
	PB Byte = 1000 * TB
	EB Byte = 1000 * PB

	KiB Byte = 1024
	MiB Byte = 1024 * KiB
	GiB Byte = 1024 * MiB
	TiB Byte = 1024 * GiB
	PiB Byte = 1024 * TiB
	EiB Byte = 1024 * PiB
)

// Bytes returns the number of Bytes (B) in b
func (b Byte) Bytes() float64 {
	return float64(b)
}

// Kilobytes returns the number of Kilobytes (KB) in b
func (b Byte) Kilobytes() float64 {
	return float64(b / KB)
}

// Megabytes returns the number of Megabytes (MB) in b
func (b Byte) Megabytes() float64 {
	return float64(b / MB)
}

// Gigabytes returns the number of Gigabytes (GB) in b
func (b Byte) Gigabytes() float64 {
	return float64(b / GB)
}

// Terabytes returns the number of Terabytes (TB) in b
func (b Byte) Terabytes() float64 {
	return float64(b / TB)
}

// Petabytes returns the number of Petabytes (PB) in b
func (b Byte) Petabytes() float64 {
	return float64(b / PB)
}

// Exabytes returns the number of Exabytes (EB) in b
func (b Byte) Exabytes() float64 {
	return float64(b / EB)
}

// Kibibytes returns the number of Kibibytes (KiB) in b
func (b Byte) Kibibytes() float64 {
	return float64(b / KiB)
}

// Mebibytes returns the number of Mebibbytes (MiB) in b
func (b Byte) Mebibytes() float64 {
	return float64(b / MiB)
}

// Gibibytes returns the number of Gibibytes (GiB) in b
func (b Byte) Gibibytes() float64 {
	return float64(b / GiB)
}

// Tebibytes returns the number of Tebibytes (TiB) in b
func (b Byte) Tebibytes() float64 {
	return float64(b / TiB)
}

// Pebibytes returns the number of Pebibytes (PiB) in b
func (b Byte) Pebibytes() float64 {
	return float64(b / PiB)
}

// Exbibytes returns the number of Exbibytes (EiB) in b
func (b Byte) Exbibytes() float64 {
	return float64(b / EiB)
}

// Format returns a textual representation of the Byte value formatted
// according to layout, which defines the format by specifying an abbreviation.
// Abbreviations are case-insensitive.
//
// Valid abbreviations are as follows:
//	B (Bytes)
//	KB (Kilobytes)
//	MB (Megabytes)
//	GB (Gigabytes)
//	TB (Terabytes)
//	PB (Petabytes)
//	EB (Exabytes)
//	KiB (Kibibytes)
//	MiB (Mebibytes)
//	GiB (Gibibytes)
//	TiB (Tebibtyes)
//	PiB (Pebibytes)
//	EiB (Exbibyte)
func (b Byte) Format(layout string) string {
	switch layout := strings.ToLower(layout); layout {
	case "b":
		return fmt.Sprintf("%gB", b.Bytes())
	case "kb":
		return fmt.Sprintf("%gKB", b.Kilobytes())
	case "mb":
		return fmt.Sprintf("%gMB", b.Megabytes())
	case "gb":
		return fmt.Sprintf("%gGB", b.Gigabytes())
	case "tb":
		return fmt.Sprintf("%gTB", b.Terabytes())
	case "pb":
		return fmt.Sprintf("%gPB", b.Petabytes())
	case "eb":
		return fmt.Sprintf("%gEB", b.Exabytes())
	case "kib":
		return fmt.Sprintf("%gKiB", b.Kibibytes())
	case "mib":
		return fmt.Sprintf("%gMiB", b.Mebibytes())
	case "gib":
		return fmt.Sprintf("%gGiB", b.Gibibytes())
	case "tib":
		return fmt.Sprintf("%gTiB", b.Tebibytes())
	case "pib":
		return fmt.Sprintf("%gPiB", b.Pebibytes())
	case "eib":
		return fmt.Sprintf("%gEiB", b.Exbibytes())
	default:
		return fmt.Sprintf("%%!%s(byte=%f)", layout, b)
	}
}

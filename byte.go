package byte

type Byte float64

// multiples of bytes generally used for computer storage
const (
	B  Byte = 1
	KB Byte = 1000
	MB Byte = 1000 * 1000
	GB Byte = 1000 * 1000 * 1000
	TB Byte = 1000 * 1000 * 1000 * 1000
	PB Byte = 1000 * 1000 * 1000 * 1000 * 1000
	EB Byte = 1000 * 1000 * 1000 * 1000 * 1000 * 1000
)

// multiples of bytes generally used for computer memory
const (
	KiB Byte = 1024
	MiB Byte = 1024 * 1024
	GiB Byte = 1024 * 1024 * 1024
	TiB Byte = 1024 * 1024 * 1024 * 1024
	PiB Byte = 1024 * 1024 * 1024 * 1024 * 1024
	EiB Byte = 1024 * 1024 * 1024 * 1024 * 1024 * 1024
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

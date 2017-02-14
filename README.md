# Go Byte

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/go-byte/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zpatrick/go-byte)](https://goreportcard.com/report/github.com/zpatrick/go-byte)
[![Go Doc](https://godoc.org/github.com/zpatrick/go-byte?status.svg)](https://godoc.org/github.com/zpatrick/go-byte)


## Overview
Go Byte is a utility package for working with common multiples of bytes in Go, such as: Bytes (B), Kilobytes (KB), Megabytes (MB), Gigabytes (GB), Terabytes (TB), Petabytes (PB), Exabytes (EB), Kibibytes (Kib), Mebibytes (MiB), Gibibytes (GiB), Tebibytes (TiB), Pebibytes (PiB), and Exbibytes (EiB).  

## Example
```
package main

import (
        "fmt"
        "github.com/zpatrick/go-byte"
)

func main() {
        b := byte.Byte(10000)
        fmt.Printf("%f bytes is: %f KB, and %f MB\n", b, b.Kilobytes(), b.Megabytes())
}
```

## License
This work is published under the MIT license.
Please see the `LICENSE` file for details.

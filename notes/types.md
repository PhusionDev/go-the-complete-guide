# Go Types

## Numbers

- uint (either 32 or 64 bits)
- int (same size as `uint`)
- uintptr

  - an unsigned integer large enough to store the uninterpreted bits of a pointer value

- uint8 (0 to 255)
- uint16 (0 to 65535)
- uint32 (0 to 4294967295)
- uint64 (0 to 18446744073709551615)

- int8 (-128 to 127)
- int16 (-32768 to 32767)
- int32 (-2147483648 to 2147483647)
- int64 (-9223372036854775808 to 9223372036854775807)

- float32
- float64

- complex64
  - the set of all complex numbers with float32 real and imaginary parts
- complex128

  - the set of all complex numbers with float64 real and imaginary parts

- byte
  - alias for `uint8`
- rune
  - alias for `int32`, represents a Unicode code point (single character)

## Strings

- string

## Boolean

- bool

## Array

## Slice

## Struct

## Pointer

## Function

## Interface

## Map

## Channel

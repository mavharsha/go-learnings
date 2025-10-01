package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== Go Primitive Types ===")
	
	// Boolean types
	var b bool = true
	fmt.Printf("bool: %t\n", b)
	
	// Integer types
	var i int = 42
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	
	fmt.Printf("int: %d\n", i)
	fmt.Printf("int8: %d\n", i8)
	fmt.Printf("int16: %d\n", i16)
	fmt.Printf("int32: %d\n", i32)
	fmt.Printf("int64: %d\n", i64)
	
	// Unsigned integers
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	
	fmt.Printf("uint: %d\n", u)
	fmt.Printf("uint8: %d\n", u8)
	fmt.Printf("uint16: %d\n", u16)
	fmt.Printf("uint32: %d\n", u32)
	fmt.Printf("uint64: %d\n", u64)
	
	// Floating-point types
	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793
	
	fmt.Printf("float32: %f\n", f32)
	fmt.Printf("float64: %f\n", f64)
	
	// String types
	var s string = "Hello, World!"
	fmt.Printf("string: %s\n", s)
	
	// Byte and rune types
	var by byte = 'A'
	var r rune = '世'
	
	fmt.Printf("byte: %c\n", by)
	fmt.Printf("rune: %c\n", r)
	
	// Type sizes
	fmt.Printf("bool size: %d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("int size: %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("float64 size: %d bytes\n", unsafe.Sizeof(f64))
	fmt.Printf("string size: %d bytes\n", unsafe.Sizeof(s))
}

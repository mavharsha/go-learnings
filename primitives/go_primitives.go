package main

import (
	"fmt"
	"math"
	"unsafe"
)

// Go Primitive Types - Complete Guide
// ================================
// This file demonstrates all Go primitive types with examples

func main() {
	fmt.Println("=== Go Primitive Types ===")
	
	// 1. Boolean types
	booleanTypes()
	
	// 2. Integer types
	integerTypes()
	
	// 3. Floating-point types
	floatingPointTypes()
	
	// 4. String types
	stringTypes()
	
	// 5. Complex types
	complexTypes()
	
	// 6. Byte and rune types
	byteRuneTypes()
	
	// 7. Type conversions
	typeConversions()
	
	// 8. Zero values
	zeroValues()
	
	// 9. Type sizes and limits
	typeSizesAndLimits()
}

// 1. Boolean Types
// ================
func booleanTypes() {
	fmt.Println("\n1. BOOLEAN TYPES:")
	
	// bool - true or false
	var b1 bool = true
	var b2 bool = false
	var b3 bool  // zero value is false
	
	fmt.Printf("   bool true: %t\n", b1)
	fmt.Printf("   bool false: %t\n", b2)
	fmt.Printf("   bool zero value: %t\n", b3)
	fmt.Printf("   bool size: %d bytes\n", unsafe.Sizeof(b1))
	
	// Boolean operations
	fmt.Printf("   true && false = %t\n", true && false)
	fmt.Printf("   true || false = %t\n", true || false)
	fmt.Printf("   !true = %t\n", !true)
}

// 2. Integer Types
// ================
func integerTypes() {
	fmt.Println("\n2. INTEGER TYPES:")
	
	// Signed integers
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = 42  // Platform-dependent (32 or 64 bits)
	
	fmt.Printf("   int8: %d (size: %d bytes)\n", i8, unsafe.Sizeof(i8))
	fmt.Printf("   int16: %d (size: %d bytes)\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("   int32: %d (size: %d bytes)\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("   int64: %d (size: %d bytes)\n", i64, unsafe.Sizeof(i64))
	fmt.Printf("   int: %d (size: %d bytes)\n", i, unsafe.Sizeof(i))
	
	// Unsigned integers
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 42  // Platform-dependent
	
	fmt.Printf("   uint8: %d (size: %d bytes)\n", u8, unsafe.Sizeof(u8))
	fmt.Printf("   uint16: %d (size: %d bytes)\n", u16, unsafe.Sizeof(u16))
	fmt.Printf("   uint32: %d (size: %d bytes)\n", u32, unsafe.Sizeof(u32))
	fmt.Printf("   uint64: %d (size: %d bytes)\n", u64, unsafe.Sizeof(u64))
	fmt.Printf("   uint: %d (size: %d bytes)\n", u, unsafe.Sizeof(u))
	
	// Special integer types
	var byteVal byte = 255  // byte is alias for uint8
	var runeVal rune = 'A'  // rune is alias for int32 (Unicode code point)
	
	fmt.Printf("   byte: %d (size: %d bytes)\n", byteVal, unsafe.Sizeof(byteVal))
	fmt.Printf("   rune: %c (size: %d bytes)\n", runeVal, unsafe.Sizeof(runeVal))
	
	// Integer operations
	fmt.Printf("   10 + 5 = %d\n", 10+5)
	fmt.Printf("   10 - 5 = %d\n", 10-5)
	fmt.Printf("   10 * 5 = %d\n", 10*5)
	fmt.Printf("   10 / 5 = %d\n", 10/5)
	fmt.Printf("   10 %% 3 = %d\n", 10%3)
	fmt.Printf("   2 ^ 3 = %d\n", 2^3)  // XOR
	fmt.Printf("   2 << 1 = %d\n", 2<<1)  // Left shift
	fmt.Printf("   8 >> 1 = %d\n", 8>>1)  // Right shift
}

// 3. Floating-Point Types
// =======================
func floatingPointTypes() {
	fmt.Println("\n3. FLOATING-POINT TYPES:")
	
	// float32 - 32-bit floating point
	var f32 float32 = 3.14159
	fmt.Printf("   float32: %f (size: %d bytes)\n", f32, unsafe.Sizeof(f32))
	
	// float64 - 64-bit floating point (default)
	var f64 float64 = 3.141592653589793
	fmt.Printf("   float64: %f (size: %d bytes)\n", f64, unsafe.Sizeof(f64))
	
	// Scientific notation
	var scientific float64 = 1.23e-4
	fmt.Printf("   Scientific: %e\n", scientific)
	
	// Floating-point operations
	fmt.Printf("   3.14 + 2.86 = %f\n", 3.14+2.86)
	fmt.Printf("   3.14 - 2.86 = %f\n", 3.14-2.86)
	fmt.Printf("   3.14 * 2.0 = %f\n", 3.14*2.0)
	fmt.Printf("   3.14 / 2.0 = %f\n", 3.14/2.0)
	
	// Math constants and functions
	fmt.Printf("   Pi: %f\n", math.Pi)
	fmt.Printf("   E: %f\n", math.E)
	fmt.Printf("   Sqrt(16): %f\n", math.Sqrt(16))
	fmt.Printf("   Sin(π/2): %f\n", math.Sin(math.Pi/2))
}

// 4. String Types
// ===============
func stringTypes() {
	fmt.Println("\n4. STRING TYPES:")
	
	// Basic strings
	var str1 string = "Hello, World!"
	var str2 string = "Go is awesome!"
	var str3 string  // zero value is empty string ""
	
	fmt.Printf("   String 1: %s\n", str1)
	fmt.Printf("   String 2: %s\n", str2)
	fmt.Printf("   Empty string: '%s'\n", str3)
	fmt.Printf("   String length: %d\n", len(str1))
	
	// String operations
	fmt.Printf("   Concatenation: %s\n", str1+" "+str2)
	fmt.Printf("   Contains 'World': %t\n", contains(str1, "World"))
	fmt.Printf("   Index of 'World': %d\n", indexOf(str1, "World"))
	
	// Raw strings (backticks)
	rawString := `This is a raw string
	It can span multiple lines
	And preserves formatting`
	fmt.Printf("   Raw string:\n%s\n", rawString)
	
	// String indexing and slicing
	fmt.Printf("   First character: %c\n", str1[0])
	fmt.Printf("   Substring (0:5): %s\n", str1[0:5])
	fmt.Printf("   Substring (7:): %s\n", str1[7:])
	
	// Unicode strings
	unicodeStr := "Hello, 世界!"
	fmt.Printf("   Unicode string: %s\n", unicodeStr)
	fmt.Printf("   Unicode length: %d bytes, %d runes\n", len(unicodeStr), len([]rune(unicodeStr)))
}

// 5. Complex Types
// ================
func complexTypes() {
	fmt.Println("\n5. COMPLEX TYPES:")
	
	// complex64 - 64-bit complex number
	var c64 complex64 = 3 + 4i
	fmt.Printf("   complex64: %v (size: %d bytes)\n", c64, unsafe.Sizeof(c64))
	
	// complex128 - 128-bit complex number
	var c128 complex128 = 3 + 4i
	fmt.Printf("   complex128: %v (size: %d bytes)\n", c128, unsafe.Sizeof(c128))
	
	// Complex number operations
	c1 := 3 + 4i
	c2 := 1 + 2i
	fmt.Printf("   Addition: %v + %v = %v\n", c1, c2, c1+c2)
	fmt.Printf("   Multiplication: %v * %v = %v\n", c1, c2, c1*c2)
	fmt.Printf("   Real part: %f\n", real(c1))
	fmt.Printf("   Imaginary part: %f\n", imag(c1))
	fmt.Printf("   Magnitude: %f\n", math.Sqrt(real(c1)*real(c1)+imag(c1)*imag(c1)))
}

// 6. Byte and Rune Types
// ======================
func byteRuneTypes() {
	fmt.Println("\n6. BYTE AND RUNE TYPES:")
	
	// byte - alias for uint8
	var b byte = 'A'
	fmt.Printf("   byte 'A': %d (%c)\n", b, b)
	
	// rune - alias for int32, represents Unicode code point
	var r rune = '世'
	fmt.Printf("   rune '世': %d (%c)\n", r, r)
	
	// String to byte slice
	str := "Hello"
	bytes := []byte(str)
	fmt.Printf("   String to bytes: %v\n", bytes)
	
	// String to rune slice
	runes := []rune(str)
	fmt.Printf("   String to runes: %v\n", runes)
	
	// Byte slice to string
	byteSlice := []byte{72, 101, 108, 108, 111}
	byteStr := string(byteSlice)
	fmt.Printf("   Bytes to string: %s\n", byteStr)
	
	// Rune slice to string
	runeSlice := []rune{'H', 'e', 'l', 'l', 'o'}
	runeStr := string(runeSlice)
	fmt.Printf("   Runes to string: %s\n", runeStr)
}

// 7. Type Conversions
// ===================
func typeConversions() {
	fmt.Println("\n7. TYPE CONVERSIONS:")
	
	// Integer to integer
	var i int = 42
	var i32 int32 = int32(i)
	var i64 int64 = int64(i)
	fmt.Printf("   int to int32: %d\n", i32)
	fmt.Printf("   int to int64: %d\n", i64)
	
	// Integer to float
	var f float64 = float64(i)
	fmt.Printf("   int to float64: %f\n", f)
	
	// Float to integer (truncation)
	var pi float64 = 3.14159
	var intPi int = int(pi)
	fmt.Printf("   float64 to int: %d (truncated)\n", intPi)
	
	// String conversions
	var numStr string = "123"
	var num int = 123
	fmt.Printf("   String to int: %d\n", num)
	fmt.Printf("   Int to string: %s\n", fmt.Sprintf("%d", num))
	
	// Boolean conversions
	var boolStr string = "true"
	var boolVal bool = (boolStr == "true")
	fmt.Printf("   String to bool: %t\n", boolVal)
}

// 8. Zero Values
// ==============
func zeroValues() {
	fmt.Println("\n8. ZERO VALUES:")
	
	// All types have zero values
	var b bool
	var i int
	var f float64
	var s string
	var c complex128
	
	fmt.Printf("   bool zero value: %t\n", b)
	fmt.Printf("   int zero value: %d\n", i)
	fmt.Printf("   float64 zero value: %f\n", f)
	fmt.Printf("   string zero value: '%s'\n", s)
	fmt.Printf("   complex128 zero value: %v\n", c)
	
	// Zero values are useful for initialization
	var counter int  // starts at 0
	var name string // starts as empty string
	var isReady bool // starts as false
	
	fmt.Printf("   Counter: %d\n", counter)
	fmt.Printf("   Name: '%s'\n", name)
	fmt.Printf("   Is ready: %t\n", isReady)
}

// 9. Type Sizes and Limits
// =========================
func typeSizesAndLimits() {
	fmt.Println("\n9. TYPE SIZES AND LIMITS:")
	
	// Show sizes of all types
	fmt.Printf("   bool: %d bytes\n", unsafe.Sizeof(bool(false)))
	fmt.Printf("   int8: %d bytes\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("   int16: %d bytes\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("   int32: %d bytes\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("   int64: %d bytes\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("   int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("   uint8: %d bytes\n", unsafe.Sizeof(uint8(0)))
	fmt.Printf("   uint16: %d bytes\n", unsafe.Sizeof(uint16(0)))
	fmt.Printf("   uint32: %d bytes\n", unsafe.Sizeof(uint32(0)))
	fmt.Printf("   uint64: %d bytes\n", unsafe.Sizeof(uint64(0)))
	fmt.Printf("   uint: %d bytes\n", unsafe.Sizeof(uint(0)))
	fmt.Printf("   float32: %d bytes\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("   float64: %d bytes\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("   complex64: %d bytes\n", unsafe.Sizeof(complex64(0)))
	fmt.Printf("   complex128: %d bytes\n", unsafe.Sizeof(complex128(0)))
	fmt.Printf("   string: %d bytes\n", unsafe.Sizeof(string("")))
	
	// Show value ranges
	fmt.Printf("   int8 range: %d to %d\n", -128, 127)
	fmt.Printf("   uint8 range: %d to %d\n", 0, 255)
	fmt.Printf("   int16 range: %d to %d\n", -32768, 32767)
	fmt.Printf("   uint16 range: %d to %d\n", 0, 65535)
}

// Helper functions
// ===============
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

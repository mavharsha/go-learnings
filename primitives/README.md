# Go Primitive Types

This folder contains comprehensive examples of Go's primitive data types.

## 📁 Files

- **`go_primitives_simple.go`** - Complete guide to Go primitive types

## 🎯 What You'll Learn

### **Boolean Types**
- `bool` - true/false values
- Boolean operations (&&, ||, !)
- Zero value is `false`

### **Integer Types**
- **Signed integers**: `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned integers**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Special types**: `byte` (uint8), `rune` (int32)
- Integer operations (+, -, *, /, %, ^, <<, >>)

### **Floating-Point Types**
- `float32` - 32-bit floating point
- `float64` - 64-bit floating point (default)
- Scientific notation
- Math constants and functions

### **String Types**
- Basic strings and string operations
- Raw strings (backticks)
- String indexing and slicing
- Unicode strings

### **Complex Types**
- `complex64` - 64-bit complex number
- `complex128` - 128-bit complex number
- Complex number operations

### **Type Conversions**
- Integer to integer conversions
- Integer to float conversions
- String conversions
- Boolean conversions

### **Zero Values**
- All types have zero values
- Useful for initialization

### **Type Sizes and Limits**
- Memory layout of different types
- Value ranges for integer types

## 🚀 How to Run

```bash
cd primitives
go run go_primitives_simple.go
```

## 📚 Key Takeaways

- **Go is statically typed** with strong type safety
- **All types have zero values** for automatic initialization
- **Type conversions are explicit** - no implicit conversions
- **Memory layout is predictable** and efficient
- **Unicode support** with `rune` type for international characters

## 🔗 Related Topics

- **Structs** - See `../structs/` folder
- **Pointers** - See `../pointers/` folder
- **Advanced Concepts** - See `../advanced-concepts/` folder

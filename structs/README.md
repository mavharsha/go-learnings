# Go Structs

This folder contains comprehensive examples of Go's struct types and their usage.

## 📁 Files

- **`go_structs.go`** - Complete guide to Go structs

## 🎯 What You'll Learn

### **Basic Struct Definition and Usage**
- Creating struct types
- Struct field access
- Struct initialization

### **Struct Initialization**
- Zero value initialization
- Field initialization
- Short declaration with field names
- Positional initialization
- Partial initialization
- Using `new()` - returns pointer
- Address of struct literal

### **Struct Fields and Access**
- Accessing fields directly
- Modifying fields
- Field access with pointers
- Dereferencing pointer to access fields

### **Anonymous Structs**
- Structs without type names
- Anonymous struct with embedded fields
- Anonymous struct as function parameter

### **Nested Structs**
- Structs containing other structs
- Accessing nested fields
- Modifying nested fields

### **Struct Methods**
- Value receiver methods
- Pointer receiver methods
- Method overriding
- Method sets

### **Struct Embedding (Composition)**
- Embedded structs
- Method inheritance
- Field access through embedding
- Method overriding

### **Struct Tags**
- JSON tags
- XML tags
- Hidden fields
- Omitempty tags

### **Struct Comparison**
- Comparing structs
- Structs with comparable fields
- Non-comparable structs

### **Struct Memory Layout**
- Memory alignment
- Field offsets
- Automatic alignment handling

## 🚀 How to Run

```bash
cd structs
go run go_structs.go
```

## 📚 Key Takeaways

- **Structs are value types** - copied when passed by value
- **Methods can have value or pointer receivers** - choose based on use case
- **Embedding provides composition** - not inheritance
- **Struct tags provide metadata** for serialization
- **Memory layout is optimized** by Go compiler

## 🔗 Related Topics

- **Primitive Types** - See `../primitives/` folder
- **Pointers** - See `../pointers/` folder
- **Advanced Concepts** - See `../advanced-concepts/` folder

# Go Language Fundamentals - Complete Learning Resource

## 🎯 **What You Need to Know About Go**

I've created a comprehensive learning resource covering all the essential Go concepts you need to know as a beginner.

## 📚 **Working Learning Files**

### **1. `go_primitives_simple.go` - Primitive Types** ✅
- **Boolean types**: `bool` (true/false)
- **Integer types**: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Floating-point types**: `float32`, `float64`
- **String types**: `string`
- **Byte and rune types**: `byte` (uint8), `rune` (int32)
- **Type conversions** and **zero values**
- **Type sizes** and **memory layout**

### **2. `go_structs.go` - Structs** ✅
- **Basic struct definition** and usage
- **Struct initialization** (multiple ways)
- **Struct fields** and access patterns
- **Anonymous structs**
- **Nested structs**
- **Struct methods** (value vs pointer receivers)
- **Struct embedding** (composition)
- **Struct tags** (JSON, XML)
- **Struct comparison** and **memory layout**

### **3. `go_pointers_simple.go` - Pointers** ✅
- **Basic pointer concepts** (declaration, address, dereference)
- **Pointer operations** (comparison, pointer to pointer)
- **Pointers and functions** (pass by reference vs value)
- **Pointers and structs** (method receivers)
- **Pointers and arrays**
- **Pointer safety** and **best practices**

### **4. `go_other_concepts_simple.go` - Advanced Concepts** ✅
- **Interfaces** (definition, implementation, composition)
- **Methods** (value vs pointer receivers, method sets)
- **Channels** (buffered/unbuffered, send/receive)
- **Goroutines** (concurrency, communication)
- **Maps** (creation, access, iteration, deletion)
- **Slices** (creation, append, copy, 2D slices)
- **Functions as values** (higher-order functions, closures)
- **Type assertions** and **type switches**
- **Error handling** (custom errors, multiple return values)

## 🚀 **How to Use These Resources**

### **Run the Examples:**
```bash
go run go_primitives_simple.go
go run go_structs.go
go run go_pointers_simple.go
go run go_other_concepts_simple.go
```

### **Study Each File:**
1. **Start with** `go_primitives_simple.go` - Learn basic data types
2. **Move to** `go_structs.go` - Understand structured data
3. **Learn** `go_pointers_simple.go` - Master memory management
4. **Explore** `go_other_concepts_simple.go` - Advanced Go features

## 📖 **Complete Go Fundamentals Overview**

### **🔢 Primitive Types**

#### **Boolean**
```go
var b bool = true
var b2 bool  // zero value is false
```

#### **Integers**
```go
var i int = 42           // Platform-dependent (32 or 64 bits)
var i8 int8 = 127        // 8-bit signed integer
var i16 int16 = 32767    // 16-bit signed integer
var i32 int32 = 2147483647  // 32-bit signed integer
var i64 int64 = 9223372036854775807  // 64-bit signed integer

var u uint = 42          // Unsigned integer
var u8 uint8 = 255       // 8-bit unsigned integer
var u16 uint16 = 65535   // 16-bit unsigned integer
var u32 uint32 = 4294967295  // 32-bit unsigned integer
var u64 uint64 = 18446744073709551615  // 64-bit unsigned integer

var b byte = 255         // byte is alias for uint8
var r rune = 'A'         // rune is alias for int32 (Unicode code point)
```

#### **Floating-Point**
```go
var f32 float32 = 3.14159    // 32-bit floating point
var f64 float64 = 3.141592653589793  // 64-bit floating point (default)
```

#### **Strings**
```go
var s string = "Hello, World!"
var s2 string  // zero value is empty string ""
```

### **🏗️ Structs**

#### **Basic Struct**
```go
type Person struct {
    Name string
    Age  int
    City string
}

// Initialization
var person1 Person
person1.Name = "Alice"
person1.Age = 30

person2 := Person{
    Name: "Bob",
    Age:  25,
    City: "London",
}
```

#### **Struct Methods**
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver method
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver method
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

#### **Struct Embedding**
```go
type Animal struct {
    Name string
    Age  int
}

type Dog struct {
    Animal  // Embedded struct
    Breed   string
    IsGood  bool
}
```

### **📍 Pointers**

#### **Basic Pointers**
```go
var x int = 42
var px *int = &x  // Get address of x
fmt.Printf("x = %d, &x = %p\n", x, px)
fmt.Printf("*px = %d\n", *px)  // Dereference pointer

*px = 100  // Modify x through pointer
fmt.Printf("x = %d\n", x)  // x is now 100
```

#### **Pointers and Functions**
```go
// Function that takes a pointer
func modifyValue(ptr *int) {
    *ptr = 100
}

// Function that returns a pointer
func createPointer(value int) *int {
    return &value  // This escapes to heap
}

// Use pointers
x := 42
modifyValue(&x)  // Pass address of x
fmt.Printf("x = %d\n", x)  // x is now 100
```

#### **Pointers and Structs**
```go
type Point struct {
    X, Y int
}

func movePoint(p *Point, dx, dy int) {
    p.X += dx
    p.Y += dy
}

point := Point{X: 10, Y: 20}
movePoint(&point, 5, 10)
fmt.Printf("Point: %+v\n", point)  // {X:15 Y:30}
```

### **🔌 Interfaces**

#### **Basic Interface**
```go
type Writer interface {
    Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
    fmt.Printf("ConsoleWriter: %s\n", string(data))
    return len(data), nil
}

// Use interface
var writer Writer = ConsoleWriter{}
writer.Write([]byte("Hello, World!"))
```

#### **Interface Composition**
```go
type Reader interface {
    Read([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

### **🔄 Channels and Goroutines**

#### **Channels**
```go
// Unbuffered channel
ch := make(chan int)

// Buffered channel
ch2 := make(chan string, 3)

// Send and receive
go func() {
    ch <- 42
}()

val := <-ch
fmt.Printf("Received: %d\n", val)
```

#### **Goroutines**
```go
// Basic goroutine
go func() {
    fmt.Println("Hello from goroutine!")
}()

// Goroutine with parameters
go func(id int) {
    fmt.Printf("Goroutine %d: Running\n", id)
}(1)
```

### **🗺️ Maps**

#### **Basic Map Operations**
```go
// Create map
m := make(map[string]int)
m2 := map[string]int{
    "apple":  5,
    "banana": 3,
    "orange": 8,
}

// Add/update values
m["key1"] = 42
m["key2"] = 100

// Access values
val := m["key1"]
fmt.Printf("m[key1]: %d\n", val)

// Check if key exists
val, exists := m["key1"]
if exists {
    fmt.Printf("Key exists: %d\n", val)
}

// Delete keys
delete(m, "key1")

// Iterate over map
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
```

### **📊 Slices**

#### **Basic Slice Operations**
```go
// Create slices
slice1 := make([]int, 5)        // length 5, capacity 5
slice2 := make([]int, 3, 10)    // length 3, capacity 10
slice3 := []int{1, 2, 3, 4, 5} // slice literal

// Append to slice
slice3 = append(slice3, 6, 7, 8)

// Slice operations
fmt.Printf("slice3[2:5]: %v\n", slice3[2:5])
fmt.Printf("slice3[:3]: %v\n", slice3[:3])
fmt.Printf("slice3[3:]: %v\n", slice3[3:])

// Copy slices
slice4 := make([]int, len(slice3))
copy(slice4, slice3)
```

### **🔧 Functions as Values**

#### **Function Types**
```go
type Operation func(int, int) int

add := func(a, b int) int {
    return a + b
}

multiply := func(a, b int) int {
    return a * b
}

var op Operation = add
fmt.Printf("add(5, 3) = %d\n", op(5, 3))

op = multiply
fmt.Printf("multiply(5, 3) = %d\n", op(5, 3))
```

#### **Higher-Order Functions**
```go
func mapInts(slice []int, fn func(int) int) []int {
    result := make([]int, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

numbers := []int{1, 2, 3, 4, 5}
doubled := mapInts(numbers, func(x int) int { return x * 2 })
fmt.Printf("Doubled: %v\n", doubled)
```

### **🎭 Type Assertions and Type Switches**

#### **Type Assertions**
```go
var i interface{} = 42

// Safe type assertion
if val, ok := i.(int); ok {
    fmt.Printf("i is int: %d\n", val)
}
```

#### **Type Switches**
```go
func processValue(v interface{}) {
    switch val := v.(type) {
    case int:
        fmt.Printf("Integer: %d\n", val)
    case string:
        fmt.Printf("String: %s\n", val)
    case bool:
        fmt.Printf("Boolean: %t\n", val)
    default:
        fmt.Printf("Unknown type: %T\n", val)
    }
}
```

### **🛠️ Error Handling**

#### **Custom Errors**
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Handle error
result, err := divide(10, 2)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("10 / 2 = %d\n", result)
}
```

## 🎯 **Key Go Concepts Summary**

### **✅ What You Need to Know:**

1. **Primitive Types**: `bool`, `int`, `float64`, `string`, `byte`, `rune`
2. **Structs**: Custom data types with fields and methods
3. **Pointers**: Memory addresses for efficient data manipulation
4. **Interfaces**: Contracts for method implementation
5. **Channels**: Communication between goroutines
6. **Goroutines**: Lightweight concurrent execution
7. **Maps**: Key-value data structures
8. **Slices**: Dynamic arrays with powerful operations
9. **Functions**: First-class citizens, can be passed around
10. **Error Handling**: Explicit error returns, no exceptions

### **🚀 Best Practices:**

- **Use pointers** for large structs to avoid copying
- **Use value receivers** for small structs, pointer receivers for large ones
- **Check for nil pointers** before dereferencing
- **Use interfaces** for flexible, testable code
- **Handle errors explicitly** - don't ignore them
- **Use channels** for goroutine communication
- **Pre-allocate slices** with known capacity
- **Use meaningful names** for variables and functions
- **Follow Go conventions** (camelCase, exported vs unexported)

### **📚 Learning Path:**

1. **Start with primitives** - Understand basic data types
2. **Learn structs** - Create custom data structures
3. **Master pointers** - Understand memory management
4. **Explore interfaces** - Write flexible, testable code
5. **Practice concurrency** - Use goroutines and channels
6. **Build projects** - Apply what you've learned

## 🎉 **Success!**

All files are working correctly and demonstrate the essential Go concepts you need to know:

- ✅ **Primitive Types** - Basic data types and operations
- ✅ **Structs** - Custom data structures with methods
- ✅ **Pointers** - Memory management and efficiency
- ✅ **Advanced Concepts** - Interfaces, concurrency, error handling

This comprehensive resource gives you everything you need to understand Go fundamentals and start building Go applications! 🚀

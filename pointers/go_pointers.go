package main

import (
	"fmt"
	"unsafe"
)

// Go Pointers - Complete Guide
// ===========================
// This file demonstrates Go pointers with comprehensive examples

func main() {
	fmt.Println("=== Go Pointers ===")
	
	// 1. Basic pointer concepts
	basicPointers()
	
	// 2. Pointer operations
	pointerOperations()
	
	// 3. Pointers to different types
	pointersToTypes()
	
	// 4. Pointers and functions
	pointersAndFunctions()
	
	// 5. Pointers and structs
	pointersAndStructs()
	
	// 6. Pointers and arrays/slices
	pointersAndArrays()
	
	// 7. Pointer arithmetic (limited in Go)
	pointerArithmetic()
	
	// 8. Pointers and memory management
	pointersAndMemory()
	
	// 9. Common pointer patterns
	commonPointerPatterns()
	
	// 10. Pointer safety and best practices
	pointerSafety()
}

// 1. Basic Pointer Concepts
// ==========================
func basicPointers() {
	fmt.Println("\n1. BASIC POINTER CONCEPTS:")
	
	// Declaring pointers
	var p1 *int        // pointer to int
	var p2 *string     // pointer to string
	var p3 *float64    // pointer to float64
	
	fmt.Printf("   p1 (nil): %v\n", p1)
	fmt.Printf("   p2 (nil): %v\n", p2)
	fmt.Printf("   p3 (nil): %v\n", p3)
	
	// Getting address of variables
	x := 42
	y := "Hello"
	z := 3.14
	
	px := &x  // address of x
	py := &y  // address of y
	pz := &z  // address of z
	
	fmt.Printf("   x = %d, &x = %p\n", x, px)
	fmt.Printf("   y = %s, &y = %p\n", y, py)
	fmt.Printf("   z = %f, &z = %p\n", z, pz)
	
	// Dereferencing pointers
	fmt.Printf("   *px = %d\n", *px)
	fmt.Printf("   *py = %s\n", *py)
	fmt.Printf("   *pz = %f\n", *pz)
	
	// Modifying values through pointers
	*px = 100
	*py = "World"
	*pz = 2.71
	
	fmt.Printf("   After modification:\n")
	fmt.Printf("   x = %d\n", x)
	fmt.Printf("   y = %s\n", y)
	fmt.Printf("   z = %f\n", z)
}

// 2. Pointer Operations
// =====================
func pointerOperations() {
	fmt.Println("\n2. POINTER OPERATIONS:")
	
	// Pointer comparison
	x := 42
	y := 42
	px := &x
	py := &y
	pz := &x
	
	fmt.Printf("   x = %d, y = %d\n", x, y)
	fmt.Printf("   px == py: %t (different addresses)\n", px == py)
	fmt.Printf("   px == pz: %t (same address)\n", px == pz)
	fmt.Printf("   *px == *py: %t (same values)\n", *px == *py)
	
	// Pointer to pointer
	ppx := &px  // pointer to pointer to int
	fmt.Printf("   ppx = %p (address of px)\n", ppx)
	fmt.Printf("   *ppx = %p (value of px)\n", *ppx)
	fmt.Printf("   **ppx = %d (value of x)\n", **ppx)
	
	// Modifying through double pointer
	**ppx = 200
	fmt.Printf("   After **ppx = 200: x = %d\n", x)
	
	// Pointer to zero value
	var zero int
	pzero := &zero
	fmt.Printf("   Zero value: %d\n", *pzero)
	
	// Checking for nil pointers
	var nilPtr *int
	fmt.Printf("   nilPtr == nil: %t\n", nilPtr == nil)
	
	// Safe dereferencing (check for nil first)
	if nilPtr != nil {
		fmt.Printf("   *nilPtr = %d\n", *nilPtr)
	} else {
		fmt.Printf("   nilPtr is nil, cannot dereference\n")
	}
}

// 3. Pointers to Different Types
// ==============================
func pointersToTypes() {
	fmt.Println("\n3. POINTERS TO DIFFERENT TYPES:")
	
	// Pointers to primitive types
	var i int = 42
	var f float64 = 3.14
	var s string = "Hello"
	var b bool = true
	
	pi := &i
	pf := &f
	ps := &s
	pb := &b
	
	fmt.Printf("   int pointer: %p -> %d\n", pi, *pi)
	fmt.Printf("   float64 pointer: %p -> %f\n", pf, *pf)
	fmt.Printf("   string pointer: %p -> %s\n", ps, *ps)
	fmt.Printf("   bool pointer: %p -> %t\n", pb, *pb)
	
	// Pointers to structs
	type Person struct {
		Name string
		Age  int
	}
	
	person := Person{Name: "Alice", Age: 30}
	pperson := &person
	
	fmt.Printf("   Person: %+v\n", *pperson)
	fmt.Printf("   Person name: %s\n", (*pperson).Name)
	fmt.Printf("   Person age: %d\n", (*pperson).Age)
	
	// Shorthand for struct pointers
	fmt.Printf("   Person name (shorthand): %s\n", pperson.Name)
	fmt.Printf("   Person age (shorthand): %d\n", pperson.Age)
	
	// Pointers to arrays
	arr := [3]int{1, 2, 3}
	parr := &arr
	
	fmt.Printf("   Array: %v\n", *parr)
	fmt.Printf("   Array[0]: %d\n", (*parr)[0])
	fmt.Printf("   Array[0] (shorthand): %d\n", parr[0])
	
	// Pointers to slices
	slice := []int{4, 5, 6}
	pslice := &slice
	
	fmt.Printf("   Slice: %v\n", *pslice)
	fmt.Printf("   Slice[0]: %d\n", (*pslice)[0])
	fmt.Printf("   Slice[0] (shorthand): %d\n", pslice[0])
}

// 4. Pointers and Functions
// =========================
func pointersAndFunctions() {
	fmt.Println("\n4. POINTERS AND FUNCTIONS:")
	
	// Test value vs pointer parameters
	x := 42
	fmt.Printf("   Original x: %d\n", x)
	
	modifyValueCopy(x)
	fmt.Printf("   After modifyValueCopy: %d\n", x)
	
	modifyValue(&x)
	fmt.Printf("   After modifyValue: %d\n", x)
	
	// Test value vs pointer returns
	val := createValue(50)
	ptr := createPointer(60)
	
	fmt.Printf("   createValue(50): %d\n", val)
	fmt.Printf("   createPointer(60): %d\n", *ptr)
	
	// Function that modifies struct through pointer
	point := Point{X: 10, Y: 20}
	fmt.Printf("   Original point: %+v\n", point)
	
	movePoint(&point, 5, 10)
	fmt.Printf("   After movePoint: %+v\n", point)
}

// 5. Pointers and Structs
// ========================
func pointersAndStructs() {
	fmt.Println("\n5. POINTERS AND STRUCTS:")
	
	type Rectangle struct {
		Width  float64
		Height float64
	}
	
	// Method with value receiver
	func (r Rectangle) Area() float64 {
		return r.Width * r.Height
	}
	
	// Method with pointer receiver
	func (r *Rectangle) SetDimensions(width, height float64) {
		r.Width = width
		r.Height = height
	}
	
	// Method with pointer receiver that modifies
	func (r *Rectangle) Scale(factor float64) {
		r.Width *= factor
		r.Height *= factor
	}
	
	// Create rectangle
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle: %+v\n", rect)
	fmt.Printf("   Area: %f\n", rect.Area())
	
	// Use pointer receiver method
	rect.SetDimensions(15, 8)
	fmt.Printf("   After SetDimensions: %+v\n", rect)
	
	// Scale the rectangle
	rect.Scale(2.0)
	fmt.Printf("   After Scale(2.0): %+v\n", rect)
	
	// Pointer to struct
	rectPtr := &rect
	fmt.Printf("   Via pointer: %+v\n", *rectPtr)
	
	// Method calls on pointer
	fmt.Printf("   Area via pointer: %f\n", rectPtr.Area())
	rectPtr.Scale(0.5)
	fmt.Printf("   After Scale(0.5): %+v\n", *rectPtr)
}

// 6. Pointers and Arrays
// =======================
func pointersAndArrays() {
	fmt.Println("\n6. POINTERS AND ARRAYS:")
	
	// Array and pointer to array
	arr := [5]int{1, 2, 3, 4, 5}
	parr := &arr
	
	fmt.Printf("   Array: %v\n", arr)
	fmt.Printf("   Array via pointer: %v\n", *parr)
	
	// Modifying array through pointer
	(*parr)[0] = 100
	fmt.Printf("   After (*parr)[0] = 100: %v\n", arr)
	
	// Shorthand for array pointer access
	parr[1] = 200
	fmt.Printf("   After parr[1] = 200: %v\n", arr)
	
	// Pointer to array element
	pelem := &arr[2]
	*pelem = 300
	fmt.Printf("   After *pelem = 300: %v\n", arr)
	
	// Slice and pointer to slice
	slice := []int{10, 20, 30, 40, 50}
	pslice := &slice
	
	fmt.Printf("   Slice: %v\n", slice)
	fmt.Printf("   Slice via pointer: %v\n", *pslice)
	
	// Modifying slice through pointer
	(*pslice)[0] = 1000
	fmt.Printf("   After (*pslice)[0] = 1000: %v\n", slice)
	
	// Appending to slice through pointer
	*pslice = append(*pslice, 60, 70)
	fmt.Printf("   After append: %v\n", slice)
}

// 7. Pointer Arithmetic (Limited in Go)
// =======================================
func pointerArithmetic() {
	fmt.Println("\n7. POINTER ARITHMETIC (LIMITED IN GO):")
	
	// Go doesn't allow pointer arithmetic like C
	// But we can demonstrate what's possible
	
	arr := [5]int{10, 20, 30, 40, 50}
	
	// Get pointer to first element
	p1 := &arr[0]
	p2 := &arr[1]
	p3 := &arr[2]
	
	fmt.Printf("   arr[0] address: %p\n", p1)
	fmt.Printf("   arr[1] address: %p\n", p2)
	fmt.Printf("   arr[2] address: %p\n", p3)
	
	// Calculate address differences
	diff1 := uintptr(unsafe.Pointer(p2)) - uintptr(unsafe.Pointer(p1))
	diff2 := uintptr(unsafe.Pointer(p3)) - uintptr(unsafe.Pointer(p2))
	
	fmt.Printf("   Address difference (p2-p1): %d bytes\n", diff1)
	fmt.Printf("   Address difference (p3-p2): %d bytes\n", diff2)
	
	// Unsafe pointer arithmetic (not recommended)
	unsafePtr := unsafe.Pointer(p1)
	nextPtr := unsafe.Pointer(uintptr(unsafePtr) + unsafe.Sizeof(int(0)))
	nextInt := (*int)(nextPtr)
	
	fmt.Printf("   Next element via unsafe: %d\n", *nextInt)
	fmt.Printf("   Should be arr[1]: %d\n", arr[1])
}

// 8. Pointers and Memory Management
// ================================
func pointersAndMemory() {
	fmt.Println("\n8. POINTERS AND MEMORY MANAGEMENT:")
	
	// Go has garbage collection, but we can still manage memory
	
	// 1. Stack allocation (automatic)
	func stackAllocation() {
		x := 42  // Stack allocated
		px := &x // Pointer to stack variable
		fmt.Printf("   Stack variable: %d\n", *px)
	} // x is automatically cleaned up
	
	stackAllocation()
	
	// 2. Heap allocation (automatic)
	func heapAllocation() *int {
		x := 100  // This escapes to heap
		return &x // Returning address
	}
	
	heapPtr := heapAllocation()
	fmt.Printf("   Heap variable: %d\n", *heapPtr)
	// heapPtr will be garbage collected when no longer referenced
	
	// 3. Memory addresses
	x := 42
	y := 43
	px := &x
	py := &y
	
	fmt.Printf("   x address: %p\n", px)
	fmt.Printf("   y address: %p\n", py)
	fmt.Printf("   Address difference: %d bytes\n", uintptr(unsafe.Pointer(py))-uintptr(unsafe.Pointer(px)))
	
	// 4. Pointer size
	fmt.Printf("   Pointer size: %d bytes\n", unsafe.Sizeof(px))
}

// 9. Common Pointer Patterns
// ===========================
func commonPointerPatterns() {
	fmt.Println("\n9. COMMON POINTER PATTERNS:")
	
	// 1. Optional values (nil pointer)
	func processOptionalValue(ptr *int) {
		if ptr != nil {
			fmt.Printf("   Processing value: %d\n", *ptr)
		} else {
			fmt.Printf("   No value provided\n")
		}
	}
	
	processOptionalValue(nil)
	val := 42
	processOptionalValue(&val)
	
	// 2. Builder pattern
	type StringBuilder struct {
		parts []string
	}
	
	func (sb *StringBuilder) Add(s string) *StringBuilder {
		sb.parts = append(sb.parts, s)
		return sb
	}
	
	func (sb *StringBuilder) String() string {
		result := ""
		for _, part := range sb.parts {
			result += part
		}
		return result
	}
	
	builder := &StringBuilder{}
	result := builder.Add("Hello").Add(" ").Add("World").String()
	fmt.Printf("   Builder result: %s\n", result)
	
	// 3. Function that modifies multiple values
	func swap(a, b *int) {
		*a, *b = *b, *a
	}
	
	x, y := 10, 20
	fmt.Printf("   Before swap: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("   After swap: x=%d, y=%d\n", x, y)
	
	// 4. Pointer to function
	func add(a, b int) int {
		return a + b
	}
	
	func multiply(a, b int) int {
		return a * b
	}
	
	var operation func(int, int) int
	operation = add
	fmt.Printf("   add(5, 3) = %d\n", operation(5, 3))
	
	operation = multiply
	fmt.Printf("   multiply(5, 3) = %d\n", operation(5, 3))
}

// 10. Pointer Safety and Best Practices
// =====================================
func pointerSafety() {
	fmt.Println("\n10. POINTER SAFETY AND BEST PRACTICES:")
	
	// 1. Always check for nil pointers
	func safeDereference(ptr *int) {
		if ptr != nil {
			fmt.Printf("   Safe dereference: %d\n", *ptr)
		} else {
			fmt.Printf("   Cannot dereference nil pointer\n")
		}
	}
	
	var nilPtr *int
	safeDereference(nilPtr)
	
	val := 42
	safeDereference(&val)
	
	// 2. Use pointers for large structs to avoid copying
	type LargeStruct struct {
		Data [1000]int
	}
	
	func processLargeStruct(ls *LargeStruct) {
		ls.Data[0] = 999
	}
	
	large := LargeStruct{}
	processLargeStruct(&large)
	fmt.Printf("   Large struct first element: %d\n", large.Data[0])
	
	// 3. Use pointers when you need to modify the original
	func modifyInPlace(ptr *int) {
		*ptr *= 2
	}
	
	num := 21
	modifyInPlace(&num)
	fmt.Printf("   Modified in place: %d\n", num)
	
	// 4. Be careful with pointer lifetimes
	func getPointer() *int {
		x := 42
		return &x  // This escapes to heap
	}
	
	ptr := getPointer()
	fmt.Printf("   Escaped pointer: %d\n", *ptr)
	
	// 5. Use value receivers for small structs, pointer receivers for large ones
	
	small := SmallStruct{Value: 100}
	fmt.Printf("   Small struct value: %d\n", small.GetValue())
	small.SetValue(200)
	fmt.Printf("   After SetValue: %d\n", small.GetValue())
}

// Helper functions
// ================
func modifyValue(ptr *int) {
	*ptr = 100
}

func modifyValueCopy(val int) {
	val = 200  // This doesn't affect the original
}

func createPointer(value int) *int {
	return &value  // This escapes to heap
}

func createValue(value int) int {
	return value
}

func movePoint(p *Point, dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Types
type Point struct {
	X, Y int
}

type SmallStruct struct {
	Value int
}

func (s SmallStruct) GetValue() int {
	return s.Value
}

func (s *SmallStruct) SetValue(v int) {
	s.Value = v
}

package main

import (
	"fmt"
	"unsafe"
)

// Go Pointers - Simple Guide
// ==========================

func main() {
	fmt.Println("=== Go Pointers ===")
	
	// 1. Basic pointer concepts
	basicPointers()
	
	// 2. Pointers and functions
	pointersAndFunctions()
	
	// 3. Pointers and structs
	pointersAndStructs()
	
	// 4. Pointers and arrays
	pointersAndArrays()
	
	// 5. Pointer safety
	pointerSafety()
}

// 1. Basic Pointer Concepts
// ==========================
func basicPointers() {
	fmt.Println("\n1. BASIC POINTER CONCEPTS:")
	
	// Declaring pointers
	var p1 *int        // pointer to int
	var p2 *string     // pointer to string
	
	fmt.Printf("   p1 (nil): %v\n", p1)
	fmt.Printf("   p2 (nil): %v\n", p2)
	
	// Getting address of variables
	x := 42
	y := "Hello"
	
	px := &x  // address of x
	py := &y  // address of y
	
	fmt.Printf("   x = %d, &x = %p\n", x, px)
	fmt.Printf("   y = %s, &y = %p\n", y, py)
	
	// Dereferencing pointers
	fmt.Printf("   *px = %d\n", *px)
	fmt.Printf("   *py = %s\n", *py)
	
	// Modifying values through pointers
	*px = 100
	*py = "World"
	
	fmt.Printf("   After modification:\n")
	fmt.Printf("   x = %d\n", x)
	fmt.Printf("   y = %s\n", y)
}

// 2. Pointers and Functions
// =========================
func pointersAndFunctions() {
	fmt.Println("\n2. POINTERS AND FUNCTIONS:")
	
	// Function that takes a pointer
	x := 42
	fmt.Printf("   Original x: %d\n", x)
	
	modifyValue(&x)
	fmt.Printf("   After modifyValue: %d\n", x)
	
	// Function that returns a pointer
	ptr := createPointer(60)
	fmt.Printf("   createPointer(60): %d\n", *ptr)
}

// 3. Pointers and Structs
// ========================
func pointersAndStructs() {
	fmt.Println("\n3. POINTERS AND STRUCTS:")
	
	// Create rectangle and use methods
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle: %+v\n", rect)
	fmt.Printf("   Area: %f\n", rect.Area())
	
	// Use pointer receiver method
	rect.SetDimensions(15, 8)
	fmt.Printf("   After SetDimensions: %+v\n", rect)
	
	// Scale the rectangle
	rect.Scale(2.0)
	fmt.Printf("   After Scale(2.0): %+v\n", rect)
}

// 4. Pointers and Arrays
// =======================
func pointersAndArrays() {
	fmt.Println("\n4. POINTERS AND ARRAYS:")
	
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
}

// 5. Pointer Safety
// =================
func pointerSafety() {
	fmt.Println("\n5. POINTER SAFETY:")
	
	// Always check for nil pointers
	var nilPtr *int
	fmt.Printf("   nilPtr == nil: %t\n", nilPtr == nil)
	
	// Safe dereferencing (check for nil first)
	if nilPtr != nil {
		fmt.Printf("   *nilPtr = %d\n", *nilPtr)
	} else {
		fmt.Printf("   nilPtr is nil, cannot dereference\n")
	}
	
	// Pointer size
	x := 42
	px := &x
	fmt.Printf("   Pointer size: %d bytes\n", unsafe.Sizeof(px))
}

// Helper functions
// ================
func modifyValue(ptr *int) {
	*ptr = 100
}

func createPointer(value int) *int {
	return &value  // This escapes to heap
}

// Types
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) SetDimensions(width, height float64) {
	r.Width = width
	r.Height = height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

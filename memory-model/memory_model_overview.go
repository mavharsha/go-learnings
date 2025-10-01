package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// Go Memory Model Overview
// =======================

// Go's memory model is based on two main areas:
// 1. Stack - Fast, automatic memory management
// 2. Heap - Slower, garbage collected memory

func main() {
	fmt.Println("=== Go Memory Model Overview ===")
	
	// Basic memory concepts
	explainBasicConcepts()
	
	// Stack allocation examples
	demonstrateStackAllocation()
	
	// Heap allocation examples
	demonstrateHeapAllocation()
	
	// Escape analysis
	demonstrateEscapeAnalysis()
	
	// Performance comparison
	comparePerformance()
}

// Basic Concepts
// =============
func explainBasicConcepts() {
	fmt.Println("\n1. BASIC CONCEPTS:")
	fmt.Println("   Stack: Fast, LIFO (Last In, First Out) memory")
	fmt.Println("   - Automatic allocation/deallocation")
	fmt.Println("   - Limited size (typically 1-8MB per goroutine)")
	fmt.Println("   - No garbage collection overhead")
	fmt.Println("   - Variables are automatically cleaned up when function returns")
	
	fmt.Println("\n   Heap: Slower, garbage collected memory")
	fmt.Println("   - Managed by Go's garbage collector")
	fmt.Println("   - Can grow dynamically")
	fmt.Println("   - Variables live beyond function scope")
	fmt.Println("   - Requires garbage collection cycles")
}

// Stack Allocation Examples
// ========================
func demonstrateStackAllocation() {
	fmt.Println("\n2. STACK ALLOCATION EXAMPLES:")
	
	// Simple variables are typically allocated on the stack
	var x int = 42
	var y float64 = 3.14
	var z string = "Hello"
	
	fmt.Printf("   x (int): %d, size: %d bytes\n", x, unsafe.Sizeof(x))
	fmt.Printf("   y (float64): %f, size: %d bytes\n", y, unsafe.Sizeof(y))
	fmt.Printf("   z (string): %s, size: %d bytes\n", z, unsafe.Sizeof(z))
	
	// Arrays with known size at compile time
	var arr [100]int
	fmt.Printf("   arr [100]int: size: %d bytes\n", unsafe.Sizeof(arr))
	
	// Structs with simple fields
	type Point struct {
		X, Y int
	}
	var p Point = Point{X: 10, Y: 20}
	fmt.Printf("   Point struct: %+v, size: %d bytes\n", p, unsafe.Sizeof(p))
}

// Heap Allocation Examples
// =======================
func demonstrateHeapAllocation() {
	fmt.Println("\n3. HEAP ALLOCATION EXAMPLES:")
	
	// Using new() - allocates on heap
	ptr := new(int)
	*ptr = 100
	fmt.Printf("   new(int): %d, address: %p\n", *ptr, ptr)
	
	// Using make() for slices, maps, channels
	slice := make([]int, 1000)
	fmt.Printf("   make([]int, 1000): length: %d, capacity: %d\n", len(slice), cap(slice))
	
	// Maps are always allocated on heap
	m := make(map[string]int)
	m["key"] = 42
	fmt.Printf("   map: %v\n", m)
	
	// Channels are always allocated on heap
	ch := make(chan int, 10)
	fmt.Printf("   channel: %p\n", ch)
}

// Escape Analysis
// ==============
func demonstrateEscapeAnalysis() {
	fmt.Println("\n4. ESCAPE ANALYSIS:")
	fmt.Println("   Go compiler determines if variables 'escape' to heap")
	
	// This function returns a pointer, so the variable escapes to heap
	escaped := createPointer()
	fmt.Printf("   Escaped variable: %d\n", *escaped)
	
	// This function returns a value, so it stays on stack
	stackValue := createValue()
	fmt.Printf("   Stack value: %d\n", stackValue)
	
	// Large variables might escape to heap
	largeArray := createLargeArray()
	fmt.Printf("   Large array size: %d\n", len(largeArray))
}

// Function that causes escape to heap
func createPointer() *int {
	x := 42  // This escapes to heap because we return its address
	return &x
}

// Function that keeps value on stack
func createValue() int {
	x := 42  // This stays on stack
	return x
}

// Function that might cause escape due to size
func createLargeArray() []int {
	// Large slices might escape to heap
	return make([]int, 10000)
}

// Performance Comparison
// =====================
func comparePerformance() {
	fmt.Println("\n5. PERFORMANCE COMPARISON:")
	
	// Stack allocation is very fast
	stackAllocation()
	
	// Heap allocation is slower due to GC
	heapAllocation()
	
	// Show memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("   Current heap size: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("   Number of GC cycles: %d\n", m.NumGC)
}

func stackAllocation() {
	// Fast stack allocation
	for i := 0; i < 1000; i++ {
		var x int = i
		_ = x  // Prevent optimization
	}
	fmt.Println("   Stack allocation: Very fast")
}

func heapAllocation() {
	// Slower heap allocation
	for i := 0; i < 1000; i++ {
		ptr := new(int)
		*ptr = i
		_ = ptr  // Prevent optimization
	}
	fmt.Println("   Heap allocation: Slower due to GC")
}

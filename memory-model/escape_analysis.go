package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// Go Escape Analysis Deep Dive
// ===========================
// Escape analysis is Go's compile-time optimization that determines
// whether variables should be allocated on the stack or heap.

func main() {
	fmt.Println("=== Go Escape Analysis ===")
	
	// Understanding escape analysis
	explainEscapeAnalysis()
	
	// Examples of variables that stay on stack
	stackExamples()
	
	// Examples of variables that escape to heap
	heapExamples()
	
	// How to check escape analysis
	checkEscapeAnalysis()
	
	// Optimization techniques
	optimizationTechniques()
}

// Understanding Escape Analysis
// ============================
func explainEscapeAnalysis() {
	fmt.Println("\n1. WHAT IS ESCAPE ANALYSIS?")
	fmt.Println("   Escape analysis determines if a variable's lifetime")
	fmt.Println("   extends beyond the function where it's declared.")
	fmt.Println("   - If variable lifetime > function lifetime → HEAP")
	fmt.Println("   - If variable lifetime = function lifetime → STACK")
	
	fmt.Println("\n   RULES FOR ESCAPE:")
	fmt.Println("   ✓ Returning address of local variable")
	fmt.Println("   ✓ Storing in global variable")
	fmt.Println("   ✓ Storing in heap-allocated structure")
	fmt.Println("   ✓ Passing to function that might store it")
	fmt.Println("   ✓ Interface method calls")
	fmt.Println("   ✓ Large variables (size threshold)")
	fmt.Println("   ✓ Dynamic stack growth")
}

// Stack Examples (Variables that DON'T escape)
// ===========================================
func stackExamples() {
	fmt.Println("\n2. VARIABLES THAT STAY ON STACK:")
	
	// Simple local variables
	simpleVariables()
	
	// Function parameters and return values
	functionStackExamples()
	
	// Struct fields (when struct doesn't escape)
	structStackExamples()
}

func simpleVariables() {
	fmt.Println("   Simple local variables:")
	
	var a int = 42
	var b float64 = 3.14
	var c string = "Hello"
	var d bool = true
	
	fmt.Printf("     a=%d, b=%f, c=%s, d=%t\n", a, b, c, d)
	fmt.Printf("     All stay on stack (no addresses taken)\n")
}

func functionStackExamples() {
	fmt.Println("   Function parameters and return values:")
	
	// Value parameters - stay on stack
	result := add(10, 20)
	fmt.Printf("     add(10, 20) = %d (stack)\n", result)
	
	// Struct passed by value
	point := Point{X: 5, Y: 10}
	moved := movePoint(point, 2, 3)
	fmt.Printf("     movePoint: %+v (stack)\n", moved)
}

func structStackExamples() {
	fmt.Println("   Struct fields (when struct doesn't escape):")
	
	// Struct allocated on stack
	person := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("     Person: %+v (stack)\n", person)
	
	// Accessing struct fields doesn't cause escape
	person.Age++
	fmt.Printf("     Person after increment: %+v (still stack)\n", person)
}

// Heap Examples (Variables that DO escape)
// =======================================
func heapExamples() {
	fmt.Println("\n3. VARIABLES THAT ESCAPE TO HEAP:")
	
	// Returning address of local variable
	addressEscape()
	
	// Storing in global variables
	globalEscape()
	
	// Interface method calls
	interfaceEscape()
	
	// Large variables
	largeVariableEscape()
	
	// Dynamic stack growth
	dynamicEscape()
}

func addressEscape() {
	fmt.Println("   Returning address of local variable:")
	
	ptr := getPointer()
	fmt.Printf("     getPointer() = %d (escaped to heap)\n", *ptr)
}

func getPointer() *int {
	x := 42  // This escapes because we return &x
	return &x
}

var globalPtr *int

func globalEscape() {
	fmt.Println("   Storing in global variable:")
	
	x := 100
	globalPtr = &x  // x escapes to heap
	fmt.Printf("     globalPtr = %d (escaped to heap)\n", *globalPtr)
}

func interfaceEscape() {
	fmt.Println("   Interface method calls:")
	
	// Interface variables are always on heap
	var writer Writer = &ConsoleWriter{}
	writer.Write("Hello")  // ConsoleWriter escapes to heap
}

func largeVariableEscape() {
	fmt.Println("   Large variables:")
	
	// Large arrays might escape to heap
	var large [10000]int
	fmt.Printf("     Large array size: %d bytes\n", unsafe.Sizeof(large))
	fmt.Println("     (Might escape to heap due to size)")
}

func dynamicEscape() {
	fmt.Println("   Dynamic stack growth:")
	
	// Recursive function might cause stack growth
	result := fibonacci(10)
	fmt.Printf("     fibonacci(10) = %d\n", result)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// How to Check Escape Analysis
// ===========================
func checkEscapeAnalysis() {
	fmt.Println("\n4. HOW TO CHECK ESCAPE ANALYSIS:")
	fmt.Println("   Use: go build -gcflags='-m' your_file.go")
	fmt.Println("   This shows which variables escape to heap")
	
	// Example of checking escape analysis
	fmt.Println("\n   Example output:")
	fmt.Println("   ./escape_analysis.go:42:6: moved escapes to heap")
	fmt.Println("   ./escape_analysis.go:45:6: &x escapes to heap")
	fmt.Println("   ./escape_analysis.go:50:6: moved escapes to heap")
}

// Optimization Techniques
// ======================
func optimizationTechniques() {
	fmt.Println("\n5. OPTIMIZATION TECHNIQUES:")
	
	// Technique 1: Avoid unnecessary pointers
	avoidUnnecessaryPointers()
	
	// Technique 2: Use value receivers when possible
	valueReceivers()
	
	// Technique 3: Pre-allocate slices with known capacity
	preAllocateSlices()
	
	// Technique 4: Use object pools for frequently allocated objects
	objectPools()
}

func avoidUnnecessaryPointers() {
	fmt.Println("   Avoid unnecessary pointers:")
	
	// BAD: Unnecessary pointer
	// func processData(data *MyData) { ... }
	
	// GOOD: Use value when possible
	// func processData(data MyData) { ... }
	
	fmt.Println("     Use values instead of pointers when possible")
}

func valueReceivers() {
	fmt.Println("   Use value receivers when possible:")
	
	// BAD: Pointer receiver when not needed
	// func (p *Point) Distance() float64 { ... }
	
	// GOOD: Value receiver
	// func (p Point) Distance() float64 { ... }
	
	fmt.Println("     Use value receivers for small structs")
}

func preAllocateSlices() {
	fmt.Println("   Pre-allocate slices with known capacity:")
	
	// BAD: Growing slice
	// var slice []int
	// for i := 0; i < 1000; i++ {
	//     slice = append(slice, i)
	// }
	
	// GOOD: Pre-allocate
	slice := make([]int, 0, 1000)  // length 0, capacity 1000
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("     Pre-allocated slice: len=%d, cap=%d\n", len(slice), cap(slice))
}

func objectPools() {
	fmt.Println("   Use object pools for frequently allocated objects:")
	
	// Example of object pool pattern
	pool := make(chan *Person, 10)
	
	// Get from pool
	var person *Person
	select {
	case person = <-pool:
		// Reuse existing object
	default:
		person = &Person{}  // Create new if pool empty
	}
	
	// Use person
	person.Name = "John"
	person.Age = 25
	
	// Return to pool
	select {
	case pool <- person:
		// Returned to pool
	default:
		// Pool full, let GC handle it
	}
	
	fmt.Println("     Object pool reduces allocation overhead")
}

// Helper functions and types
// =========================
func add(a, b int) int {
	return a + b
}

func movePoint(p Point, dx, dy int) Point {
	p.X += dx
	p.Y += dy
	return p
}

type Point struct {
	X, Y int
}

type Person struct {
	Name string
	Age  int
}

type Writer interface {
	Write(string)
}

type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(s string) {
	fmt.Printf("     ConsoleWriter: %s\n", s)
}

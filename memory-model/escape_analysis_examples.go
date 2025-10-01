package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// Go Escape Analysis Examples
// ==========================
// This file demonstrates Go's escape analysis with practical examples
// showing when variables are allocated on stack vs heap.

func main() {
	fmt.Println("=== Go Escape Analysis Examples ===")
	
	// Example 1: Variables that stay on stack
	stackAllocationExamples()
	
	// Example 2: Variables that escape to heap
	heapAllocationExamples()
	
	// Example 3: Function parameters and return values
	functionAllocationExamples()
	
	// Example 4: Struct allocation patterns
	structAllocationExamples()
	
	// Example 5: Interface and method calls
	interfaceAllocationExamples()
	
	// Example 6: Slice and array allocation
	sliceArrayAllocationExamples()
	
	// Example 7: Closure and goroutine allocation
	closureAllocationExamples()
	
	// Example 8: Large variable allocation
	largeVariableExamples()
	
	// Example 9: Global variable allocation
	globalVariableExamples()
	
	// Example 10: How to check escape analysis
	checkEscapeAnalysis()
}

// Example 1: Variables that stay on stack
// ======================================
func stackAllocationExamples() {
	fmt.Println("\n1. VARIABLES THAT STAY ON STACK:")
	
	// Simple local variables - STACK
	var a int = 42
	var b float64 = 3.14
	var c string = "Hello"
	var d bool = true
	
	fmt.Printf("   Simple variables: a=%d, b=%f, c=%s, d=%t\n", a, b, c, d)
	fmt.Println("   ✓ These stay on stack (no addresses taken)")
	
	// Arrays with known size - STACK
	var arr [5]int
	for i := range arr {
		arr[i] = i * 2
	}
	fmt.Printf("   Array: %v\n", arr)
	fmt.Println("   ✓ Small arrays stay on stack")
	
	// Struct with simple fields - STACK
	type Point struct {
		X, Y int
	}
	var p Point = Point{X: 10, Y: 20}
	fmt.Printf("   Struct: %+v\n", p)
	fmt.Println("   ✓ Simple structs stay on stack")
	
	// Function parameters and return values - STACK
	result := add(10, 20)
	fmt.Printf("   Function result: %d\n", result)
	fmt.Println("   ✓ Value parameters and returns stay on stack")
}

// Example 2: Variables that escape to heap
// =======================================
func heapAllocationExamples() {
	fmt.Println("\n2. VARIABLES THAT ESCAPE TO HEAP:")
	
	// Returning address of local variable - HEAP
	ptr := getPointer()
	fmt.Printf("   Returned pointer: %d\n", *ptr)
	fmt.Println("   ✗ Escapes to heap (returning address)")
	
	// Storing in global variable - HEAP
	storeInGlobal(100)
	fmt.Printf("   Global value: %d\n", *globalInt)
	fmt.Println("   ✗ Escapes to heap (stored in global)")
	
	// Passing to function that might store it - HEAP
	value := 200
	storeInMap(value)
	fmt.Println("   ✗ Escapes to heap (stored in map)")
	
	// Interface method calls - HEAP
	var writer Writer = &ConsoleWriter{}
	writer.Write("Hello")
	fmt.Println("   ✗ Escapes to heap (interface method call)")
}

// Example 3: Function parameters and return values
// ===============================================
func functionAllocationExamples() {
	fmt.Println("\n3. FUNCTION ALLOCATION PATTERNS:")
	
	// Value parameters - STACK
	result1 := add(10, 20)
	fmt.Printf("   add(10, 20) = %d (stack)\n", result1)
	
	// Value return - STACK
	result2 := createValue()
	fmt.Printf("   createValue() = %d (stack)\n", result2)
	
	// Pointer parameter - depends on usage
	modifyValue(30)
	fmt.Println("   modifyValue(30) - stack (value passed)")
	
	// Pointer return - HEAP
	ptr := createPointer()
	fmt.Printf("   createPointer() = %d (heap)\n", *ptr)
	
	// Struct by value - STACK
	point := Point{X: 5, Y: 10}
	moved := movePoint(point, 2, 3)
	fmt.Printf("   movePoint: %+v (stack)\n", moved)
	
	// Struct by pointer - HEAP
	pointPtr := &Point{X: 5, Y: 10}
	movePointPointer(pointPtr, 2, 3)
	fmt.Printf("   movePointPointer: %+v (heap)\n", *pointPtr)
}

// Example 4: Struct allocation patterns
// ====================================
func structAllocationExamples() {
	fmt.Println("\n4. STRUCT ALLOCATION PATTERNS:")
	
	// Struct literal - STACK
	person1 := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("   Struct literal: %+v (stack)\n", person1)
	
	// Using new() - HEAP
	person2 := new(Person)
	person2.Name = "Bob"
	person2.Age = 25
	fmt.Printf("   new(Person): %+v (heap)\n", *person2)
	
	// Taking address - HEAP
	person3 := &Person{
		Name: "Charlie",
		Age:  35,
	}
	fmt.Printf("   &Person{}: %+v (heap)\n", *person3)
	
	// Struct field access - depends on struct location
	person1.Age++  // STACK (struct is on stack)
	person2.Age++  // HEAP (struct is on heap)
	person3.Age++  // HEAP (struct is on heap)
	
	fmt.Printf("   After increment: %+v, %+v, %+v\n", person1, *person2, *person3)
}

// Example 5: Interface and method calls
// ====================================
func interfaceAllocationExamples() {
	fmt.Println("\n5. INTERFACE ALLOCATION PATTERNS:")
	
	// Interface variables - HEAP
	var writer Writer = &ConsoleWriter{}
	writer.Write("Hello from interface!")
	fmt.Println("   ✗ Interface variables escape to heap")
	
	// Empty interface - HEAP
	var any interface{} = 42
	fmt.Printf("   interface{}: %v (heap)\n", any)
	
	// Type assertion - HEAP
	if val, ok := any.(int); ok {
		fmt.Printf("   Type assertion: %d (heap)\n", val)
	}
	
	// Method calls on interfaces - HEAP
	var reader Reader = &StringReader{data: "Hello World"}
	content := reader.Read()
	fmt.Printf("   Reader content: %s (heap)\n", content)
}

// Example 6: Slice and array allocation
// ===================================
func sliceArrayAllocationExamples() {
	fmt.Println("\n6. SLICE AND ARRAY ALLOCATION:")
	
	// Small array - STACK
	var arr [5]int
	for i := range arr {
		arr[i] = i
	}
	fmt.Printf("   Small array: %v (stack)\n", arr)
	
	// Large array - might escape to HEAP
	var largeArr [1000]int
	fmt.Printf("   Large array: %d elements (might escape)\n", len(largeArr))
	
	// Slice - HEAP
	slice := make([]int, 5)
	for i := range slice {
		slice[i] = i * 2
	}
	fmt.Printf("   Slice: %v (heap)\n", slice)
	
	// Slice literal - HEAP
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Slice literal: %v (heap)\n", slice2)
	
	// Slice of structs - HEAP
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	fmt.Printf("   Slice of structs: %+v (heap)\n", people)
}

// Example 7: Closure and goroutine allocation
// =========================================
func closureAllocationExamples() {
	fmt.Println("\n7. CLOSURE ALLOCATION PATTERNS:")
	
	// Closure that captures variables - HEAP
	counter := createCounter()
	fmt.Printf("   Counter 1: %d (heap)\n", counter())
	fmt.Printf("   Counter 2: %d (heap)\n", counter())
	fmt.Printf("   Counter 3: %d (heap)\n", counter())
	fmt.Println("   ✗ Closures that capture variables escape to heap")
	
	// Closure with parameters - HEAP
	multiplier := createMultiplier(5)
	fmt.Printf("   Multiplier(3): %d (heap)\n", multiplier(3))
	fmt.Printf("   Multiplier(4): %d (heap)\n", multiplier(4))
	
	// Goroutine - each has its own stack
	go func() {
		var local [100]int
		for i := range local {
			local[i] = i
		}
		fmt.Printf("   Goroutine local array: %d elements (stack)\n", len(local))
	}()
}

// Example 8: Large variable allocation
// ===================================
func largeVariableExamples() {
	fmt.Println("\n8. LARGE VARIABLE ALLOCATION:")
	
	// Large struct - might escape to HEAP
	type LargeStruct struct {
		Data [1000]int
	}
	
	var large LargeStruct
	fmt.Printf("   Large struct size: %d bytes\n", unsafe.Sizeof(large))
	fmt.Println("   ? Large structs might escape to heap")
	
	// Large slice - HEAP
	largeSlice := make([]int, 10000)
	fmt.Printf("   Large slice: %d elements (heap)\n", len(largeSlice))
	
	// Dynamic allocation - HEAP
	dynamicSize := 5000
	dynamicSlice := make([]int, dynamicSize)
	fmt.Printf("   Dynamic slice: %d elements (heap)\n", len(dynamicSlice))
}

// Example 9: Global variable allocation
// ===================================
func globalVariableExamples() {
	fmt.Println("\n9. GLOBAL VARIABLE ALLOCATION:")
	
	// Storing in global variable - HEAP
	storeInGlobal(42)
	fmt.Printf("   Global int: %d (heap)\n", *globalInt)
	
	// Storing in global map - HEAP
	storeInGlobalMap("key", "value")
	fmt.Printf("   Global map: %v (heap)\n", globalMap)
	
	// Storing in global slice - HEAP
	storeInGlobalSlice(100)
	fmt.Printf("   Global slice: %d elements (heap)\n", len(globalSlice))
}

// Example 10: How to check escape analysis
// ========================================
func checkEscapeAnalysis() {
	fmt.Println("\n10. HOW TO CHECK ESCAPE ANALYSIS:")
	fmt.Println("   Use: go build -gcflags='-m' your_file.go")
	fmt.Println("   This shows which variables escape to heap")
	
	// Show current memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("   Current heap size: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("   GC cycles: %d\n", m.NumGC)
	
	fmt.Println("\n   Example escape analysis output:")
	fmt.Println("   ./escape_analysis_examples.go:42:6: moved escapes to heap")
	fmt.Println("   ./escape_analysis_examples.go:45:6: &x escapes to heap")
	fmt.Println("   ./escape_analysis_examples.go:50:6: moved escapes to heap")
}

// Helper functions for examples
// ============================

// Stack allocation functions
func add(a, b int) int {
	return a + b
}

func createValue() int {
	x := 42
	return x
}

func modifyValue(x int) {
	x++
}

func movePoint(p Point, dx, dy int) Point {
	p.X += dx
	p.Y += dy
	return p
}

// Heap allocation functions
func getPointer() *int {
	x := 42  // This escapes because we return &x
	return &x
}

func createPointer() *int {
	x := 100  // This escapes because we return &x
	return &x
}

func movePointPointer(p *Point, dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Global variables
var globalInt *int
var globalMap map[string]string
var globalSlice []int

func storeInGlobal(value int) {
	x := value
	globalInt = &x  // x escapes to heap
}

func storeInMap(value int) {
	if globalMap == nil {
		globalMap = make(map[string]string)
	}
	globalMap["value"] = fmt.Sprintf("%d", value)
}

func storeInGlobalMap(key, value string) {
	if globalMap == nil {
		globalMap = make(map[string]string)
	}
	globalMap[key] = value
}

func storeInGlobalSlice(value int) {
	globalSlice = append(globalSlice, value)
}

// Interface definitions
type Writer interface {
	Write(string)
}

type Reader interface {
	Read() string
}

type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(s string) {
	fmt.Printf("   ConsoleWriter: %s\n", s)
}

type StringReader struct {
	data string
}

func (sr *StringReader) Read() string {
	return sr.data
}

// Closure functions
func createCounter() func() int {
	count := 0  // This escapes to heap
	return func() int {
		count++
		return count
	}
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor  // factor is captured
	}
}

// Types
type Point struct {
	X, Y int
}

type Person struct {
	Name string
	Age  int
}

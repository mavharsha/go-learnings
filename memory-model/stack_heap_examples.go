package main

import (
	"fmt"
	"runtime"
	"time"
)

// Detailed Stack vs Heap Examples
// ===============================

func main() {
	fmt.Println("=== Stack vs Heap Allocation Examples ===")
	
	// Example 1: Basic variable allocation
	basicAllocation()
	
	// Example 2: Function parameters and return values
	functionAllocation()
	
	// Example 3: Struct allocation
	structAllocation()
	
	// Example 4: Slice and array allocation
	sliceArrayAllocation()
	
	// Example 5: Interface allocation
	interfaceAllocation()
	
	// Example 6: Closure allocation
	closureAllocation()
	
	// Example 7: Performance comparison
	performanceComparison()
}

// Example 1: Basic Variable Allocation
// ====================================
func basicAllocation() {
	fmt.Println("\n1. BASIC VARIABLE ALLOCATION:")
	
	// These are allocated on the STACK
	var a int = 10
	var b float64 = 3.14
	var c string = "Hello"
	var d bool = true
	
	fmt.Printf("   Stack variables: a=%d, b=%f, c=%s, d=%t\n", a, b, c, d)
	
	// These are allocated on the HEAP
	ptrA := &a  // Taking address of stack variable
	ptrB := new(float64)  // Explicit heap allocation
	*ptrB = 2.71
	
	fmt.Printf("   Heap pointers: ptrA=%p, ptrB=%p, *ptrB=%f\n", ptrA, ptrB, *ptrB)
}

// Example 2: Function Parameters and Return Values
// ===============================================
func functionAllocation() {
	fmt.Println("\n2. FUNCTION ALLOCATION:")
	
	// Stack allocation - value passed by copy
	result1 := addNumbers(10, 20)
	fmt.Printf("   Stack result: %d\n", result1)
	
	// Heap allocation - pointer returned
	result2 := addNumbersPointer(10, 20)
	fmt.Printf("   Heap result: %d\n", *result2)
	
	// Stack allocation - struct passed by value
	point := Point{X: 5, Y: 10}
	movedPoint := movePoint(point, 2, 3)
	fmt.Printf("   Stack point: %+v\n", movedPoint)
	
	// Heap allocation - pointer to struct
	pointPtr := &Point{X: 5, Y: 10}
	movePointPointer(pointPtr, 2, 3)
	fmt.Printf("   Heap point: %+v\n", *pointPtr)
}

// Stack allocation - returns value
func addNumbers(a, b int) int {
	return a + b  // Result stays on stack
}

// Heap allocation - returns pointer
func addNumbersPointer(a, b int) *int {
	result := a + b  // This escapes to heap
	return &result
}

type Point struct {
	X, Y int
}

// Stack allocation - struct passed by value
func movePoint(p Point, dx, dy int) Point {
	p.X += dx
	p.Y += dy
	return p  // Copy returned
}

// Heap allocation - pointer to struct
func movePointPointer(p *Point, dx, dy int) {
	p.X += dx  // Modifies original
	p.Y += dy
}

// Example 3: Struct Allocation
// ===========================
func structAllocation() {
	fmt.Println("\n3. STRUCT ALLOCATION:")
	
	// Stack allocation - struct literal
	person1 := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("   Stack person: %+v\n", person1)
	
	// Heap allocation - using new()
	person2 := new(Person)
	person2.Name = "Bob"
	person2.Age = 25
	fmt.Printf("   Heap person: %+v\n", *person2)
	
	// Heap allocation - taking address
	person3 := &Person{
		Name: "Charlie",
		Age:  35,
	}
	fmt.Printf("   Heap person (address): %+v\n", *person3)
}

type Person struct {
	Name string
	Age  int
}

// Example 4: Slice and Array Allocation
// ====================================
func sliceArrayAllocation() {
	fmt.Println("\n4. SLICE AND ARRAY ALLOCATION:")
	
	// Array - allocated on stack (if small)
	var arr [5]int
	for i := range arr {
		arr[i] = i * 2
	}
	fmt.Printf("   Stack array: %v\n", arr)
	
	// Large array - might escape to heap
	var largeArr [1000]int
	fmt.Printf("   Large array length: %d\n", len(largeArr))
	
	// Slice - always allocated on heap
	slice := make([]int, 5)
	for i := range slice {
		slice[i] = i * 3
	}
	fmt.Printf("   Heap slice: %v\n", slice)
	
	// Slice literal - also on heap
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Heap slice literal: %v\n", slice2)
}

// Example 5: Interface Allocation
// ==============================
func interfaceAllocation() {
	fmt.Println("\n5. INTERFACE ALLOCATION:")
	
	// Interface variables are allocated on heap
	var writer Writer = &ConsoleWriter{}
	writer.Write("Hello from interface!")
	
	// Interface{} (empty interface) - heap allocation
	var any interface{} = 42
	fmt.Printf("   Interface value: %v\n", any)
	
	// Type assertion
	if val, ok := any.(int); ok {
		fmt.Printf("   Type assertion: %d\n", val)
	}
}

type Writer interface {
	Write(string)
}

type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(s string) {
	fmt.Printf("   ConsoleWriter: %s\n", s)
}

// Example 6: Closure Allocation
// ============================
func closureAllocation() {
	fmt.Println("\n6. CLOSURE ALLOCATION:")
	
	// Closure captures variables - might escape to heap
	counter := createCounter()
	fmt.Printf("   Counter 1: %d\n", counter())
	fmt.Printf("   Counter 2: %d\n", counter())
	fmt.Printf("   Counter 3: %d\n", counter())
	
	// Closure with parameters
	multiplier := createMultiplier(5)
	fmt.Printf("   Multiplier(3): %d\n", multiplier(3))
	fmt.Printf("   Multiplier(4): %d\n", multiplier(4))
}

// Closure that captures variable - escapes to heap
func createCounter() func() int {
	count := 0  // This escapes to heap
	return func() int {
		count++
		return count
	}
}

// Closure with parameter
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor  // factor is captured
	}
}

// Example 7: Performance Comparison
// ================================
func performanceComparison() {
	fmt.Println("\n7. PERFORMANCE COMPARISON:")
	
	// Test stack allocation performance
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		stackOperation()
	}
	stackTime := time.Since(start)
	fmt.Printf("   Stack allocation time: %v\n", stackTime)
	
	// Test heap allocation performance
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		heapOperation()
	}
	heapTime := time.Since(start)
	fmt.Printf("   Heap allocation time: %v\n", heapTime)
	
	// Show memory statistics
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("   Heap size: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("   GC cycles: %d\n", m.NumGC)
}

func stackOperation() {
	// Fast stack operations
	var a, b, c int = 1, 2, 3
	var result int = a + b + c
	_ = result  // Prevent optimization
}

func heapOperation() {
	// Slower heap operations
	a := new(int)
	b := new(int)
	c := new(int)
	*a, *b, *c = 1, 2, 3
	result := new(int)
	*result = *a + *b + *c
	_ = result  // Prevent optimization
}

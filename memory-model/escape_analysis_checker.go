package main

import (
	"fmt"
	"runtime"
	"time"
)

// Escape Analysis Checker
// =====================
// This file demonstrates how to check and understand Go's escape analysis
// with practical examples and memory profiling.

func main() {
	fmt.Println("=== Escape Analysis Checker ===")
	
	// How to check escape analysis
	howToCheckEscapeAnalysis()
	
	// Examples with escape analysis output
	escapeAnalysisExamples()
	
	// Memory profiling examples
	memoryProfilingExamples()
	
	// Performance comparison
	performanceComparison()
	
	// Best practices for avoiding heap allocation
	bestPractices()
}

// How to Check Escape Analysis
// =============================
func howToCheckEscapeAnalysis() {
	fmt.Println("\n1. HOW TO CHECK ESCAPE ANALYSIS:")
	
	fmt.Println("   Command: go build -gcflags='-m' your_file.go")
	fmt.Println("   This shows which variables escape to heap")
	
	fmt.Println("\n   Example output:")
	fmt.Println("   ./escape_analysis_checker.go:42:6: moved escapes to heap")
	fmt.Println("   ./escape_analysis_checker.go:45:6: &x escapes to heap")
	fmt.Println("   ./escape_analysis_checker.go:50:6: moved escapes to heap")
	
	fmt.Println("\n   Understanding the output:")
	fmt.Println("   - 'escapes to heap' means variable is allocated on heap")
	fmt.Println("   - No message means variable stays on stack")
	fmt.Println("   - Line numbers show where the escape occurs")
}

// Examples with Escape Analysis Output
// ===================================
func escapeAnalysisExamples() {
	fmt.Println("\n2. ESCAPE ANALYSIS EXAMPLES:")
	
	// Example 1: Stack allocation (no escape)
	stackExample()
	
	// Example 2: Heap allocation (escape)
	heapExample()
	
	// Example 3: Function return patterns
	returnPatterns()
	
	// Example 4: Interface allocation
	interfaceExample()
	
	// Example 5: Closure allocation
	closureExample()
}

func stackExample() {
	fmt.Println("   Stack Allocation Example:")
	
	// These variables stay on stack
	var a int = 42
	var b float64 = 3.14
	var c string = "Hello"
	
	fmt.Printf("     Variables: a=%d, b=%f, c=%s\n", a, b, c)
	fmt.Println("     ✓ No escape analysis output (stays on stack)")
}

func heapExample() {
	fmt.Println("   Heap Allocation Example:")
	
	// This variable escapes to heap
	ptr := getPointer()
	fmt.Printf("     Pointer: %d\n", *ptr)
	fmt.Println("     ✗ Escape analysis will show: &x escapes to heap")
}

func returnPatterns() {
	fmt.Println("   Return Pattern Examples:")
	
	// Stack return
	result1 := returnValue()
	fmt.Printf("     Return value: %d\n", result1)
	fmt.Println("     ✓ No escape (value return)")
	
	// Heap return
	result2 := returnPointer()
	fmt.Printf("     Return pointer: %d\n", *result2)
	fmt.Println("     ✗ Escape analysis will show: &x escapes to heap")
}

func interfaceExample() {
	fmt.Println("   Interface Example:")
	
	// Interface variables escape to heap
	var writer Writer = &ConsoleWriter{}
	writer.Write("Hello")
	fmt.Println("     ✗ Escape analysis will show: interface{} escapes to heap")
}

func closureExample() {
	fmt.Println("   Closure Example:")
	
	// Closure capturing variable escapes to heap
	counter := createCounter()
	fmt.Printf("     Counter: %d\n", counter())
	fmt.Println("     ✗ Escape analysis will show: moved escapes to heap")
}

// Memory Profiling Examples
// =========================
func memoryProfilingExamples() {
	fmt.Println("\n3. MEMORY PROFILING EXAMPLES:")
	
	// Show current memory stats
	showMemoryStats()
	
	// Demonstrate heap allocation
	demonstrateHeapAllocation()
	
	// Show GC impact
	showGCImpact()
}

func showMemoryStats() {
	fmt.Println("   Current Memory Stats:")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("     Heap size: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("     Stack size: %d KB\n", m.StackInuse/1024)
	fmt.Printf("     GC cycles: %d\n", m.NumGC)
	fmt.Printf("     GC time: %v\n", time.Duration(m.PauseTotalNs))
}

func demonstrateHeapAllocation() {
	fmt.Println("   Demonstrating Heap Allocation:")
	
	// Create heap allocations
	for i := 0; i < 1000; i++ {
		ptr := new(int)
		*ptr = i
		_ = ptr
	}
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("     After heap allocation: %d KB\n", m.HeapAlloc/1024)
}

func showGCImpact() {
	fmt.Println("   GC Impact:")
	
	// Force GC
	runtime.GC()
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("     After GC: %d KB\n", m.HeapAlloc/1024)
	fmt.Printf("     GC cycles: %d\n", m.NumGC)
}

// Performance Comparison
// ====================
func performanceComparison() {
	fmt.Println("\n4. PERFORMANCE COMPARISON:")
	
	// Stack allocation benchmark
	stackBenchmark()
	
	// Heap allocation benchmark
	heapBenchmark()
	
	// Mixed allocation benchmark
	mixedBenchmark()
}

func stackBenchmark() {
	fmt.Println("   Stack Allocation Benchmark:")
	
	iterations := 1000000
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		var a, b, c int = i, i+1, i+2
		var result int = a + b + c
		_ = result
	}
	
	duration := time.Since(start)
	fmt.Printf("     %d iterations in %v\n", iterations, duration)
	fmt.Printf("     Average: %v per operation\n", duration/time.Duration(iterations))
}

func heapBenchmark() {
	fmt.Println("   Heap Allocation Benchmark:")
	
	iterations := 1000000
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		a := new(int)
		b := new(int)
		c := new(int)
		*a, *b, *c = i, i+1, i+2
		result := new(int)
		*result = *a + *b + *c
		_ = result
	}
	
	duration := time.Since(start)
	fmt.Printf("     %d iterations in %v\n", iterations, duration)
	fmt.Printf("     Average: %v per operation\n", duration/time.Duration(iterations))
}

func mixedBenchmark() {
	fmt.Println("   Mixed Allocation Benchmark:")
	
	iterations := 1000000
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		var stackVar int = i
		heapVar := new(int)
		*heapVar = i + 1
		result := stackVar + *heapVar
		_ = result
	}
	
	duration := time.Since(start)
	fmt.Printf("     %d iterations in %v\n", iterations, duration)
	fmt.Printf("     Average: %v per operation\n", duration/time.Duration(iterations))
}

// Best Practices for Avoiding Heap Allocation
// ===========================================
func bestPractices() {
	fmt.Println("\n5. BEST PRACTICES FOR AVOIDING HEAP ALLOCATION:")
	
	// Use value types when possible
	useValueTypes()
	
	// Avoid unnecessary pointers
	avoidUnnecessaryPointers()
	
	// Use value receivers
	useValueReceivers()
	
	// Pre-allocate slices
	preAllocateSlices()
	
	// Use object pools
	useObjectPools()
}

func useValueTypes() {
	fmt.Println("   Use Value Types When Possible:")
	
	// GOOD: Value type
	type Point struct {
		X, Y int
	}
	
	p := Point{X: 10, Y: 20}
	p.X += 5
	fmt.Printf("     Point: %+v (stack)\n", p)
	fmt.Println("     ✓ Use value types for small structs")
}

func avoidUnnecessaryPointers() {
	fmt.Println("   Avoid Unnecessary Pointers:")
	
	// GOOD: Pass by value
	func() {
		s := SmallStruct{Value: 42}
		result := processSmallStruct(s)
		fmt.Printf("     Small struct: %d (stack)\n", result)
	}()
	
	// BAD: Unnecessary pointer
	func() {
		s := &SmallStruct{Value: 42}
		result := processSmallStructPointer(s)
		fmt.Printf("     Small struct: %d (heap)\n", result)
	}()
}

func useValueReceivers() {
	fmt.Println("   Use Value Receivers:")
	
	// GOOD: Value receiver for small structs
	p := Point{X: 3, Y: 4}
	distance := p.Distance()
	fmt.Printf("     Distance: %f (value receiver)\n", distance)
	fmt.Println("     ✓ Use value receivers for small structs")
}

func preAllocateSlices() {
	fmt.Println("   Pre-allocate Slices:")
	
	// BAD: Growing slice
	func() {
		var slice []int
		for i := 0; i < 1000; i++ {
			slice = append(slice, i)
		}
		fmt.Printf("     Growing slice: %d elements\n", len(slice))
	}()
	
	// GOOD: Pre-allocate
	func() {
		slice := make([]int, 0, 1000)
		for i := 0; i < 1000; i++ {
			slice = append(slice, i)
		}
		fmt.Printf("     Pre-allocated slice: %d elements\n", len(slice))
		fmt.Println("     ✓ Pre-allocate slices with known capacity")
	}()
}

func useObjectPools() {
	fmt.Println("   Use Object Pools:")
	
	// Object pool for frequently allocated objects
	pool := make(chan *Person, 10)
	
	// Get from pool
	person := getFromPool(pool)
	person.Name = "John"
	person.Age = 30
	
	// Return to pool
	returnToPool(pool, person)
	fmt.Println("     ✓ Object pools reduce allocation overhead")
}

// Helper functions
// ===============

func getPointer() *int {
	x := 42
	return &x
}

func returnValue() int {
	x := 42
	return x
}

func returnPointer() *int {
	x := 42
	return &x
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func processSmallStruct(s SmallStruct) int {
	return s.Value * 2
}

func processSmallStructPointer(s *SmallStruct) int {
	return s.Value * 2
}

func getFromPool(pool chan *Person) *Person {
	select {
	case person := <-pool:
		return person
	default:
		return &Person{}
	}
}

func returnToPool(pool chan *Person, person *Person) {
	select {
	case pool <- person:
		// Returned to pool
	default:
		// Pool full, let GC handle it
	}
}

// Interface definitions
type Writer interface {
	Write(string)
}

type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(s string) {
	fmt.Printf("     ConsoleWriter: %s\n", s)
}

// Types
type SmallStruct struct {
	Value int
}

type Person struct {
	Name string
	Age  int
}

type Point struct {
	X, Y int
}

func (p Point) Distance() float64 {
	return float64(p.X*p.X + p.Y*p.Y)
}
